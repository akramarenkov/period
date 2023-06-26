package period

import (
	"errors"
	"time"
	"unicode"
)

var (
	ErrIncompleteNumber       = errors.New("incomplete named number")
	ErrInvalidExpression      = errors.New("invalid expression")
	ErrNumberUnitIsNotUnique  = errors.New("named number unit is not unique")
	ErrUnexpectedNumberFormat = errors.New("unexpected number format")
	ErrUnexpectedSymbol       = errors.New("unexpected symbol")
)

type namedNumber struct {
	Number []rune
	Unit   Unit
}

func isSpecialZero(input []rune) bool {
	if len(input) != 1 {
		return false
	}

	return input[0] == '0'
}

func isNegative(
	input []rune,
	minusSign rune,
	plusSign rune,
	fractionalSeparator rune,
) (bool, int, error) {
	minusFound := false
	plusFound := false

	for id, symbol := range input {
		if unicode.IsSpace(symbol) {
			continue
		}

		if symbol == minusSign && !minusFound && !plusFound {
			minusFound = true
			continue
		}

		if symbol == plusSign && !minusFound && !plusFound {
			plusFound = true
			continue
		}

		if unicode.IsDigit(symbol) || symbol == fractionalSeparator {
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
	input []rune,
	units UnitsTable,
	fractionalSeparator rune,
	unitsMustBeUnique bool,
) ([]namedNumber, error) {
	retrieved := make([]namedNumber, 0)
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
	input []rune,
	units UnitsTable,
	fractionalSeparator rune,
) ([]rune, int, bool, Unit, error) {
	begin := -1
	separated := false

	for id, symbol := range input {
		if unicode.IsSpace(symbol) {
			if begin != -1 {
				return nil, 0, false, UnitUnknown, ErrIncompleteNumber
			}

			continue
		}

		if unicode.IsDigit(symbol) {
			if begin == -1 {
				begin = id
			}

			continue
		}

		if symbol == fractionalSeparator && !separated {
			if begin == -1 {
				begin = id
			}

			separated = true

			continue
		}

		unit, found, next := findUnit(input[id:], fractionalSeparator, units)
		if found {
			if begin == -1 {
				return nil, 0, false, UnitUnknown, ErrIncompleteNumber
			}

			return input[begin:id], id + next, true, unit, nil
		}

		return nil, 0, false, UnitUnknown, ErrUnexpectedSymbol
	}

	if begin != -1 {
		return nil, 0, false, UnitUnknown, ErrIncompleteNumber
	}

	return nil, 0, false, UnitUnknown, nil
}

func findUnit(
	input []rune,
	fractionalSeparator rune,
	units UnitsTable,
) (Unit, bool, int) {
	for unit, modifiers := range units {
		for _, modifier := range modifiers {
			runed := []rune(modifier)

			if !isModifierPossibleMatch(input, runed, fractionalSeparator) {
				continue
			}

			challenger := input[:len(runed)]

			if string(challenger) == modifier {
				return unit, true, len(runed)
			}
		}
	}

	return UnitUnknown, false, 0
}

func isModifierPossibleMatch(
	input []rune,
	modifier []rune,
	fractionalSeparator rune,
) bool {
	if len(input) < len(modifier) {
		return false
	}

	if len(input) == len(modifier) {
		return true
	}

	after := input[len(modifier)]

	switch {
	case unicode.IsSpace(after):
		return true
	case unicode.IsDigit(after):
		return true
	case after == fractionalSeparator:
		return true
	}

	return false
}

func parseDuration(
	named namedNumber,
	numberBase uint,
	fractionalSeparator rune,
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

	duration, err := safeSumInt(basic, additional)
	if err != nil {
		return 0, err
	}

	return duration, nil
}

func splitNumber(input []rune, fractionalSeparator rune) ([]rune, []rune, error) {
	edge := -1

	for id, symbol := range input {
		if unicode.IsDigit(symbol) {
			continue
		}

		if symbol == fractionalSeparator {
			if edge != -1 {
				return nil, nil, ErrUnexpectedNumberFormat
			}

			edge = id

			continue
		}

		return nil, nil, ErrUnexpectedSymbol
	}

	if edge == -1 {
		return input, nil, nil
	}

	return input[:edge], input[edge+1:], nil
}

func parseIntegerDuration(
	integerPart []rune,
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
		number, err = safeProductInt(number, int64(numberBase))
		if err != nil {
			return 0, err
		}

		// we will assume that overflow is impossible for int64(digit)
		number, err = safeSumInt(number, int64(digit))
		if err != nil {
			return 0, err
		}
	}

	dimension, err := getDurationDimension(unit)
	if err != nil {
		return 0, err
	}

	// overflow is impossible for int64(dimension)
	number, err = safeProductInt(number, int64(dimension))
	if err != nil {
		return 0, err
	}

	return time.Duration(number), nil
}

func parseFractionalDuration(
	fractionalPart []rune,
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
		product, err := safeProductInt(number, int64(numberBase))
		if err != nil {
			break
		}

		// we will assume that overflow is impossible for int64(digit)
		sum, err := safeSumInt(product, int64(digit))
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
