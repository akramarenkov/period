package period

import (
	"errors"
	"strconv"
	"strings"
	"time"
	"unicode"
)

const (
	dotSign   = '.'
	minusSign = '-'
)

var (
	ErrEmptyUnitModifier       = errors.New("unit modifier is empty")
	ErrIncompleteNumber        = errors.New("incomplete named number")
	ErrInvalidExpression       = errors.New("invalid expression")
	ErrInvalidUnit             = errors.New("invalid unit")
	ErrMissingUnit             = errors.New("missing unit")
	ErrMissingUnitModifier     = errors.New("unit modifier is missing")
	ErrNumberUnitIsNotUnique   = errors.New("named number unit is not unique")
	ErrUnexpectedNumberFormat  = errors.New("unexpected number format")
	ErrUnexpectedSymbol        = errors.New("unexpected symbol")
	ErrUnitModifierIsNotUnique = errors.New("unit modifier is not unique")
	ErrValueOverflow           = errors.New("value overflow")
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

const (
	validUnitsQuantity = 9
)

type UnitsTable map[Unit][]string

var defaultUnits = UnitsTable{
	UnitYear: {
		"y",
	},
	UnitMonth: {
		"mo",
	},
	UnitDay: {
		"d",
	},
	UnitHour: {
		"h",
	},
	UnitMinute: {
		"m",
	},
	UnitSecond: {
		"s",
	},
	UnitMillisecond: {
		"ms",
	},
	UnitMicrosecond: {
		"us",
		"µs",
	},
	UnitNanosecond: {
		"ns",
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
	return parse(input, defaultUnits)
}

func ParseCustom(input string, table UnitsTable) (Period, bool, error) {
	if err := IsValidUnitsTable(table); err != nil {
		return Period{}, false, err
	}

	return parse(input, table)
}

func ParseCustomUnsafe(input string, table UnitsTable) (Period, bool, error) {
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

	period := Period{
		negative: negative,
		table:    table,
	}

	if len(found) == 0 {
		return period, false, nil
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
	if isYMDUnit(unit) {
		return prd.parseYMDNumber(number, unit)
	}

	return prd.parseHMSNumber(number, unit)
}

func (prd Period) parseYMDNumber(number string, unit Unit) (Period, error) {
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
	if integer, err := strconv.ParseInt(number, 10, 64); err == nil {
		return prd.addInt(integer, unit)
	}

	if float, err := strconv.ParseFloat(number, 64); err == nil {
		return prd.addFloat(float, unit)
	}

	return Period{}, ErrUnexpectedNumberFormat
}

func (prd Period) addInt(parsed int64, unit Unit) (Period, error) {
	dimension := getDimension(unit)

	added, overflow := safeProductInt(time.Duration(parsed), dimension)
	if overflow {
		return Period{}, ErrValueOverflow
	}

	sum, overflow := safeSumInt(prd.duration, added)
	if overflow {
		return Period{}, ErrValueOverflow
	}

	prd.duration = sum

	return prd, nil
}

func (prd Period) addFloat(parsed float64, unit Unit) (Period, error) {
	dimension := getDimension(unit)

	added, overflow := safeFloatToInt[float64, time.Duration](
		parsed * float64(dimension),
	)

	if overflow {
		return Period{}, ErrValueOverflow
	}

	sum, overflow := safeSumInt(prd.duration, added)
	if overflow {
		return Period{}, ErrValueOverflow
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

func (prd Period) IsNegative() bool {
	return prd.negative
}

func (prd Period) SetNegative(negative bool) Period {
	prd.negative = negative
	return prd
}

func (prd Period) Years() int {
	if prd.negative {
		return -prd.years
	}

	return prd.years
}

func (prd Period) Months() int {
	if prd.negative {
		return -prd.months
	}

	return prd.months
}

func (prd Period) Days() int {
	if prd.negative {
		return -prd.days
	}

	return prd.days
}

func (prd Period) Duration() time.Duration {
	if prd.negative {
		return -prd.duration
	}

	return prd.duration
}

func (prd Period) sumYMD(original int, added int) (int, error) {
	if prd.negative {
		if isMaxNegative(added) {
			return 0, ErrValueOverflow
		}

		added = -added
	}

	sum, overflow := safeSumInt(original, added)
	if overflow {
		return 0, ErrValueOverflow
	}

	return sum, nil
}

func (prd Period) AddDate(years int, months int, days int) (Period, error) {
	sumYears, err := prd.sumYMD(prd.years, years)
	if err != nil {
		return Period{}, err
	}

	sumMonths, err := prd.sumYMD(prd.months, months)
	if err != nil {
		return Period{}, err
	}

	sumDays, err := prd.sumYMD(prd.days, days)
	if err != nil {
		return Period{}, err
	}

	prd.years = sumYears
	prd.months = sumMonths
	prd.days = sumDays

	return prd, nil
}

func (prd Period) sumDuration(
	original time.Duration,
	added time.Duration,
) (time.Duration, error) {
	if prd.negative {
		if isMaxNegative(added) {
			return 0, ErrValueOverflow
		}

		added = -added
	}

	sum, overflow := safeSumInt(original, added)
	if overflow {
		return 0, ErrValueOverflow
	}

	return sum, nil
}

func (prd Period) AddDuration(duration time.Duration) (Period, error) {
	sum, err := prd.sumDuration(prd.duration, duration)
	if err != nil {
		return Period{}, err
	}

	prd.duration = sum

	return prd, nil
}

func (prd Period) String() string {
	builder := &strings.Builder{}

	if prd.negative {
		builder.WriteByte(minusSign)
	}

	if prd.years != 0 {
		builder.WriteString(strconv.Itoa(prd.years))
		builder.WriteString(prd.table[UnitYear][0])
	}

	if prd.months != 0 {
		builder.WriteString(strconv.Itoa(prd.months))
		builder.WriteString(prd.table[UnitMonth][0])
	}

	if prd.days != 0 {
		builder.WriteString(strconv.Itoa(prd.days))
		builder.WriteString(prd.table[UnitDay][0])
	}

	hours, minutes, seconds := prd.countHMS()

	if hours != 0 {
		builder.WriteString(strconv.FormatInt(hours, 10))
		builder.WriteString(prd.table[UnitHour][0])
	}

	if minutes != 0 {
		builder.WriteString(strconv.FormatInt(minutes, 10))
		builder.WriteString(prd.table[UnitMinute][0])
	}

	if seconds != 0 || builder.Len() == 0 {
		builder.WriteString(strconv.FormatFloat(seconds, 'f', -1, 64))
		builder.WriteString(prd.table[UnitSecond][0])
	}

	return builder.String()
}

func (prd Period) countHMS() (int64, int64, float64) {
	remainder := prd.duration

	hours := remainder / time.Hour
	remainder -= hours * time.Hour

	minutes := remainder / time.Minute
	remainder -= minutes * time.Minute

	seconds := remainder / time.Second
	remainder -= seconds * time.Second

	milli := remainder / time.Millisecond
	remainder -= milli * time.Millisecond

	micro := remainder / time.Microsecond
	remainder -= micro * time.Microsecond

	nano := remainder

	float := float64(seconds)
	float += float64(milli) * float64(time.Millisecond) / float64(time.Second)
	float += float64(micro) * float64(time.Microsecond) / float64(time.Second)
	float += float64(nano) * float64(time.Nanosecond) / float64(time.Second)

	return int64(hours), int64(minutes), float
}

func isNegative(input []rune) (bool, int, error) {
	found := false

	for id, symbol := range input {
		if unicode.IsSpace(symbol) {
			continue
		}

		if symbol == minusSign {
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

		if symbol == dotSign && !doted {
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

func findUnit(input []rune, table UnitsTable) (Unit, bool, int) {
	for unit, modifiers := range table {
		for _, modifier := range modifiers {
			runed := []rune(modifier)

			if !isModifierPossibleMatch(input, runed) {
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

func IsValidUnitsTable(table UnitsTable) error {
	unitsQuantity := 0
	uniqueModifiers := make(map[string]struct{}, len(table))

	for unit, modifiers := range table {
		if err := isValidUnit(unit); err != nil {
			return err
		}

		unitsQuantity++

		if err := isValidModifiers(modifiers, uniqueModifiers); err != nil {
			return err
		}
	}

	if unitsQuantity != validUnitsQuantity {
		return ErrMissingUnit
	}

	return nil
}

func isValidModifiers(modifiers []string, uniqueModifiers map[string]struct{}) error {
	if len(modifiers) == 0 {
		return ErrMissingUnitModifier
	}

	for _, modifier := range modifiers {
		if len(modifier) == 0 {
			return ErrEmptyUnitModifier
		}

		if _, exists := uniqueModifiers[modifier]; exists {
			return ErrUnitModifierIsNotUnique
		}

		uniqueModifiers[modifier] = struct{}{}
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

func getDimension(unit Unit) time.Duration {
	switch unit {
	case UnitHour:
		return time.Hour
	case UnitMinute:
		return time.Minute
	case UnitSecond:
		return time.Second
	case UnitMillisecond:
		return time.Millisecond
	case UnitMicrosecond:
		return time.Microsecond
	}

	return time.Nanosecond
}

func isYMDUnit(unit Unit) bool {
	switch unit {
	case UnitYear:
	case UnitMonth:
	case UnitDay:
	default:
		return false
	}

	return true
}
