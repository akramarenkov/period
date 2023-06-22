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
	period, found, err := Parse("10d")
	require.NoError(t, err)
	require.Equal(t, Period{days: 10, table: defaultUnits}, period)
	require.Equal(t, true, found)

	period, found, err = Parse("   10d")
	require.NoError(t, err)
	require.Equal(t, Period{days: 10, table: defaultUnits}, period)
	require.Equal(t, true, found)

	period, found, err = Parse("-10d")
	require.NoError(t, err)
	require.Equal(t, Period{negative: true, days: 10, table: defaultUnits}, period)
	require.Equal(t, true, found)

	period, found, err = Parse("   -   10d")
	require.NoError(t, err)
	require.Equal(t, Period{negative: true, days: 10, table: defaultUnits}, period)
	require.Equal(t, true, found)

	expected := Period{
		negative: true,

		years:  2,
		months: 3,
		days:   10,

		duration: 86398010030010,

		table: defaultUnits,
	}

	period, found, err = Parse(" - 3mo 10d 2y 23h 59m 58s 10ms 30us 10ns")
	require.NoError(t, err)
	require.Equal(t, expected, period)
	require.Equal(t, true, found)

	period, found, err = Parse(" - 3mo 10d 2y 23h59m58s10ms30us10ns")
	require.NoError(t, err)
	require.Equal(t, expected, period)
	require.Equal(t, true, found)

	duration, err := time.ParseDuration("-23h59m58s10ms30us10ns")
	require.NoError(t, err)
	require.Equal(t, duration, -period.duration)

	period, found, err = Parse(" - 3mo 10d 2y 23h59m58s10ms30µs10ns")
	require.NoError(t, err)
	require.Equal(t, expected, period)
	require.Equal(t, true, found)

	duration, err = time.ParseDuration("-23h59m58s10ms30µs10ns")
	require.NoError(t, err)
	require.Equal(t, duration, -period.duration)

	expected = Period{
		negative: true,

		years:  2,
		months: 3,
		days:   10,

		duration: 191941010030000,

		table: defaultUnits,
	}

	period, found, err = Parse(" - 3mo 10d 2y 52h 78m 61s 10ms 30us")
	require.NoError(t, err)
	require.Equal(t, expected, period)
	require.Equal(t, true, found)

	period, found, err = Parse(" - 3mo 10d 2y 52h78m61s10ms30us")
	require.NoError(t, err)
	require.Equal(t, expected, period)
	require.Equal(t, true, found)

	duration, err = time.ParseDuration("-52h78m61s10ms30us")
	require.NoError(t, err)
	require.Equal(t, duration, -period.duration)

	period, found, err = Parse("")
	require.NoError(t, err)
	require.Equal(t, Period{table: defaultUnits}, period)
	require.Equal(t, false, found)

	period, found, err = Parse("   ")
	require.NoError(t, err)
	require.Equal(t, Period{table: defaultUnits}, period)
	require.Equal(t, false, found)

	period, found, err = Parse("d")
	require.Error(t, err)
	require.Equal(t, Period{}, period)
	require.Equal(t, false, found)

	period, found, err = Parse(" d")
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

func TestCorretness(t *testing.T) {
	period := Period{
		years: 1,
	}

	date := time.Date(2023, time.April, 1, 0, 0, 0, 0, time.UTC)

	expectedDate := time.Date(2024, time.April, 1, 0, 0, 0, 0, time.UTC)
	expectedDuration := date.AddDate(1, 0, 0).Sub(date)

	unexpectedDuration := 365 * 24 * time.Hour
	unexpectedDate := date.Add(unexpectedDuration)

	require.Equal(t, expectedDate, period.ShiftTime(date))
	require.NotEqual(t, unexpectedDate, period.ShiftTime(date))

	require.Equal(t, expectedDuration, period.RelativeDuration(date))
	require.NotEqual(t, unexpectedDuration, period.RelativeDuration(date))

	period, found, err := Parse("365d24h")
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

func TestCorretnessNegative(t *testing.T) {
	period := Period{
		negative: true,
		years:    1,
	}

	date := time.Date(2024, time.April, 1, 0, 0, 0, 0, time.UTC)

	expectedDate := time.Date(2023, time.April, 1, 0, 0, 0, 0, time.UTC)
	expectedDuration := date.AddDate(-1, 0, 0).Sub(date)

	unexpectedDuration := -365 * 24 * time.Hour
	unexpectedDate := date.Add(unexpectedDuration)

	require.Equal(t, expectedDate, period.ShiftTime(date))
	require.NotEqual(t, unexpectedDate, period.ShiftTime(date))

	require.Equal(t, expectedDuration, period.RelativeDuration(date))
	require.NotEqual(t, unexpectedDuration, period.RelativeDuration(date))

	period, found, err := Parse("-365d24h")
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

	if period.negative {
		require.Equal(t, duration, -period.duration)
	} else {
		require.Equal(t, duration, period.duration)
	}

	require.Equal(t, duration.String(), period.String())
}
