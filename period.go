package period

import (
	"errors"
	"strconv"
	"strings"
	"time"

	"github.com/akramarenkov/safe"
	"golang.org/x/exp/constraints"
)

var (
	ErrUnexpectedNumberSign = errors.New("unexpected number sign")
)

type Opts struct {
	// Provides more accurate parsing in the presence of non-significant zeros in
	// the input string
	ExtraZerosResistance bool
	// Disables validates units table
	NotValidateUnits bool
	Units            UnitsTable
	// Enables checking for units uniqueness in the input string
	UnitsMustBeUnique bool
}

type Period struct {
	opts Opts

	negative bool

	years  int
	months int
	days   int

	duration time.Duration
}

// Creates empty Period instance with default units table
func New() Period {
	opts := Opts{
		Units: defaultUnits,
	}

	return newPeriod(opts)
}

// Creates empty Period instance with custom units table.
//
// Units table validates before create instance
func NewCustom(units UnitsTable) (Period, error) {
	if err := IsValidUnitsTable(units); err != nil {
		return Period{}, err
	}

	opts := Opts{
		Units: units,
	}

	return newPeriod(opts), nil
}

// Creates empty Period instance with custom units table.
//
// Units table not validates before create instance. Use IsValidUnitsTable()
// yourself before creates instance
func NewCustomUnsafe(units UnitsTable) Period {
	opts := Opts{
		Units: units,
	}

	return newPeriod(opts)
}

// Creates empty Period instance with options.
func NewWithOpts(opts Opts) (Period, error) {
	if !opts.NotValidateUnits {
		if err := IsValidUnitsTable(opts.Units); err != nil {
			return Period{}, err
		}
	}

	return newPeriod(opts), nil
}

func newPeriod(opts Opts) Period {
	prd := Period{
		opts: opts,
	}

	return prd
}

// Creates Period instance from input string with default units table.
func Parse(input string) (Period, bool, error) {
	opts := Opts{
		Units: defaultUnits,
	}

	return parse(input, opts)
}

// Creates Period instance from input string with custom units table.
//
// Units table validates before create instance
func ParseCustom(input string, units UnitsTable) (Period, bool, error) {
	if err := IsValidUnitsTable(units); err != nil {
		return Period{}, false, err
	}

	opts := Opts{
		Units: units,
	}

	return parse(input, opts)
}

// Creates Period instance from input string with custom units table.
//
// Units table not validates before create instance. Use IsValidUnitsTable()
// yourself before creates instance
func ParseCustomUnsafe(input string, units UnitsTable) (Period, bool, error) {
	opts := Opts{
		Units: units,
	}

	return parse(input, opts)
}

// Creates Period instance from input string with options.
func ParseWithOpts(input string, opts Opts) (Period, bool, error) {
	if !opts.NotValidateUnits {
		if err := IsValidUnitsTable(opts.Units); err != nil {
			return Period{}, false, err
		}
	}

	return parse(input, opts)
}

func (prd *Period) Parse(input string) (bool, error) {
	period, found, err := parse(input, prd.opts)
	if err != nil {
		return false, err
	}

	*prd = period

	return found, nil
}

func parse(input string, opts Opts) (Period, bool, error) {
	runes := []rune(input)

	negative, shift, err := isNegative(
		runes,
		defaultMinusSign,
		defaultPlusSign,
		defaultFractionalSeparator,
	)
	if err != nil {
		return Period{}, false, err
	}

	if isSpecialZero(runes[shift:]) {
		return Period{opts: opts}, true, nil
	}

	found, err := findNamedNumbers(
		runes[shift:],
		opts.Units,
		defaultFractionalSeparator,
		opts.UnitsMustBeUnique,
	)
	if err != nil {
		return Period{}, false, err
	}

	period := Period{
		opts:     opts,
		negative: negative,
	}

	if len(found) == 0 {
		return period, false, nil
	}

	for _, named := range found {
		updated, err := period.parseNumber(named)
		if err != nil {
			return Period{}, false, err
		}

		period = updated
	}

	return period, true, nil
}

func (prd Period) parseNumber(named namedNumber) (Period, error) {
	if isYMDUnit(named.Unit) {
		return prd.parseYMDNumber(named)
	}

	return prd.parseHMSNumber(named)
}

func (prd Period) parseYMDNumber(named namedNumber) (Period, error) {
	parsed, err := strconv.ParseInt(string(named.Number), int(defaultNumberBase), 0)
	if err != nil {
		return Period{}, err
	}

	switch named.Unit {
	case UnitYear:
		years, err := safe.SumInt(prd.years, int(parsed))
		if err != nil {
			return Period{}, ErrValueOverflow // For backward compatibility
		}

		prd.years = years
	case UnitMonth:
		months, err := safe.SumInt(prd.months, int(parsed))
		if err != nil {
			return Period{}, ErrValueOverflow // For backward compatibility
		}

		prd.months = months
	case UnitDay:
		days, err := safe.SumInt(prd.days, int(parsed))
		if err != nil {
			return Period{}, ErrValueOverflow // For backward compatibility
		}

		prd.days = days
	}

	return prd, nil
}

func (prd Period) parseHMSNumber(named namedNumber) (Period, error) {
	duration, err := parseDuration(
		named,
		defaultNumberBase,
		defaultFractionalSeparator,
		prd.opts.ExtraZerosResistance,
	)
	if err != nil {
		return Period{}, err
	}

	duration, err = safe.SumInt(prd.duration, duration)
	if err != nil {
		return Period{}, ErrValueOverflow // For backward compatibility
	}

	prd.duration = duration

	return prd, nil
}

// Shifts base time to Period value
func (prd Period) ShiftTime(base time.Time) time.Time {
	if prd.negative {
		return base.AddDate(-prd.years, -prd.months, -prd.days).Add(-prd.duration)
	}

	return base.AddDate(prd.years, prd.months, prd.days).Add(prd.duration)
}

// Calculates Period value in time.Duration.
//
// Base time is necessary because shift to days, months and years
// not deterministic and depends on time around which it occurs
func (prd Period) RelativeDuration(base time.Time) time.Duration {
	return prd.ShiftTime(base).Sub(base)
}

// Returns Period sign (negative or positive)
func (prd Period) IsNegative() bool {
	return prd.negative
}

// Sets Period sign (negative or positive)
func (prd *Period) SetNegative(negative bool) {
	prd.negative = negative
}

// Returns years separately
func (prd Period) Years() int {
	if prd.negative {
		return -prd.years
	}

	return prd.years
}

// Sets years separately
func (prd *Period) SetYears(years int) error {
	years, err := normalizeValue(prd.negative, years)
	if err != nil {
		return err
	}

	prd.years = years

	return nil
}

// Returns months separately
func (prd Period) Months() int {
	if prd.negative {
		return -prd.months
	}

	return prd.months
}

// Sets months separately
func (prd *Period) SetMonths(months int) error {
	months, err := normalizeValue(prd.negative, months)
	if err != nil {
		return err
	}

	prd.months = months

	return nil
}

// Returns days separately
func (prd Period) Days() int {
	if prd.negative {
		return -prd.days
	}

	return prd.days
}

// Sets days separately
func (prd *Period) SetDays(days int) error {
	days, err := normalizeValue(prd.negative, days)
	if err != nil {
		return err
	}

	prd.days = days

	return nil
}

// Returns duration part separately.
//
// It is not Period duration, it is part of Period with value of
// hours, minutes, seconds and etc.
//
// For get Period duration use RelativeDuration()
func (prd Period) Duration() time.Duration {
	if prd.negative {
		return -prd.duration
	}

	return prd.duration
}

// Sets duration part separately.
//
// It is not Period duration, it is part of Period with value of
// hours, minutes, seconds and etc.
func (prd *Period) SetDuration(duration time.Duration) error {
	duration, err := normalizeValue(prd.negative, duration)
	if err != nil {
		return err
	}

	prd.duration = duration

	return nil
}

func normalizeValue[Type constraints.Signed](
	negative bool,
	value Type,
) (Type, error) {
	if value < 0 && !negative {
		return 0, ErrUnexpectedNumberSign
	}

	if value > 0 && negative {
		return 0, ErrUnexpectedNumberSign
	}

	if value < 0 {
		inverted, err := safe.Invert(value)
		if err != nil {
			return 0, ErrValueOverflow // For backward compatibility
		}

		return inverted, nil
	}

	return value, nil
}

// Increases or decreases value of years, months and days
func (prd *Period) AddDate(years int, months int, days int) error {
	sumYears, err := addValue(prd.negative, prd.years, years)
	if err != nil {
		return err
	}

	sumMonths, err := addValue(prd.negative, prd.months, months)
	if err != nil {
		return err
	}

	sumDays, err := addValue(prd.negative, prd.days, days)
	if err != nil {
		return err
	}

	prd.years = sumYears
	prd.months = sumMonths
	prd.days = sumDays

	return nil
}

// Increases or decreases duration part.
//
// It is not Period duration, it is part of Period with value of
// hours, minutes, seconds and etc.
func (prd *Period) AddDuration(duration time.Duration) error {
	sum, err := addValue(prd.negative, prd.duration, duration)
	if err != nil {
		return err
	}

	prd.duration = sum

	return nil
}

func addValue[Type constraints.Signed](
	negative bool,
	original Type,
	added Type,
) (Type, error) {
	if negative {
		inverted, err := safe.Invert(added)
		if err != nil {
			return 0, ErrValueOverflow // For backward compatibility
		}

		sum, err := safe.SumInt(original, inverted)
		if err != nil {
			return 0, ErrValueOverflow // For backward compatibility
		}

		return sum, nil
	}

	sum, err := safe.SumInt(original, added)
	if err != nil {
		return 0, ErrValueOverflow // For backward compatibility
	}

	return sum, nil
}

// Converts Period value into string
func (prd Period) String() string {
	builder := &strings.Builder{}

	if prd.isZero() {
		return "0s"
	}

	if prd.negative && !prd.isZero() {
		builder.WriteByte(defaultMinusSign)
	}

	upperWritten := prd.writeYMD(builder)

	prd.writeHMS(builder, upperWritten)

	return builder.String()
}

func (prd Period) isZero() bool {
	return prd.years == 0 && prd.months == 0 && prd.days == 0 && prd.duration == 0
}

func (prd Period) writeYMD(builder *strings.Builder) bool {
	upperWritten := false

	if prd.years != 0 {
		upperWritten = true

		prd.writeNumber(builder, int64(prd.years), 0, UnitYear)
	}

	if prd.months != 0 || upperWritten {
		upperWritten = true

		prd.writeNumber(builder, int64(prd.months), 0, UnitMonth)
	}

	if prd.days != 0 || upperWritten {
		upperWritten = true

		prd.writeNumber(builder, int64(prd.days), 0, UnitDay)
	}

	return upperWritten
}

func (prd Period) writeHMS(builder *strings.Builder, upperWritten bool) {
	hours, minutes, seconds, remainder := calcHMS(prd.duration)

	if hours != 0 || upperWritten {
		upperWritten = true

		prd.writeNumber(builder, int64(hours), 0, UnitHour)
	}

	if minutes != 0 || upperWritten {
		upperWritten = true

		prd.writeNumber(builder, int64(minutes), 0, UnitMinute)
	}

	if seconds != 0 || upperWritten {
		prd.writeNumber(builder, int64(seconds), int64(remainder), UnitSecond)
		return
	}

	milli, milliFractional, micro, microFractional, nano := calcMMN(remainder)

	if milli != 0 {
		prd.writeNumber(builder, int64(milli), int64(milliFractional), UnitMillisecond)
		return
	}

	if micro != 0 {
		prd.writeNumber(builder, int64(micro), int64(microFractional), UnitMicrosecond)
		return
	}

	if nano != 0 {
		prd.writeNumber(builder, int64(remainder), 0, UnitNanosecond)
	}
}

func calcHMS(duration time.Duration) (
	time.Duration,
	time.Duration,
	time.Duration,
	time.Duration,
) {
	hours := duration / time.Hour
	duration -= hours * time.Hour

	minutes := duration / time.Minute
	duration -= minutes * time.Minute

	seconds := duration / time.Second
	duration -= seconds * time.Second

	return hours, minutes, seconds, duration
}

func calcMMN(remainder time.Duration) (
	time.Duration,
	time.Duration,
	time.Duration,
	time.Duration,
	time.Duration,
) {
	milli := remainder / time.Millisecond
	remainder -= milli * time.Millisecond
	milliFractional := remainder * time.Second / time.Millisecond

	micro := remainder / time.Microsecond
	remainder -= micro * time.Microsecond
	microFractional := remainder * time.Second / time.Microsecond

	return milli, milliFractional, micro, microFractional, remainder
}

func (prd Period) writeNumber(
	builder *strings.Builder,
	integer int64,
	fractional int64,
	unit Unit,
) {
	builder.WriteString(strconv.FormatInt(integer, int(defaultNumberBase)))

	if fractional != 0 {
		formated, err := formatFractional(
			fractional,
			defaultNumberBase,
			defaultFormatFractionalSize,
			defaultFractionalSeparator,
		)
		if err == nil {
			builder.WriteString(formated)
		}
	}

	builder.WriteString(prd.opts.Units[unit][0])
}
