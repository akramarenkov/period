package period

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestIsNegative(t *testing.T) {
	negative, next, err := isNegative([]rune("-10"))
	require.NoError(t, err)
	require.Equal(t, true, negative)
	require.Equal(t, 1, next)

	negative, next, err = isNegative([]rune("   -10"))
	require.NoError(t, err)
	require.Equal(t, true, negative)
	require.Equal(t, 4, next)

	negative, next, err = isNegative([]rune("- 10"))
	require.NoError(t, err)
	require.Equal(t, true, negative)
	require.Equal(t, 2, next)

	negative, next, err = isNegative([]rune("   -   10"))
	require.NoError(t, err)
	require.Equal(t, true, negative)
	require.Equal(t, 7, next)

	negative, next, err = isNegative([]rune(""))
	require.NoError(t, err)
	require.Equal(t, false, negative)
	require.Equal(t, 0, next)

	negative, next, err = isNegative([]rune("   "))
	require.NoError(t, err)
	require.Equal(t, false, negative)
	require.Equal(t, 0, next)

	negative, next, err = isNegative([]rune("10"))
	require.NoError(t, err)
	require.Equal(t, false, negative)
	require.Equal(t, 0, next)

	negative, next, err = isNegative([]rune("   10"))
	require.NoError(t, err)
	require.Equal(t, false, negative)
	require.Equal(t, 0, next)
}

func TestIsNegativeRequireError(t *testing.T) {
	negative, next, err := isNegative([]rune("-"))
	require.Error(t, err)
	require.Equal(t, false, negative)
	require.Equal(t, 0, next)

	negative, next, err = isNegative([]rune(" - "))
	require.Error(t, err)
	require.Equal(t, false, negative)
	require.Equal(t, 0, next)

	negative, next, err = isNegative([]rune("d"))
	require.Error(t, err)
	require.Equal(t, false, negative)
	require.Equal(t, 0, next)

	negative, next, err = isNegative([]rune("-d"))
	require.Error(t, err)
	require.Equal(t, false, negative)
	require.Equal(t, 0, next)

	negative, next, err = isNegative([]rune("   d"))
	require.Error(t, err)
	require.Equal(t, false, negative)
	require.Equal(t, 0, next)

	negative, next, err = isNegative([]rune("   -d"))
	require.Error(t, err)
	require.Equal(t, false, negative)
	require.Equal(t, 0, next)

	negative, next, err = isNegative([]rune("   -   d"))
	require.Error(t, err)
	require.Equal(t, false, negative)
	require.Equal(t, 0, next)
}

func TestFindUnit(t *testing.T) {
	unit, found, next := findUnit([]rune("y"), defaultUnits)
	require.Equal(t, UnitYear, unit)
	require.Equal(t, true, found)
	require.Equal(t, 1, next)

	unit, found, next = findUnit([]rune("y   "), defaultUnits)
	require.Equal(t, UnitYear, unit)
	require.Equal(t, true, found)
	require.Equal(t, 1, next)

	unit, found, next = findUnit([]rune("mo"), defaultUnits)
	require.Equal(t, UnitMonth, unit)
	require.Equal(t, true, found)
	require.Equal(t, 2, next)

	unit, found, next = findUnit([]rune("d"), defaultUnits)
	require.Equal(t, UnitDay, unit)
	require.Equal(t, true, found)
	require.Equal(t, 1, next)

	unit, found, next = findUnit([]rune("h"), defaultUnits)
	require.Equal(t, UnitHour, unit)
	require.Equal(t, true, found)
	require.Equal(t, 1, next)

	unit, found, next = findUnit([]rune("m"), defaultUnits)
	require.Equal(t, UnitMinute, unit)
	require.Equal(t, true, found)
	require.Equal(t, 1, next)

	unit, found, next = findUnit([]rune("s"), defaultUnits)
	require.Equal(t, UnitSecond, unit)
	require.Equal(t, true, found)
	require.Equal(t, 1, next)

	unit, found, next = findUnit([]rune("ms"), defaultUnits)
	require.Equal(t, UnitMillisecond, unit)
	require.Equal(t, true, found)
	require.Equal(t, 2, next)

	unit, found, next = findUnit([]rune("us"), defaultUnits)
	require.Equal(t, UnitMicrosecond, unit)
	require.Equal(t, true, found)
	require.Equal(t, 2, next)

	unit, found, next = findUnit([]rune("µs"), defaultUnits)
	require.Equal(t, UnitMicrosecond, unit)
	require.Equal(t, true, found)
	require.Equal(t, 2, next)

	unit, found, next = findUnit([]rune("ns"), defaultUnits)
	require.Equal(t, UnitNanosecond, unit)
	require.Equal(t, true, found)
	require.Equal(t, 2, next)
}

func TestFindUnitNotFound(t *testing.T) {
	unit, found, next := findUnit([]rune("u"), defaultUnits)
	require.Equal(t, UnitUnknown, unit)
	require.Equal(t, false, found)
	require.Equal(t, 0, next)

	unit, found, next = findUnit([]rune("n "), defaultUnits)
	require.Equal(t, UnitUnknown, unit)
	require.Equal(t, false, found)
	require.Equal(t, 0, next)
}

func TestFindNumber(t *testing.T) {
	number, next, found, unit, err := findNumber([]rune("10d"), defaultUnits)
	require.NoError(t, err)
	require.Equal(t, []rune("10"), number)
	require.Equal(t, 3, next)
	require.Equal(t, true, found)
	require.Equal(t, UnitDay, unit)

	number, next, found, unit, err = findNumber([]rune("   10d"), defaultUnits)
	require.NoError(t, err)
	require.Equal(t, []rune("10"), number)
	require.Equal(t, 6, next)
	require.Equal(t, true, found)
	require.Equal(t, UnitDay, unit)

	number, next, found, unit, err = findNumber([]rune("10d2m"), defaultUnits)
	require.NoError(t, err)
	require.Equal(t, []rune("10"), number)
	require.Equal(t, 3, next)
	require.Equal(t, true, found)
	require.Equal(t, UnitDay, unit)

	number, next, found, unit, err = findNumber([]rune("   10d2m"), defaultUnits)
	require.NoError(t, err)
	require.Equal(t, []rune("10"), number)
	require.Equal(t, 6, next)
	require.Equal(t, true, found)
	require.Equal(t, UnitDay, unit)

	number, next, found, unit, err = findNumber([]rune("1.10d"), defaultUnits)
	require.NoError(t, err)
	require.Equal(t, []rune("1.10"), number)
	require.Equal(t, 5, next)
	require.Equal(t, true, found)
	require.Equal(t, UnitDay, unit)

	number, next, found, unit, err = findNumber([]rune("   1.10d"), defaultUnits)
	require.NoError(t, err)
	require.Equal(t, []rune("1.10"), number)
	require.Equal(t, 8, next)
	require.Equal(t, true, found)
	require.Equal(t, UnitDay, unit)

	number, next, found, unit, err = findNumber([]rune(".10d"), defaultUnits)
	require.NoError(t, err)
	require.Equal(t, []rune(".10"), number)
	require.Equal(t, 4, next)
	require.Equal(t, true, found)
	require.Equal(t, UnitDay, unit)

	number, next, found, unit, err = findNumber([]rune("   .10d"), defaultUnits)
	require.NoError(t, err)
	require.Equal(t, []rune(".10"), number)
	require.Equal(t, 7, next)
	require.Equal(t, true, found)
	require.Equal(t, UnitDay, unit)

	number, next, found, unit, err = findNumber([]rune(""), defaultUnits)
	require.NoError(t, err)
	require.Equal(t, []rune(nil), number)
	require.Equal(t, 0, next)
	require.Equal(t, false, found)
	require.Equal(t, UnitUnknown, unit)

	number, next, found, unit, err = findNumber([]rune("  "), defaultUnits)
	require.NoError(t, err)
	require.Equal(t, []rune(nil), number)
	require.Equal(t, 0, next)
	require.Equal(t, false, found)
	require.Equal(t, UnitUnknown, unit)
}

func TestFindNumberRequireError(t *testing.T) {
	number, next, found, unit, err := findNumber([]rune("d"), defaultUnits)
	require.Error(t, err)
	require.Equal(t, []rune(nil), number)
	require.Equal(t, 0, next)
	require.Equal(t, false, found)
	require.Equal(t, UnitUnknown, unit)

	number, next, found, unit, err = findNumber([]rune("  d"), defaultUnits)
	require.Error(t, err)
	require.Equal(t, []rune(nil), number)
	require.Equal(t, 0, next)
	require.Equal(t, false, found)
	require.Equal(t, UnitUnknown, unit)

	number, next, found, unit, err = findNumber([]rune("-10d"), defaultUnits)
	require.Error(t, err)
	require.Equal(t, []rune(nil), number)
	require.Equal(t, 0, next)
	require.Equal(t, false, found)
	require.Equal(t, UnitUnknown, unit)

	number, next, found, unit, err = findNumber([]rune("10"), defaultUnits)
	require.Error(t, err)
	require.Equal(t, []rune(nil), number)
	require.Equal(t, 0, next)
	require.Equal(t, false, found)
	require.Equal(t, UnitUnknown, unit)

	number, next, found, unit, err = findNumber([]rune("   1 0d"), defaultUnits)
	require.Error(t, err)
	require.Equal(t, []rune(nil), number)
	require.Equal(t, 0, next)
	require.Equal(t, false, found)
	require.Equal(t, UnitUnknown, unit)

	number, next, found, unit, err = findNumber([]rune("1..10d"), defaultUnits)
	require.Error(t, err)
	require.Equal(t, []rune(nil), number)
	require.Equal(t, 0, next)
	require.Equal(t, false, found)
	require.Equal(t, UnitUnknown, unit)

	number, next, found, unit, err = findNumber([]rune("   1..10d"), defaultUnits)
	require.Error(t, err)
	require.Equal(t, []rune(nil), number)
	require.Equal(t, 0, next)
	require.Equal(t, false, found)
	require.Equal(t, UnitUnknown, unit)

	number, next, found, unit, err = findNumber([]rune("..10d"), defaultUnits)
	require.Error(t, err)
	require.Equal(t, []rune(nil), number)
	require.Equal(t, 0, next)
	require.Equal(t, false, found)
	require.Equal(t, UnitUnknown, unit)

	number, next, found, unit, err = findNumber([]rune("   ..10d"), defaultUnits)
	require.Error(t, err)
	require.Equal(t, []rune(nil), number)
	require.Equal(t, 0, next)
	require.Equal(t, false, found)
	require.Equal(t, UnitUnknown, unit)
}

func TestIsValidUnitsTable(t *testing.T) {
	require.NoError(t, IsValidUnitsTable(defaultUnits))
}

func TestIsValidUnitsTableInvalidUnit(t *testing.T) {
	table := UnitsTable{
		UnitUnknown: {
			"u",
		},
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

	require.Error(t, IsValidUnitsTable(table))
}

func TestIsValidUnitsTableMissingUnit(t *testing.T) {
	table := UnitsTable{
		UnitYear: {
			"y",
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

	require.Error(t, IsValidUnitsTable(table))
}

func TestIsValidUnitsTableMissingUnitModifier(t *testing.T) {
	table := UnitsTable{
		UnitYear: {
			"y",
		},
		UnitMonth: {},
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

	require.Error(t, IsValidUnitsTable(table))
}

func TestIsValidUnitsTableEmptyUnitModifier(t *testing.T) {
	table := UnitsTable{
		UnitYear: {
			"y",
		},
		UnitMonth: {
			"",
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

	require.Error(t, IsValidUnitsTable(table))
}

func TestIsValidUnitsTableModifierIsNotUnique(t *testing.T) {
	table := UnitsTable{
		UnitYear: {
			"y",
		},
		UnitMonth: {
			"m",
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

	require.Error(t, IsValidUnitsTable(table))
}

func TestParse(t *testing.T) {
	period, found, err := Parse(" 3mo 10d 2y 23h 59m 58s 10ms 30us 10ns")
	require.NoError(t, err)
	require.Equal(t, true, found)

	require.Equal(t, 2, period.Years())
	require.Equal(t, 3, period.Months())
	require.Equal(t, 10, period.Days())
	require.Equal(t, time.Duration(86398010030010), period.Duration())

	period, found, err = Parse(" 3mo 10d 2y 23h59m58s10ms30us10ns")
	require.NoError(t, err)
	require.Equal(t, true, found)

	require.Equal(t, 2, period.Years())
	require.Equal(t, 3, period.Months())
	require.Equal(t, 10, period.Days())
	require.Equal(t, time.Duration(86398010030010), period.Duration())

	duration, err := time.ParseDuration("23h59m58s10ms30us10ns")
	require.NoError(t, err)
	require.Equal(t, duration, period.Duration())

	period, found, err = Parse("  3mo 10d 2y 23h59m58s10ms30µs10ns")
	require.NoError(t, err)
	require.Equal(t, true, found)

	require.Equal(t, 2, period.Years())
	require.Equal(t, 3, period.Months())
	require.Equal(t, 10, period.Days())
	require.Equal(t, time.Duration(86398010030010), period.Duration())

	duration, err = time.ParseDuration("23h59m58s10ms30µs10ns")
	require.NoError(t, err)
	require.Equal(t, duration, period.Duration())
}

func TestParseNegative(t *testing.T) {
	period, found, err := Parse(" - 3mo 10d 2y 23h 59m 58s 10ms 30us 10ns")
	require.NoError(t, err)
	require.Equal(t, true, found)

	require.Equal(t, -2, period.Years())
	require.Equal(t, -3, period.Months())
	require.Equal(t, -10, period.Days())
	require.Equal(t, time.Duration(-86398010030010), period.Duration())

	period, found, err = Parse(" - 3mo 10d 2y 23h59m58s10ms30us10ns")
	require.NoError(t, err)
	require.Equal(t, true, found)

	require.Equal(t, -2, period.Years())
	require.Equal(t, -3, period.Months())
	require.Equal(t, -10, period.Days())
	require.Equal(t, time.Duration(-86398010030010), period.Duration())

	duration, err := time.ParseDuration("-23h59m58s10ms30us10ns")
	require.NoError(t, err)
	require.Equal(t, duration, period.Duration())

	period, found, err = Parse(" - 3mo 10d 2y 23h59m58s10ms30µs10ns")
	require.NoError(t, err)
	require.Equal(t, true, found)

	require.Equal(t, -2, period.Years())
	require.Equal(t, -3, period.Months())
	require.Equal(t, -10, period.Days())
	require.Equal(t, time.Duration(-86398010030010), period.Duration())

	duration, err = time.ParseDuration("-23h59m58s10ms30µs10ns")
	require.NoError(t, err)
	require.Equal(t, duration, period.Duration())
}

func TestParseOver(t *testing.T) {
	period, found, err := Parse(" - 3mo 10d 2y 52h 78m 61s 10ms 30us")
	require.NoError(t, err)
	require.Equal(t, true, found)

	require.Equal(t, -2, period.Years())
	require.Equal(t, -3, period.Months())
	require.Equal(t, -10, period.Days())
	require.Equal(t, time.Duration(-191941010030000), period.Duration())

	period, found, err = Parse(" - 3mo 10d 2y 52h78m61s10ms30us")
	require.NoError(t, err)
	require.Equal(t, true, found)

	require.Equal(t, -2, period.Years())
	require.Equal(t, -3, period.Months())
	require.Equal(t, -10, period.Days())
	require.Equal(t, time.Duration(-191941010030000), period.Duration())

	duration, err := time.ParseDuration("-52h78m61s10ms30us")
	require.NoError(t, err)
	require.Equal(t, duration, period.Duration())
}

func TestParseEmpty(t *testing.T) {
	period, found, err := Parse("")
	require.NoError(t, err)
	require.Equal(t, Period{table: defaultUnits}, period)
	require.Equal(t, false, found)

	period, found, err = Parse("   ")
	require.NoError(t, err)
	require.Equal(t, Period{table: defaultUnits}, period)
	require.Equal(t, false, found)
}

func TestParseCustom(t *testing.T) {
	input := " 3mo 10d 2y 23h 59m 58s 10ms 30us 10ns"

	periodRegular, found, err := Parse(input)
	require.NoError(t, err)
	require.Equal(t, true, found)

	periodCustom, found, err := ParseCustom(input, defaultUnits)
	require.NoError(t, err)
	require.Equal(t, true, found)

	periodUnsafe, found, err := ParseCustomUnsafe(input, defaultUnits)
	require.NoError(t, err)
	require.Equal(t, true, found)

	require.Equal(t, periodRegular, periodCustom)
	require.Equal(t, periodRegular, periodUnsafe)
}

func TestParseCustomInvalidUnitsTable(t *testing.T) {
	input := " 3mo 10d 2y 23h 59m 58s 10ms 30us 10ns"

	table := UnitsTable{
		UnitYear: {
			"y",
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

	_, _, err := ParseCustom(input, table)
	require.Error(t, err)
}

func TestParseRequireError(t *testing.T) {
	period, found, err := Parse("d")
	require.Error(t, err)
	require.Equal(t, Period{}, period)
	require.Equal(t, false, found)

	period, found, err = Parse(" - 3mo 10d 2y 23h 59m 58s 10ms 30us 20µs 10ns 1")
	require.Error(t, err)
	require.Equal(t, Period{}, period)
	require.Equal(t, false, found)

	period, found, err = Parse(" - 3mo 10d 2y 23h 59m 58s 10ms 30us 10zs")
	require.Error(t, err)
	require.Equal(t, Period{}, period)
	require.Equal(t, false, found)

	period, found, err = Parse(" - ৩mo 10d 2y 23h 59m 58s 10ms 30us 10ns")
	require.Error(t, err)
	require.Equal(t, Period{}, period)
	require.Equal(t, false, found)

	period, found, err = Parse(" - 3mo 10d 2y 23h 59m ৩s 10ms 30us 10ns")
	require.Error(t, err)
	require.Equal(t, Period{}, period)
	require.Equal(t, false, found)
}

func TestParseOverflow(t *testing.T) {
	_, _, err := Parse("9223372036854775807ns")
	require.NoError(t, err)

	_, _, err = Parse("9223372036854775808ns")
	require.Error(t, err)

	_, _, err = Parse("9223372036854775us")
	require.NoError(t, err)

	_, _, err = Parse("9223372036854776us")
	require.Error(t, err)

	_, _, err = Parse("9223372036854ms")
	require.NoError(t, err)

	_, _, err = Parse("9223372036855ms")
	require.Error(t, err)

	_, _, err = Parse("9223372036s")
	require.NoError(t, err)

	_, _, err = Parse("9223372037s")
	require.Error(t, err)

	_, _, err = Parse("153722867m")
	require.NoError(t, err)

	_, _, err = Parse("153722868m")
	require.Error(t, err)

	_, _, err = Parse("2562047h")
	require.NoError(t, err)

	_, _, err = Parse("2562048h")
	require.Error(t, err)

	_, _, err = Parse("2562047h2836s")
	require.NoError(t, err)

	_, _, err = Parse("2562047h2837s")
	require.Error(t, err)

	_, _, err = Parse("2562046.5h30.5m2837.5s")
	require.Error(t, err)

	_, _, err = Parse("9223372036854775808y")
	require.Error(t, err)

	_, _, err = Parse("9223372036854775808mo")
	require.Error(t, err)

	_, _, err = Parse("9223372036854775808d")
	require.Error(t, err)
}

func TestShiftTime(t *testing.T) {
	period, found, err := Parse("1y")
	require.NoError(t, err)
	require.Equal(t, true, found)

	date := time.Date(2023, time.April, 1, 0, 0, 0, 0, time.UTC)

	expectedDate := time.Date(2024, time.April, 1, 0, 0, 0, 0, time.UTC)
	expectedDuration := date.AddDate(1, 0, 0).Sub(date)

	unexpectedDuration := 365 * 24 * time.Hour
	unexpectedDate := date.Add(unexpectedDuration)

	require.Equal(t, expectedDate, period.ShiftTime(date))
	require.NotEqual(t, unexpectedDate, period.ShiftTime(date))

	require.Equal(t, expectedDuration, period.RelativeDuration(date))
	require.NotEqual(t, unexpectedDuration, period.RelativeDuration(date))

	period, found, err = Parse("365d24h")
	require.NoError(t, err)
	require.Equal(t, true, found)
	require.Equal(t, expectedDate, period.ShiftTime(date))
	require.Equal(t, expectedDuration, period.RelativeDuration(date))

	period, found, err = Parse("8760h1d")
	require.NoError(t, err)
	require.Equal(t, true, found)
	require.Equal(t, expectedDate, period.ShiftTime(date))
	require.Equal(t, expectedDuration, period.RelativeDuration(date))
}

func TestShiftTimeNegative(t *testing.T) {
	period, found, err := Parse("-1y")
	require.NoError(t, err)
	require.Equal(t, true, found)

	date := time.Date(2024, time.April, 1, 0, 0, 0, 0, time.UTC)

	expectedDate := time.Date(2023, time.April, 1, 0, 0, 0, 0, time.UTC)
	expectedDuration := date.AddDate(-1, 0, 0).Sub(date)

	unexpectedDuration := -365 * 24 * time.Hour
	unexpectedDate := date.Add(unexpectedDuration)

	require.Equal(t, expectedDate, period.ShiftTime(date))
	require.NotEqual(t, unexpectedDate, period.ShiftTime(date))

	require.Equal(t, expectedDuration, period.RelativeDuration(date))
	require.NotEqual(t, unexpectedDuration, period.RelativeDuration(date))

	period, found, err = Parse("-365d24h")
	require.NoError(t, err)
	require.Equal(t, true, found)
	require.Equal(t, expectedDate, period.ShiftTime(date))
	require.Equal(t, expectedDuration, period.RelativeDuration(date))

	period, found, err = Parse("-8760h1d")
	require.NoError(t, err)
	require.Equal(t, true, found)
	require.Equal(t, expectedDate, period.ShiftTime(date))
	require.Equal(t, expectedDuration, period.RelativeDuration(date))
}

func TestStringWithoutDuration(t *testing.T) {
	source := "-2y3mo10d"

	period, found, err := Parse(source)
	require.NoError(t, err)
	require.Equal(t, true, found)

	require.Equal(t, source, period.String())
}

func TestStringWithDuration(t *testing.T) {
	source := "-2y3mo10d23h59m58.01003001s"

	period, found, err := Parse(source)
	require.NoError(t, err)
	require.Equal(t, true, found)

	require.Equal(t, source, period.String())
}

func TestDurationImitation(t *testing.T) {
	source := "-23.1h59.1m58.01003001s10.1ms10.1us1.1ns"

	period, found, err := Parse(source)
	require.NoError(t, err)
	require.Equal(t, true, found)

	duration, err := time.ParseDuration(source)
	require.NoError(t, err)

	require.Equal(t, duration, period.Duration())
	require.Equal(t, duration.String(), period.String())
}

func benchmarkParseString(b *testing.B, name string) {
	input := " - 3mo 10d 2y 23h59m58s10ms30µs10ns"
	output := "-2y3mo10d23h59m58.01003001s"

	expected := Period{
		negative: true,

		years:  2,
		months: 3,
		days:   10,

		duration: 86398010030010,

		table: defaultUnits,
	}

	for attempt := 0; attempt < 100000; attempt++ {
		var (
			period Period
			found  bool
			err    error
		)

		switch name {
		case "parse":
			period, found, err = Parse(input)
		case "custom":
			period, found, err = ParseCustom(input, defaultUnits)
		case "unsafe":
			period, found, err = ParseCustomUnsafe(input, defaultUnits)
		}

		require.NoError(b, err)
		require.Equal(b, expected, period)
		require.Equal(b, true, found)

		require.Equal(b, output, period.String())
	}
}

func BenchmarkParseString(b *testing.B) {
	benchmarkParseString(b, "parse")
}

func BenchmarkParseCustomString(b *testing.B) {
	benchmarkParseString(b, "custom")
}

func BenchmarkParseCustomUnsafeString(b *testing.B) {
	benchmarkParseString(b, "unsafe")
}
