package period

import (
	"errors"
	"strconv"
	"time"
	"unicode"
)

type Unit string

const (
	UnitUnknown     Unit = ""
	UnitYear        Unit = "year"
	UnitMonth       Unit = "month"
	UnitDay         Unit = "day"
	UnitHour        Unit = "hour"
	UnitMinute      Unit = "minute"
	UnitSecond      Unit = "second"
	UnitMillisecond Unit = "millisecond"
	UnitMicrosecond Unit = "microsecond"
	UnitNanosecond  Unit = "nanosecond"
)

var (
	ErrInvalidExpression     = errors.New("invalid expression")
	ErrNumberNotFound        = errors.New("named number not found")
	ErrNumberUnitIsNotUnique = errors.New("named number unit is not unique")
	ErrUnexpectedSymbol      = errors.New("unexpected symbol")
)

func knownUnits() map[Unit][][]rune {
	units := map[Unit][][]rune{
		UnitYear: {
			[]rune("y"),
			[]rune(UnitYear),
			[]rune(UnitYear + "s"),
		},
		UnitMonth: {
			[]rune("mo"),
			[]rune(UnitMonth),
			[]rune(UnitMonth + "s"),
		},
		UnitDay: {
			[]rune("d"),
			[]rune(UnitDay),
			[]rune(UnitDay + "s"),
		},
		UnitHour: {
			[]rune("h"),
			[]rune(UnitHour),
			[]rune(UnitHour + "s"),
		},
		UnitMinute: {
			[]rune("m"),
			[]rune(UnitMinute),
			[]rune(UnitMinute + "s"),
		},
		UnitSecond: {
			[]rune("s"),
			[]rune(UnitSecond),
			[]rune(UnitSecond + "s"),
		},
		UnitMillisecond: {
			[]rune("ms"),
			[]rune(UnitMillisecond),
			[]rune(UnitMillisecond + "s"),
		},
		UnitMicrosecond: {
			[]rune("us"),
			[]rune("µs"),
			[]rune(UnitMicrosecond),
			[]rune(UnitMicrosecond + "s"),
		},
		UnitNanosecond: {
			[]rune("ns"),
			[]rune(UnitNanosecond),
			[]rune(UnitNanosecond + "s"),
		},
	}

	return units
}

type Period struct {
	Years  int
	Months int
	Days   int

	Duration time.Duration
}

func Parse(input string) (Period, bool, error) {
	runes := []rune(input)

	retrieved := map[Unit][]rune{}

	negative, shift, err := isNegative(runes)
	if err != nil {
		return Period{}, false, err
	}

	for shift != len(runes) {
		number, next, found, unit, err := findOneNamedNumber(runes[shift:])
		if err != nil {
			return Period{}, false, err
		}

		if !found {
			return Period{}, false, nil
		}

		shift += next

		if _, exists := retrieved[unit]; exists {
			return Period{}, false, ErrNumberUnitIsNotUnique
		}

		retrieved[unit] = number
	}

	if len(retrieved) == 0 {
		return Period{}, false, nil
	}

	period := Period{}

	for unit, number := range retrieved {
		parsed, err := strconv.Atoi(string(number))
		if err != nil {
			return Period{}, false, err
		}

		if negative {
			parsed = -parsed
		}

		switch unit {
		case UnitYear:
			period.Years = parsed
		case UnitMonth:
			period.Months = parsed
		case UnitDay:
			period.Days = parsed
		case UnitHour:
			period.Duration += time.Duration(parsed) * time.Hour
		case UnitMinute:
			period.Duration += time.Duration(parsed) * time.Minute
		case UnitSecond:
			period.Duration += time.Duration(parsed) * time.Second
		case UnitMillisecond:
			period.Duration += time.Duration(parsed) * time.Millisecond
		case UnitMicrosecond:
			period.Duration += time.Duration(parsed) * time.Microsecond
		case UnitNanosecond:
			period.Duration += time.Duration(parsed) * time.Nanosecond
		}
	}

	return period, true, nil
}

func (prd Period) ShiftTime(base time.Time) time.Time {
	return base.AddDate(prd.Years, prd.Months, prd.Days).Add(prd.Duration)
}

func (prd Period) RelativeDuration(base time.Time) time.Duration {
	updated := base.AddDate(prd.Years, prd.Months, prd.Days).Add(prd.Duration)
	return base.Sub(updated)
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
			if !isPossibleMatch(input, modifier) {
				continue
			}

			challenger := input[:len(modifier)]

			if string(challenger) == string(modifier) {
				return unit, true, len(modifier)
			}
		}
	}

	return "", false, 0
}

func isPossibleMatch(input []rune, modifier []rune) bool {
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

func findOneNamedNumber(input []rune) ([]rune, int, bool, Unit, error) {
	begin := -1

	for id, symbol := range input {
		if unicode.IsSpace(symbol) {
			if begin != -1 {
				return nil, 0, false, UnitUnknown, ErrNumberNotFound
			}

			continue
		}

		if unicode.IsDigit(symbol) {
			if begin == -1 {
				begin = id
			}

			continue
		}

		unit, found, next := findUnit(input[id:])
		if found {
			if begin == -1 {
				return nil, 0, false, UnitUnknown, ErrNumberNotFound
			}

			return input[begin:id], id + next, true, unit, nil
		}

		return nil, 0, false, UnitUnknown, ErrUnexpectedSymbol
	}

	if begin != -1 {
		return nil, 0, false, UnitUnknown, ErrNumberNotFound
	}

	return nil, 0, false, UnitUnknown, nil
}
