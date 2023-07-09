package period

import (
	"errors"
	"time"
	"unicode"

	"github.com/akramarenkov/safe"
)

const (
	typicalNamedNumbers = 9
)

var (
	ErrIncompleteNumber       = errors.New("incomplete named number")
	ErrInvalidExpression      = errors.New("invalid expression")
	ErrNumberUnitIsNotUnique  = errors.New("named number unit is not unique")
	ErrUnexpectedNumberFormat = errors.New("unexpected number format")
	ErrUnexpectedSymbol       = errors.New("unexpected symbol")
)

type namedNumber struct {
	Number string
	Unit   Unit
}

func isSpecialZero(input string) bool {
	if len(input) != 1 {
		return false
	}

	return input[0] == '0'
}

func isNegative(
	input string,
	minusSign byte,
	plusSign byte,
	fractionalSeparator byte,
) (bool, int, error) {
	minusFound := false
	plusFound := false

	for id, symbol := range input {
		if unicode.IsSpace(symbol) {
			continue
		}

		if symbol == rune(minusSign) && !minusFound && !plusFound {
			minusFound = true
			continue
		}

		if symbol == rune(plusSign) && !minusFound && !plusFound {
			plusFound = true
			continue
		}

		if unicode.IsDigit(symbol) || symbol == rune(fractionalSeparator) {
			if minusFound {
				return true, id, nil
			}

			if plusFound {
				return false, id, nil
			}

			return false, 0, nil
		}

		return false, 0, ErrUnexpectedSymbol
	}

	if minusFound || plusFound {
		return false, 0, ErrInvalidExpression
	}

	return false, 0, nil
}

func findNamedNumbers(
	input string,
	units UnitsTable,
	fractionalSeparator byte,
	unitsMustBeUnique bool,
) ([]namedNumber, error) {
	retrieved := make([]namedNumber, 0, typicalNamedNumbers)
	unique := make(map[Unit]struct{})

	shift := 0

	for shift != len(input) {
		number, next, found, unit, err := findNamedNumber(
			input[shift:],
			units,
			fractionalSeparator,
		)
		if err != nil {
			return nil, err
		}

		if !found {
			return nil, nil
		}

		if unitsMustBeUnique {
			if err := isUniqueUnit(unique, unit); err != nil {
				return nil, err
			}
		}

		named := namedNumber{
			Unit:   unit,
			Number: number,
		}

		retrieved = append(retrieved, named)

		shift += next
	}

	return retrieved, nil
}

func isUniqueUnit(unique map[Unit]struct{}, unit Unit) error {
	if _, exists := unique[unit]; exists {
		return ErrNumberUnitIsNotUnique
	}

	unique[unit] = struct{}{}

	return nil
}

func findNamedNumber(
	input string,
	units UnitsTable,
	fractionalSeparator byte,
) (string, int, bool, Unit, error) {
	begin := -1
	separated := false

	for id, symbol := range input {
		if unicode.IsSpace(symbol) {
			if begin != -1 {
				return "", 0, false, UnitUnknown, ErrIncompleteNumber
			}

			continue
		}

		if unicode.IsDigit(symbol) {
			if begin == -1 {
				begin = id
			}

			continue
		}

		if symbol == rune(fractionalSeparator) && !separated {
			if begin == -1 {
				begin = id
			}

			separated = true

			continue
		}

		unit, found, next := findUnit(input[id:], fractionalSeparator, units)
		if found {
			if begin == -1 {
				return "", 0, false, UnitUnknown, ErrIncompleteNumber
			}

			return input[begin:id], id + next, true, unit, nil
		}

		return "", 0, false, UnitUnknown, ErrUnexpectedSymbol
	}

	if begin != -1 {
		return "", 0, false, UnitUnknown, ErrIncompleteNumber
	}

	return "", 0, false, UnitUnknown, nil
}

func findUnit(
	input string,
	fractionalSeparator byte,
	units UnitsTable,
) (Unit, bool, int) {
	possible := pickOutPossibleUnit(input, fractionalSeparator)

	for unit, modifiers := range units {
		for _, modifier := range modifiers {
			if possible == modifier {
				return unit, true, len(modifier)
			}
		}
	}

	return UnitUnknown, false, 0
}

func pickOutPossibleUnit(input string, fractionalSeparator byte) string {
	for id, symbol := range input {
		switch {
		case unicode.IsSpace(symbol):
			return input[:id]
		case unicode.IsDigit(symbol):
			return input[:id]
		case symbol == rune(fractionalSeparator):
			return input[:id]
		}
	}

	return input
}

func parseDuration(
	named namedNumber,
	numberBase uint,
	fractionalSeparator byte,
	clear bool,
) (time.Duration, error) {
	integer, fractional, err := splitNumber(named.Number, fractionalSeparator)
	if err != nil {
		return 0, err
	}

	if clear {
		integer = clearInteger(integer)
		fractional = clearFractional(fractional, fractionalSeparator)
	}

	basic, err := parseIntegerDuration(integer, numberBase, named.Unit)
	if err != nil {
		return 0, err
	}

	additional, err := parseFractionalDuration(fractional, numberBase, named.Unit)
	if err != nil {
		return 0, err
	}

	duration, err := safe.SumInt(basic, additional)
	if err != nil {
		return 0, ErrValueOverflow // For backward compatibility
	}

	return duration, nil
}

func splitNumber(input string, fractionalSeparator byte) (string, string, error) {
	edge := -1

	for id, symbol := range input {
		if unicode.IsDigit(symbol) {
			continue
		}

		if symbol == rune(fractionalSeparator) {
			if edge != -1 {
				return "", "", ErrUnexpectedNumberFormat
			}

			edge = id

			continue
		}

		return "", "", ErrUnexpectedSymbol
	}

	if edge == -1 {
		return input, "", nil
	}

	return input[:edge], input[edge+1:], nil
}

func parseIntegerDuration(
	integerPart string,
	numberBase uint,
	unit Unit,
) (time.Duration, error) {
	number := int64(0)

	for _, symbol := range integerPart {
		digit, err := symbolToDigit(symbol)
		if err != nil {
			return 0, err
		}

		// we will assume that overflow is impossible for int64(numberBase)
		number, err = safe.ProductInt(number, int64(numberBase))
		if err != nil {
			return 0, ErrValueOverflow // For backward compatibility
		}

		// we will assume that overflow is impossible for int64(digit)
		number, err = safe.SumInt(number, int64(digit))
		if err != nil {
			return 0, ErrValueOverflow // For backward compatibility
		}
	}

	dimension, err := getDurationDimension(unit)
	if err != nil {
		return 0, err
	}

	// overflow is impossible for int64(dimension)
	number, err = safe.ProductInt(number, int64(dimension))
	if err != nil {
		return 0, ErrValueOverflow // For backward compatibility
	}

	return time.Duration(number), nil
}

func parseFractionalDuration(
	fractionalPart string,
	numberBase uint,
	unit Unit,
) (time.Duration, error) {
	number := int64(0)
	size := float64(1)

	for _, symbol := range fractionalPart {
		digit, err := symbolToDigit(symbol)
		if err != nil {
			return 0, err
		}

		// we will assume that overflow is impossible for int64(numberBase)
		product, err := safe.ProductInt(number, int64(numberBase))
		if err != nil {
			break
		}

		// we will assume that overflow is impossible for int64(digit)
		sum, err := safe.SumInt(product, int64(digit))
		if err != nil {
			break
		}

		number = sum
		size *= float64(numberBase)
	}

	dimension, err := getDurationDimension(unit)
	if err != nil {
		return 0, err
	}

	floated := float64(number) * (float64(dimension) / size)

	// overflow is not possible
	return time.Duration(floated), nil
}
