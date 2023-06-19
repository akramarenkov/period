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
	unit, found, next, err := findUnit([]rune("y"))
	require.NoError(t, err)
	require.Equal(t, UnitYear, unit)
	require.Equal(t, true, found)
	require.Equal(t, 1, next)

	unit, found, next, err = findUnit([]rune("year"))
	require.NoError(t, err)
	require.Equal(t, UnitYear, unit)
	require.Equal(t, true, found)
	require.Equal(t, 4, next)

	unit, found, next, err = findUnit([]rune("years"))
	require.NoError(t, err)
	require.Equal(t, UnitYear, unit)
	require.Equal(t, true, found)
	require.Equal(t, 5, next)

	unit, found, next, err = findUnit([]rune("mo"))
	require.NoError(t, err)
	require.Equal(t, UnitMonth, unit)
	require.Equal(t, true, found)
	require.Equal(t, 2, next)

	unit, found, next, err = findUnit([]rune("month"))
	require.NoError(t, err)
	require.Equal(t, UnitMonth, unit)
	require.Equal(t, true, found)
	require.Equal(t, 5, next)

	unit, found, next, err = findUnit([]rune("months"))
	require.NoError(t, err)
	require.Equal(t, UnitMonth, unit)
	require.Equal(t, true, found)
	require.Equal(t, 6, next)

	unit, found, next, err = findUnit([]rune("d"))
	require.NoError(t, err)
	require.Equal(t, UnitDay, unit)
	require.Equal(t, true, found)
	require.Equal(t, 1, next)

	unit, found, next, err = findUnit([]rune("day"))
	require.NoError(t, err)
	require.Equal(t, UnitDay, unit)
	require.Equal(t, true, found)
	require.Equal(t, 3, next)

	unit, found, next, err = findUnit([]rune("days"))
	require.NoError(t, err)
	require.Equal(t, UnitDay, unit)
	require.Equal(t, true, found)
	require.Equal(t, 4, next)

	unit, found, next, err = findUnit([]rune("h"))
	require.NoError(t, err)
	require.Equal(t, UnitHour, unit)
	require.Equal(t, true, found)
	require.Equal(t, 1, next)

	unit, found, next, err = findUnit([]rune("hour"))
	require.NoError(t, err)
	require.Equal(t, UnitHour, unit)
	require.Equal(t, true, found)
	require.Equal(t, 4, next)

	unit, found, next, err = findUnit([]rune("hours"))
	require.NoError(t, err)
	require.Equal(t, UnitHour, unit)
	require.Equal(t, true, found)
	require.Equal(t, 5, next)

	unit, found, next, err = findUnit([]rune("m"))
	require.NoError(t, err)
	require.Equal(t, UnitMinute, unit)
	require.Equal(t, true, found)
	require.Equal(t, 1, next)

	unit, found, next, err = findUnit([]rune("minute"))
	require.NoError(t, err)
	require.Equal(t, UnitMinute, unit)
	require.Equal(t, true, found)
	require.Equal(t, 6, next)

	unit, found, next, err = findUnit([]rune("minutes"))
	require.NoError(t, err)
	require.Equal(t, UnitMinute, unit)
	require.Equal(t, true, found)
	require.Equal(t, 7, next)

	unit, found, next, err = findUnit([]rune("s"))
	require.NoError(t, err)
	require.Equal(t, UnitSecond, unit)
	require.Equal(t, true, found)
	require.Equal(t, 1, next)

	unit, found, next, err = findUnit([]rune("second"))
	require.NoError(t, err)
	require.Equal(t, UnitSecond, unit)
	require.Equal(t, true, found)
	require.Equal(t, 6, next)

	unit, found, next, err = findUnit([]rune("seconds"))
	require.NoError(t, err)
	require.Equal(t, UnitSecond, unit)
	require.Equal(t, true, found)
	require.Equal(t, 7, next)

	unit, found, next, err = findUnit([]rune("ms"))
	require.NoError(t, err)
	require.Equal(t, UnitMillisecond, unit)
	require.Equal(t, true, found)
	require.Equal(t, 2, next)

	unit, found, next, err = findUnit([]rune("millisecond"))
	require.NoError(t, err)
	require.Equal(t, UnitMillisecond, unit)
	require.Equal(t, true, found)
	require.Equal(t, 11, next)

	unit, found, next, err = findUnit([]rune("milliseconds"))
	require.NoError(t, err)
	require.Equal(t, UnitMillisecond, unit)
	require.Equal(t, true, found)
	require.Equal(t, 12, next)

	unit, found, next, err = findUnit([]rune("us"))
	require.NoError(t, err)
	require.Equal(t, UnitMicrosecond, unit)
	require.Equal(t, true, found)
	require.Equal(t, 2, next)

	unit, found, next, err = findUnit([]rune("µs"))
	require.NoError(t, err)
	require.Equal(t, UnitMicrosecond, unit)
	require.Equal(t, true, found)
	require.Equal(t, 2, next)

	unit, found, next, err = findUnit([]rune("microsecond"))
	require.NoError(t, err)
	require.Equal(t, UnitMicrosecond, unit)
	require.Equal(t, true, found)
	require.Equal(t, 11, next)

	unit, found, next, err = findUnit([]rune("microseconds"))
	require.NoError(t, err)
	require.Equal(t, UnitMicrosecond, unit)
	require.Equal(t, true, found)
	require.Equal(t, 12, next)

	unit, found, next, err = findUnit([]rune("ns"))
	require.NoError(t, err)
	require.Equal(t, UnitNanosecond, unit)
	require.Equal(t, true, found)
	require.Equal(t, 2, next)

	unit, found, next, err = findUnit([]rune("nanosecond"))
	require.NoError(t, err)
	require.Equal(t, UnitNanosecond, unit)
	require.Equal(t, true, found)
	require.Equal(t, 10, next)

	unit, found, next, err = findUnit([]rune("nanoseconds"))
	require.NoError(t, err)
	require.Equal(t, UnitNanosecond, unit)
	require.Equal(t, true, found)
	require.Equal(t, 11, next)

	unit, found, next, err = findUnit([]rune("yea"))
	require.NoError(t, err)
	require.Equal(t, UnitUnknown, unit)
	require.Equal(t, false, found)
	require.Equal(t, 0, next)
}

func TestFindOneNamedNumber(t *testing.T) {
	number, next, found, unit, err := findOneNamedNumber([]rune("10d"))
	require.NoError(t, err)
	require.Equal(t, []rune("10"), number)
	require.Equal(t, 3, next)
	require.Equal(t, true, found)
	require.Equal(t, UnitDay, unit)

	number, next, found, unit, err = findOneNamedNumber([]rune("   10d"))
	require.NoError(t, err)
	require.Equal(t, []rune("10"), number)
	require.Equal(t, 6, next)
	require.Equal(t, true, found)
	require.Equal(t, UnitDay, unit)

	number, next, found, unit, err = findOneNamedNumber([]rune("10d2m"))
	require.NoError(t, err)
	require.Equal(t, []rune("10"), number)
	require.Equal(t, 3, next)
	require.Equal(t, true, found)
	require.Equal(t, UnitDay, unit)

	number, next, found, unit, err = findOneNamedNumber([]rune("   10d2m"))
	require.NoError(t, err)
	require.Equal(t, []rune("10"), number)
	require.Equal(t, 6, next)
	require.Equal(t, true, found)
	require.Equal(t, UnitDay, unit)

	number, next, found, unit, err = findOneNamedNumber([]rune(""))
	require.NoError(t, err)
	require.Equal(t, []rune(nil), number)
	require.Equal(t, 0, next)
	require.Equal(t, false, found)
	require.Equal(t, UnitUnknown, unit)

	number, next, found, unit, err = findOneNamedNumber([]rune("  "))
	require.NoError(t, err)
	require.Equal(t, []rune(nil), number)
	require.Equal(t, 0, next)
	require.Equal(t, false, found)
	require.Equal(t, UnitUnknown, unit)

	number, next, found, unit, err = findOneNamedNumber([]rune(" 1 0d"))
	require.NoError(t, err)
	require.Equal(t, []rune("1"), number)
	require.Equal(t, 2, next)
	require.Equal(t, true, found)
	require.Equal(t, UnitNanosecond, unit)

	number, next, found, unit, err = findOneNamedNumber([]rune("d"))
	require.Error(t, err)
	require.Equal(t, []rune(nil), number)
	require.Equal(t, 0, next)
	require.Equal(t, false, found)
	require.Equal(t, UnitUnknown, unit)

	number, next, found, unit, err = findOneNamedNumber([]rune("  d"))
	require.Error(t, err)
	require.Equal(t, []rune(nil), number)
	require.Equal(t, 0, next)
	require.Equal(t, false, found)
	require.Equal(t, UnitUnknown, unit)

	number, next, found, unit, err = findOneNamedNumber([]rune("-10d"))
	require.Error(t, err)
	require.Equal(t, []rune(nil), number)
	require.Equal(t, 0, next)
	require.Equal(t, false, found)
	require.Equal(t, UnitUnknown, unit)
}

func TestParse(t *testing.T) {
	period, found, err := Parse("10d")
	require.NoError(t, err)
	require.Equal(t, Period{Days: 10}, period)
	require.Equal(t, true, found)

	period, found, err = Parse("   10d")
	require.NoError(t, err)
	require.Equal(t, Period{Days: 10}, period)
	require.Equal(t, true, found)

	period, found, err = Parse("-10d")
	require.NoError(t, err)
	require.Equal(t, Period{Days: -10}, period)
	require.Equal(t, true, found)

	period, found, err = Parse("   -10d")
	require.NoError(t, err)
	require.Equal(t, Period{Days: -10}, period)
	require.Equal(t, true, found)

	period, found, err = Parse("10d2y")
	require.NoError(t, err)
	require.Equal(t, Period{Days: 10, Years: 2}, period)
	require.Equal(t, true, found)

	period, found, err = Parse("   10d   2y")
	require.NoError(t, err)
	require.Equal(t, Period{Days: 10, Years: 2}, period)
	require.Equal(t, true, found)

	period, found, err = Parse("   -  10d   2y")
	require.NoError(t, err)
	require.Equal(t, Period{Days: -10, Years: -2}, period)
	require.Equal(t, true, found)

	period, found, err = Parse("10d2y3mo")
	require.NoError(t, err)
	require.Equal(t, Period{Days: 10, Years: 2, Months: 3}, period)
	require.Equal(t, true, found)

	period, found, err = Parse("-10d2y3mo")
	require.NoError(t, err)
	require.Equal(t, Period{Days: -10, Years: -2, Months: -3}, period)
	require.Equal(t, true, found)

	period, found, err = Parse("    10d   2y  3mo")
	require.NoError(t, err)
	require.Equal(t, Period{Days: 10, Years: 2, Months: 3}, period)
	require.Equal(t, true, found)

	period, found, err = Parse("   -  10d   2y  3mo")
	require.NoError(t, err)
	require.Equal(t, Period{Days: -10, Years: -2, Months: -3}, period)
	require.Equal(t, true, found)

	period, found, err = Parse("-3mo2y10d")
	require.NoError(t, err)
	require.Equal(t, Period{Days: -10, Years: -2, Months: -3}, period)
	require.Equal(t, true, found)

	period, found, err = Parse("- 3mo 2y 10d")
	require.NoError(t, err)
	require.Equal(t, Period{Days: -10, Years: -2, Months: -3}, period)
	require.Equal(t, true, found)

	period, found, err = Parse(" - 3mo 2y 001d")
	require.NoError(t, err)
	require.Equal(t, Period{Days: -1, Years: -2, Months: -3}, period)
	require.Equal(t, true, found)

	period, found, err = Parse("10")
	require.NoError(t, err)
	require.Equal(t, Period{Duration: 10 * time.Nanosecond}, period)
	require.Equal(t, true, found)

	period, found, err = Parse("   10")
	require.NoError(t, err)
	require.Equal(t, Period{Duration: 10 * time.Nanosecond}, period)
	require.Equal(t, true, found)

	period, found, err = Parse("10d2")
	require.NoError(t, err)
	require.Equal(t, Period{Days: 10, Duration: 2 * time.Nanosecond}, period)
	require.Equal(t, true, found)

	period, found, err = Parse("   10d2")
	require.NoError(t, err)
	require.Equal(t, Period{Days: 10, Duration: 2 * time.Nanosecond}, period)
	require.Equal(t, true, found)

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

	period, found, err = Parse("10d2y3mo1y")
	require.Error(t, err)
	require.Equal(t, Period{}, period)
	require.Equal(t, false, found)
}
