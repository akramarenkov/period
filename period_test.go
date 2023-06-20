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

	negative, next, err = isNegative([]rune("-"))
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
	unit, found, next := findUnit([]rune("y"))
	require.Equal(t, unitYear, unit)
	require.Equal(t, true, found)
	require.Equal(t, 1, next)

	unit, found, next = findUnit([]rune("year"))
	require.Equal(t, unitYear, unit)
	require.Equal(t, true, found)
	require.Equal(t, 4, next)

	unit, found, next = findUnit([]rune("years"))

	require.Equal(t, unitYear, unit)
	require.Equal(t, true, found)
	require.Equal(t, 5, next)

	unit, found, next = findUnit([]rune("mo"))
	require.Equal(t, unitMonth, unit)
	require.Equal(t, true, found)
	require.Equal(t, 2, next)

	unit, found, next = findUnit([]rune("month"))
	require.Equal(t, unitMonth, unit)
	require.Equal(t, true, found)
	require.Equal(t, 5, next)

	unit, found, next = findUnit([]rune("months"))
	require.Equal(t, unitMonth, unit)
	require.Equal(t, true, found)
	require.Equal(t, 6, next)

	unit, found, next = findUnit([]rune("d"))
	require.Equal(t, unitDay, unit)
	require.Equal(t, true, found)
	require.Equal(t, 1, next)

	unit, found, next = findUnit([]rune("day"))
	require.Equal(t, unitDay, unit)
	require.Equal(t, true, found)
	require.Equal(t, 3, next)

	unit, found, next = findUnit([]rune("days"))
	require.Equal(t, unitDay, unit)
	require.Equal(t, true, found)
	require.Equal(t, 4, next)

	unit, found, next = findUnit([]rune("h"))
	require.Equal(t, unitHour, unit)
	require.Equal(t, true, found)
	require.Equal(t, 1, next)

	unit, found, next = findUnit([]rune("hour"))
	require.Equal(t, unitHour, unit)
	require.Equal(t, true, found)
	require.Equal(t, 4, next)

	unit, found, next = findUnit([]rune("hours"))
	require.Equal(t, unitHour, unit)
	require.Equal(t, true, found)
	require.Equal(t, 5, next)

	unit, found, next = findUnit([]rune("m"))
	require.Equal(t, unitMinute, unit)
	require.Equal(t, true, found)
	require.Equal(t, 1, next)

	unit, found, next = findUnit([]rune("minute"))
	require.Equal(t, unitMinute, unit)
	require.Equal(t, true, found)
	require.Equal(t, 6, next)

	unit, found, next = findUnit([]rune("minutes"))
	require.Equal(t, unitMinute, unit)
	require.Equal(t, true, found)
	require.Equal(t, 7, next)

	unit, found, next = findUnit([]rune("s"))

	require.Equal(t, unitSecond, unit)
	require.Equal(t, true, found)
	require.Equal(t, 1, next)

	unit, found, next = findUnit([]rune("second"))
	require.Equal(t, unitSecond, unit)
	require.Equal(t, true, found)
	require.Equal(t, 6, next)

	unit, found, next = findUnit([]rune("seconds"))
	require.Equal(t, unitSecond, unit)
	require.Equal(t, true, found)
	require.Equal(t, 7, next)

	unit, found, next = findUnit([]rune("ms"))
	require.Equal(t, unitMillisecond, unit)
	require.Equal(t, true, found)
	require.Equal(t, 2, next)

	unit, found, next = findUnit([]rune("millisecond"))
	require.Equal(t, unitMillisecond, unit)
	require.Equal(t, true, found)
	require.Equal(t, 11, next)

	unit, found, next = findUnit([]rune("milliseconds"))
	require.Equal(t, unitMillisecond, unit)
	require.Equal(t, true, found)
	require.Equal(t, 12, next)

	unit, found, next = findUnit([]rune("us"))
	require.Equal(t, unitMicrosecond, unit)
	require.Equal(t, true, found)
	require.Equal(t, 2, next)

	unit, found, next = findUnit([]rune("µs"))
	require.Equal(t, unitMicrosecond, unit)
	require.Equal(t, true, found)
	require.Equal(t, 2, next)

	unit, found, next = findUnit([]rune("microsecond"))
	require.Equal(t, unitMicrosecond, unit)
	require.Equal(t, true, found)
	require.Equal(t, 11, next)

	unit, found, next = findUnit([]rune("microseconds"))
	require.Equal(t, unitMicrosecond, unit)
	require.Equal(t, true, found)
	require.Equal(t, 12, next)

	unit, found, next = findUnit([]rune("ns"))
	require.Equal(t, unitNanosecond, unit)
	require.Equal(t, true, found)
	require.Equal(t, 2, next)

	unit, found, next = findUnit([]rune("nanosecond"))
	require.Equal(t, unitNanosecond, unit)
	require.Equal(t, true, found)
	require.Equal(t, 10, next)

	unit, found, next = findUnit([]rune("nanoseconds"))
	require.Equal(t, unitNanosecond, unit)
	require.Equal(t, true, found)
	require.Equal(t, 11, next)

	unit, found, next = findUnit([]rune("y "))
	require.Equal(t, unitYear, unit)
	require.Equal(t, true, found)
	require.Equal(t, 1, next)

	unit, found, next = findUnit([]rune("yea"))
	require.Equal(t, unitUnknown, unit)
	require.Equal(t, false, found)
	require.Equal(t, 0, next)

	unit, found, next = findUnit([]rune("yea "))
	require.Equal(t, unitUnknown, unit)
	require.Equal(t, false, found)
	require.Equal(t, 0, next)
}

func TestFindNamedNumber(t *testing.T) {
	number, next, found, unit, err := findNamedNumber([]rune("10d"))
	require.NoError(t, err)
	require.Equal(t, []rune("10"), number)
	require.Equal(t, 3, next)
	require.Equal(t, true, found)
	require.Equal(t, unitDay, unit)

	number, next, found, unit, err = findNamedNumber([]rune("   10d"))
	require.NoError(t, err)
	require.Equal(t, []rune("10"), number)
	require.Equal(t, 6, next)
	require.Equal(t, true, found)
	require.Equal(t, unitDay, unit)

	number, next, found, unit, err = findNamedNumber([]rune("10d2m"))
	require.NoError(t, err)
	require.Equal(t, []rune("10"), number)
	require.Equal(t, 3, next)
	require.Equal(t, true, found)
	require.Equal(t, unitDay, unit)

	number, next, found, unit, err = findNamedNumber([]rune("   10d2m"))
	require.NoError(t, err)
	require.Equal(t, []rune("10"), number)
	require.Equal(t, 6, next)
	require.Equal(t, true, found)
	require.Equal(t, unitDay, unit)

	number, next, found, unit, err = findNamedNumber([]rune("1.10d"))
	require.NoError(t, err)
	require.Equal(t, []rune("1.10"), number)
	require.Equal(t, 5, next)
	require.Equal(t, true, found)
	require.Equal(t, unitDay, unit)

	number, next, found, unit, err = findNamedNumber([]rune("   1.10d"))
	require.NoError(t, err)
	require.Equal(t, []rune("1.10"), number)
	require.Equal(t, 8, next)
	require.Equal(t, true, found)
	require.Equal(t, unitDay, unit)

	number, next, found, unit, err = findNamedNumber([]rune(".10d"))
	require.NoError(t, err)
	require.Equal(t, []rune(".10"), number)
	require.Equal(t, 4, next)
	require.Equal(t, true, found)
	require.Equal(t, unitDay, unit)

	number, next, found, unit, err = findNamedNumber([]rune("   .10d"))
	require.NoError(t, err)
	require.Equal(t, []rune(".10"), number)
	require.Equal(t, 7, next)
	require.Equal(t, true, found)
	require.Equal(t, unitDay, unit)

	number, next, found, unit, err = findNamedNumber([]rune(""))
	require.NoError(t, err)
	require.Equal(t, []rune(nil), number)
	require.Equal(t, 0, next)
	require.Equal(t, false, found)
	require.Equal(t, unitUnknown, unit)

	number, next, found, unit, err = findNamedNumber([]rune("  "))
	require.NoError(t, err)
	require.Equal(t, []rune(nil), number)
	require.Equal(t, 0, next)
	require.Equal(t, false, found)
	require.Equal(t, unitUnknown, unit)

	number, next, found, unit, err = findNamedNumber([]rune("d"))
	require.Error(t, err)
	require.Equal(t, []rune(nil), number)
	require.Equal(t, 0, next)
	require.Equal(t, false, found)
	require.Equal(t, unitUnknown, unit)

	number, next, found, unit, err = findNamedNumber([]rune("  d"))
	require.Error(t, err)
	require.Equal(t, []rune(nil), number)
	require.Equal(t, 0, next)
	require.Equal(t, false, found)
	require.Equal(t, unitUnknown, unit)

	number, next, found, unit, err = findNamedNumber([]rune("-10d"))
	require.Error(t, err)
	require.Equal(t, []rune(nil), number)
	require.Equal(t, 0, next)
	require.Equal(t, false, found)
	require.Equal(t, unitUnknown, unit)

	number, next, found, unit, err = findNamedNumber([]rune("10"))
	require.Error(t, err)
	require.Equal(t, []rune(nil), number)
	require.Equal(t, 0, next)
	require.Equal(t, false, found)
	require.Equal(t, unitUnknown, unit)

	number, next, found, unit, err = findNamedNumber([]rune("   1 0d"))
	require.Error(t, err)
	require.Equal(t, []rune(nil), number)
	require.Equal(t, 0, next)
	require.Equal(t, false, found)
	require.Equal(t, unitUnknown, unit)

	number, next, found, unit, err = findNamedNumber([]rune("1..10d"))
	require.Error(t, err)
	require.Equal(t, []rune(nil), number)
	require.Equal(t, 0, next)
	require.Equal(t, false, found)
	require.Equal(t, unitUnknown, unit)

	number, next, found, unit, err = findNamedNumber([]rune("   1..10d"))
	require.Error(t, err)
	require.Equal(t, []rune(nil), number)
	require.Equal(t, 0, next)
	require.Equal(t, false, found)
	require.Equal(t, unitUnknown, unit)

	number, next, found, unit, err = findNamedNumber([]rune("..10d"))
	require.Error(t, err)
	require.Equal(t, []rune(nil), number)
	require.Equal(t, 0, next)
	require.Equal(t, false, found)
	require.Equal(t, unitUnknown, unit)

	number, next, found, unit, err = findNamedNumber([]rune("   ..10d"))
	require.Error(t, err)
	require.Equal(t, []rune(nil), number)
	require.Equal(t, 0, next)
	require.Equal(t, false, found)
	require.Equal(t, unitUnknown, unit)
}

func TestParse(t *testing.T) {
	period, found, err := Parse("10d")
	require.NoError(t, err)
	require.Equal(t, Period{days: 10}, period)
	require.Equal(t, true, found)

	period, found, err = Parse("   10d")
	require.NoError(t, err)
	require.Equal(t, Period{days: 10}, period)
	require.Equal(t, true, found)

	period, found, err = Parse("-10d")
	require.NoError(t, err)
	require.Equal(t, Period{negative: true, days: 10}, period)
	require.Equal(t, true, found)

	period, found, err = Parse("   -   10d")
	require.NoError(t, err)
	require.Equal(t, Period{negative: true, days: 10}, period)
	require.Equal(t, true, found)

	expected := Period{
		negative: true,

		years:  2,
		months: 3,
		days:   10,

		duration: 86398010030010,
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
	require.Equal(t, Period{}, period)
	require.Equal(t, false, found)

	period, found, err = Parse("   ")
	require.NoError(t, err)
	require.Equal(t, Period{}, period)
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

func TestString(t *testing.T) {
	source := "-2y3mo10d23h59m58s10ms30us10ns"

	period, found, err := Parse(source)
	require.NoError(t, err)
	require.Equal(t, true, found)

	require.Equal(t, source, period.String())
}
