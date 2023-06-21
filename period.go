package period

import (
	"errors"
	"strconv"
	"strings"
	"time"
	"unicode"
)

var (
	ErrDurationOverflow      = errors.New("duration value overflow")
	ErrIncompleteNumber      = errors.New("incomplete named number")
	ErrInvalidExpression     = errors.New("invalid expression")
	ErrInvalidUnit           = errors.New("invalid unit")
	ErrMissingUnitModifier   = errors.New("unit modifier is missing")
	ErrNumberUnitIsNotUnique = errors.New("named number unit is not unique")
	ErrUnexpectedSymbol      = errors.New("unexpected symbol")
)

type Unit int

const (
	UnitUnknown Unit = iota
	UnitYear
	UnitMonth
	UnitDay
	UnitHour
	UnitMinute
	UnitSecond
	UnitMillisecond
	UnitMicrosecond
	UnitNanosecond
)

type UnitsTable map[Unit][][]rune

var defaultKnownUnits = UnitsTable{
	UnitYear: {
		[]rune("y"),
		[]rune("year"),
		[]rune("years"),
	},
	UnitMonth: {
		[]rune("mo"),
		[]rune("month"),
		[]rune("months"),
	},
	UnitDay: {
		[]rune("d"),
		[]rune("day"),
		[]rune("days"),
	},
	UnitHour: {
		[]rune("h"),
		[]rune("hour"),
		[]rune("hours"),
	},
	UnitMinute: {
		[]rune("m"),
		[]rune("minute"),
		[]rune("minutes"),
	},
	UnitSecond: {
		[]rune("s"),
		[]rune("second"),
		[]rune("seconds"),
	},
	UnitMillisecond: {
		[]rune("ms"),
		[]rune("millisecond"),
		[]rune("milliseconds"),
	},
	UnitMicrosecond: {
		[]rune("us"),
		[]rune("µs"),
		[]rune("microsecond"),
		[]rune("microseconds"),
	},
	UnitNanosecond: {
		[]rune("ns"),
		[]rune("nanosecond"),
		[]rune("nanoseconds"),
	},
}

type Period struct {
	negative bool

	years  int
	months int
	days   int

	duration time.Duration

	table UnitsTable
}

func Parse(input string) (Period, bool, error) {
	return parse(input, defaultKnownUnits)
}

func ParseCustom(input string, table UnitsTable) (Period, bool, error) {
	if err := isValidUnitsTable(table); err != nil {
		return Period{}, false, err
	}

	return parse(input, table)
}

func parse(input string, table UnitsTable) (Period, bool, error) {
	runes := []rune(input)

	negative, shift, err := isNegative(runes)
	if err != nil {
		return Period{}, false, err
	}

	found, err := findNumbers(runes[shift:], table)
	if err != nil {
		return Period{}, false, err
	}

	if len(found) == 0 {
		return Period{}, false, nil
	}

	period := Period{
		negative: negative,
		table:    table,
	}

	for unit, number := range found {
		updated, err := period.parseNumber(string(number), unit)
		if err != nil {
			return Period{}, false, err
		}

		period = updated
	}

	return period, true, nil
}

func (prd Period) parseNumber(number string, unit Unit) (Period, error) {
	updated, err := prd.parseYMDNumber(number, unit)
	if err != nil {
		return Period{}, err
	}

	return updated.parseHMSNumber(number, unit)
}

func (prd Period) parseYMDNumber(number string, unit Unit) (Period, error) {
	switch unit {
	case UnitYear:
	case UnitMonth:
	case UnitDay:
	default:
		return prd, nil
	}

	parsed, err := strconv.Atoi(number)
	if err != nil {
		return Period{}, err
	}

	switch unit {
	case UnitYear:
		prd.years = parsed
	case UnitMonth:
		prd.months = parsed
	case UnitDay:
		prd.days = parsed
	}

	return prd, nil
}

func (prd Period) parseHMSNumber(number string, unit Unit) (Period, error) {
	switch unit {
	case UnitHour:
	case UnitMinute:
	case UnitSecond:
	case UnitMillisecond:
	case UnitMicrosecond:
	case UnitNanosecond:
	default:
		return prd, nil
	}

	updated, err := prd.parseHMSIntNumber(number, unit)
	if err == nil {
		return updated, nil
	}

	return prd.parseHMSFloatNumber(number, unit)
}

func (prd Period) parseHMSIntNumber(number string, unit Unit) (Period, error) {
	parsed, err := strconv.Atoi(number)
	if err != nil {
		return Period{}, err
	}

	return prd.addInt(parsed, unit)
}

func (prd Period) addInt(parsed int, unit Unit) (Period, error) {
	added := time.Duration(parsed)

	switch unit {
	case UnitHour:
		added = added * time.Hour
	case UnitMinute:
		added = added * time.Minute
	case UnitSecond:
		added = added * time.Second
	case UnitMillisecond:
		added = added * time.Millisecond
	case UnitMicrosecond:
		added = added * time.Microsecond
	case UnitNanosecond:
		added = added * time.Nanosecond
	}

	sum, overflow := safeSum(prd.duration, added)
	if overflow {
		return Period{}, ErrDurationOverflow
	}

	prd.duration = sum

	return prd, nil
}

func (prd Period) parseHMSFloatNumber(number string, unit Unit) (Period, error) {
	parsed, err := strconv.ParseFloat(number, 64)
	if err != nil {
		return Period{}, err
	}

	return prd.addFloat(parsed, unit)
}

func (prd Period) addFloat(parsed float64, unit Unit) (Period, error) {
	var added time.Duration

	switch unit {
	case UnitHour:
		added = time.Duration(parsed * float64(time.Hour))
	case UnitMinute:
		added = time.Duration(parsed * float64(time.Minute))
	case UnitSecond:
		added = time.Duration(parsed * float64(time.Second))
	case UnitMillisecond:
		added = time.Duration(parsed * float64(time.Millisecond))
	case UnitMicrosecond:
		added = time.Duration(parsed * float64(time.Microsecond))
	case UnitNanosecond:
		added = time.Duration(parsed * float64(time.Nanosecond))
	}

	sum, overflow := safeSum(prd.duration, added)
	if overflow {
		return Period{}, ErrDurationOverflow
	}

	prd.duration = sum

	return prd, nil
}

func (prd Period) ShiftTime(base time.Time) time.Time {
	if prd.negative {
		return base.AddDate(-prd.years, -prd.months, -prd.days).Add(-prd.duration)
	}

	return base.AddDate(prd.years, prd.months, prd.days).Add(prd.duration)
}

func (prd Period) RelativeDuration(base time.Time) time.Duration {
	return prd.ShiftTime(base).Sub(base)
}

func (prd Period) String() string {
	builder := &strings.Builder{}

	if prd.negative {
		builder.WriteString("-")
	}

	if prd.years != 0 {
		builder.WriteString(strconv.Itoa(prd.years) + string(prd.table[UnitYear][0]))
	}

	if prd.months != 0 {
		builder.WriteString(strconv.Itoa(prd.months) + string(prd.table[UnitMonth][0]))
	}

	if prd.days != 0 {
		builder.WriteString(strconv.Itoa(prd.days) + string(prd.table[UnitDay][0]))
	}

	hours := prd.duration / time.Hour
	prd.duration -= hours * time.Hour

	minutes := prd.duration / time.Minute
	prd.duration -= minutes * time.Minute

	seconds := prd.duration / time.Second
	prd.duration -= seconds * time.Second

	milliseconds := prd.duration / time.Millisecond
	prd.duration -= milliseconds * time.Millisecond

	microseconds := prd.duration / time.Microsecond
	prd.duration -= microseconds * time.Microsecond

	nanoseconds := prd.duration

	secondsImitation := float64(seconds)
	secondsImitation += float64(milliseconds) * float64(time.Millisecond) / float64(time.Second)
	secondsImitation += float64(microseconds) * float64(time.Microsecond) / float64(time.Second)
	secondsImitation += float64(nanoseconds) * float64(time.Nanosecond) / float64(time.Second)

	if hours != 0 {
		builder.WriteString(strconv.Itoa(int(hours)) + string(prd.table[UnitHour][0]))
	}

	if minutes != 0 {
		builder.WriteString(strconv.Itoa(int(minutes)) + string(prd.table[UnitMinute][0]))
	}

	if secondsImitation != 0 {
		builder.WriteString(strconv.FormatFloat(secondsImitation, 'f', -1, 64) + string(prd.table[UnitSecond][0]))
	}

	return builder.String()
}

func isNegative(input []rune) (bool, int, error) {
	found := false

	for id, symbol := range input {
		if unicode.IsSpace(symbol) {
			continue
		}

		if symbol == '-' {
			found = true
			continue
		}

		if unicode.IsDigit(symbol) {
			if found {
				return true, id, nil
			}

			return false, 0, nil
		}

		return false, 0, ErrUnexpectedSymbol
	}

	if found {
		return false, 0, ErrInvalidExpression
	}

	return false, 0, nil
}

func findUnit(input []rune, table UnitsTable) (Unit, bool, int) {
	for unit, list := range table {
		for _, modifier := range list {
			if !isModifierPossibleMatch(input, modifier) {
				continue
			}

			challenger := input[:len(modifier)]

			if string(challenger) == string(modifier) {
				return unit, true, len(modifier)
			}
		}
	}

	return UnitUnknown, false, 0
}

func isModifierPossibleMatch(input []rune, modifier []rune) bool {
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
	}

	return false
}

func findNumber(input []rune, table UnitsTable) ([]rune, int, bool, Unit, error) {
	begin := -1
	doted := false

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

		if symbol == '.' && !doted {
			if begin == -1 {
				begin = id
			}

			doted = true

			continue
		}

		unit, found, next := findUnit(input[id:], table)
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

func findNumbers(input []rune, table UnitsTable) (map[Unit][]rune, error) {
	retrieved := map[Unit][]rune{}

	shift := 0

	for shift != len(input) {
		number, next, found, unit, err := findNumber(input[shift:], table)
		if err != nil {
			return nil, err
		}

		if !found {
			return nil, nil
		}

		if _, exists := retrieved[unit]; exists {
			return nil, ErrNumberUnitIsNotUnique
		}

		retrieved[unit] = number

		shift += next
	}

	return retrieved, nil
}

func isValidUnitsTable(table UnitsTable) error {
	for unit, modifier := range table {
		if err := isValidUnit(unit); err != nil {
			return err
		}

		if len(modifier) == 0 {
			return ErrMissingUnitModifier
		}

		if len(modifier[0]) == 0 {
			return ErrMissingUnitModifier
		}
	}

	return nil
}

func isValidUnit(unit Unit) error {
	switch unit {
	case UnitYear:
	case UnitMonth:
	case UnitDay:
	case UnitHour:
	case UnitMinute:
	case UnitSecond:
	case UnitMillisecond:
	case UnitMicrosecond:
	case UnitNanosecond:
	default:
		return ErrInvalidUnit
	}

	return nil
}
