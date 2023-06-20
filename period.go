package period

import (
	"errors"
	"strconv"
	"strings"
	"time"
	"unicode"
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

var (
	ErrIncompleteNumber      = errors.New("incomplete named number")
	ErrInvalidExpression     = errors.New("invalid expression")
	ErrNumberUnitIsNotUnique = errors.New("named number unit is not unique")
	ErrUnexpectedSymbol      = errors.New("unexpected symbol")
)

func knownUnits() map[Unit][][]rune {
	units := map[Unit][][]rune{
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

	return units
}

type Period struct {
	negative bool

	years  int
	months int
	days   int

	duration time.Duration
}

func Parse(input string) (Period, bool, error) {
	runes := []rune(input)

	negative, shift, err := isNegative(runes)
	if err != nil {
		return Period{}, false, err
	}

	found, err := findNamedNumbers(runes[shift:])
	if err != nil {
		return Period{}, false, err
	}

	if len(found) == 0 {
		return Period{}, false, nil
	}

	period := Period{
		negative: negative,
	}

	for unit, number := range found {
		parsed, err := strconv.Atoi(string(number))
		if err != nil {
			return Period{}, false, err
		}

		switch unit {
		case UnitYear:
			period.years = parsed
		case UnitMonth:
			period.months = parsed
		case UnitDay:
			period.days = parsed
		case UnitHour:
			period.duration += time.Duration(parsed) * time.Hour
		case UnitMinute:
			period.duration += time.Duration(parsed) * time.Minute
		case UnitSecond:
			period.duration += time.Duration(parsed) * time.Second
		case UnitMillisecond:
			period.duration += time.Duration(parsed) * time.Millisecond
		case UnitMicrosecond:
			period.duration += time.Duration(parsed) * time.Microsecond
		case UnitNanosecond:
			period.duration += time.Duration(parsed) * time.Nanosecond
		}
	}

	return period, true, nil
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
	builder := strings.Builder{}

	units := knownUnits()

	if prd.negative {
		builder.WriteString("-")
	}

	if prd.years != 0 {
		builder.WriteString(strconv.Itoa(prd.years) + string(units[UnitYear][0]))
	}

	if prd.months != 0 {
		builder.WriteString(strconv.Itoa(prd.months) + string(units[UnitMonth][0]))
	}

	if prd.days != 0 {
		builder.WriteString(strconv.Itoa(prd.days) + string(units[UnitDay][0]))
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

	if hours != 0 {
		builder.WriteString(strconv.Itoa(int(hours)) + string(units[UnitHour][0]))
	}

	if minutes != 0 {
		builder.WriteString(strconv.Itoa(int(minutes)) + string(units[UnitMinute][0]))
	}

	if seconds != 0 {
		builder.WriteString(strconv.Itoa(int(seconds)) + string(units[UnitSecond][0]))
	}

	if milliseconds != 0 {
		builder.WriteString(strconv.Itoa(int(milliseconds)) + string(units[UnitMillisecond][0]))
	}

	if microseconds != 0 {
		builder.WriteString(strconv.Itoa(int(microseconds)) + string(units[UnitMicrosecond][0]))
	}

	if nanoseconds != 0 {
		builder.WriteString(strconv.Itoa(int(nanoseconds)) + string(units[UnitNanosecond][0]))
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

func findUnit(input []rune) (Unit, bool, int) {
	for unit, list := range knownUnits() {
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

func findNamedNumber(input []rune) ([]rune, int, bool, Unit, error) {
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

		unit, found, next := findUnit(input[id:])
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

func findNamedNumbers(input []rune) (map[Unit][]rune, error) {
	retrieved := map[Unit][]rune{}

	shift := 0

	for shift != len(input) {
		number, next, found, unit, err := findNamedNumber(input[shift:])
		if err != nil {
			return nil, err
		}

		if !found {
			return nil, nil
		}

		shift += next

		if _, exists := retrieved[unit]; exists {
			return nil, ErrNumberUnitIsNotUnique
		}

		retrieved[unit] = number
	}

	return retrieved, nil
}
