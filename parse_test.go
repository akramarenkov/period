package period

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestIsSpecialZero(t *testing.T) {
	require.Equal(t, true, isSpecialZero([]rune("0")))
	require.Equal(t, false, isSpecialZero([]rune("1")))
	require.Equal(t, false, isSpecialZero([]rune("")))
	require.Equal(t, false, isSpecialZero([]rune("00")))
}

func TestIsNegative(t *testing.T) {
	negative, next, err := isNegative(
		[]rune("-10"),
		defaultMinusSign,
		defaultPlusSign,
		defaultFractionalSeparator,
	)
	require.NoError(t, err)
	require.Equal(t, true, negative)
	require.Equal(t, 1, next)

	negative, next, err = isNegative(
		[]rune("+10"),
		defaultMinusSign,
		defaultPlusSign,
		defaultFractionalSeparator,
	)
	require.NoError(t, err)
	require.Equal(t, false, negative)
	require.Equal(t, 1, next)

	negative, next, err = isNegative([]rune("   -10"),
		defaultMinusSign,
		defaultPlusSign,
		defaultFractionalSeparator,
	)
	require.NoError(t, err)
	require.Equal(t, true, negative)
	require.Equal(t, 4, next)

	negative, next, err = isNegative([]rune("- 10"),
		defaultMinusSign,
		defaultPlusSign,
		defaultFractionalSeparator,
	)
	require.NoError(t, err)
	require.Equal(t, true, negative)
	require.Equal(t, 2, next)

	negative, next, err = isNegative([]rune("   -   10"),
		defaultMinusSign,
		defaultPlusSign,
		defaultFractionalSeparator,
	)
	require.NoError(t, err)
	require.Equal(t, true, negative)
	require.Equal(t, 7, next)

	negative, next, err = isNegative([]rune(""),
		defaultMinusSign,
		defaultPlusSign,
		defaultFractionalSeparator,
	)
	require.NoError(t, err)
	require.Equal(t, false, negative)
	require.Equal(t, 0, next)

	negative, next, err = isNegative([]rune("   "),
		defaultMinusSign,
		defaultPlusSign,
		defaultFractionalSeparator,
	)
	require.NoError(t, err)
	require.Equal(t, false, negative)
	require.Equal(t, 0, next)

	negative, next, err = isNegative([]rune("10"),
		defaultMinusSign,
		defaultPlusSign,
		defaultFractionalSeparator,
	)
	require.NoError(t, err)
	require.Equal(t, false, negative)
	require.Equal(t, 0, next)

	negative, next, err = isNegative([]rune("   10"),
		defaultMinusSign,
		defaultPlusSign,
		defaultFractionalSeparator,
	)
	require.NoError(t, err)
	require.Equal(t, false, negative)
	require.Equal(t, 0, next)

	negative, next, err = isNegative([]rune(".0"),
		defaultMinusSign,
		defaultPlusSign,
		defaultFractionalSeparator,
	)
	require.NoError(t, err)
	require.Equal(t, false, negative)
	require.Equal(t, 0, next)

	negative, next, err = isNegative([]rune("-.0"),
		defaultMinusSign,
		defaultPlusSign,
		defaultFractionalSeparator,
	)
	require.NoError(t, err)
	require.Equal(t, true, negative)
	require.Equal(t, 1, next)

	negative, next, err = isNegative([]rune("+.0"),
		defaultMinusSign,
		defaultPlusSign,
		defaultFractionalSeparator,
	)
	require.NoError(t, err)
	require.Equal(t, false, negative)
	require.Equal(t, 1, next)
}

func TestIsNegativeRequireError(t *testing.T) {
	inputs := []string{
		"-",
		" - ",
		"d",
		"-d",
		"   d",
		"   -d",
		"   -   d",
		"--10",
		"   -   -   10",
		"-+10",
		"   -   +   10",
		"++10",
		"   +   +   10",
	}

	for _, input := range inputs {
		negative, next, err := isNegative([]rune(input),
			defaultMinusSign,
			defaultPlusSign,
			defaultFractionalSeparator,
		)
		require.Error(t, err)
		require.Equal(t, false, negative)
		require.Equal(t, 0, next)
	}
}

func TestFindUnit(t *testing.T) {
	unit, found, next := findUnit(
		[]rune("y"),
		defaultFractionalSeparator,
		defaultUnits,
	)
	require.Equal(t, UnitYear, unit)
	require.Equal(t, true, found)
	require.Equal(t, 1, next)

	unit, found, next = findUnit(
		[]rune("y   "),
		defaultFractionalSeparator,
		defaultUnits,
	)
	require.Equal(t, UnitYear, unit)
	require.Equal(t, true, found)
	require.Equal(t, 1, next)

	unit, found, next = findUnit(
		[]rune("mo"),
		defaultFractionalSeparator,
		defaultUnits,
	)
	require.Equal(t, UnitMonth, unit)
	require.Equal(t, true, found)
	require.Equal(t, 2, next)

	unit, found, next = findUnit(
		[]rune("d"),
		defaultFractionalSeparator,
		defaultUnits,
	)
	require.Equal(t, UnitDay, unit)
	require.Equal(t, true, found)
	require.Equal(t, 1, next)

	unit, found, next = findUnit(
		[]rune("h"),
		defaultFractionalSeparator,
		defaultUnits,
	)
	require.Equal(t, UnitHour, unit)
	require.Equal(t, true, found)
	require.Equal(t, 1, next)

	unit, found, next = findUnit(
		[]rune("m"),
		defaultFractionalSeparator,
		defaultUnits,
	)
	require.Equal(t, UnitMinute, unit)
	require.Equal(t, true, found)
	require.Equal(t, 1, next)

	unit, found, next = findUnit(
		[]rune("s"),
		defaultFractionalSeparator,
		defaultUnits,
	)
	require.Equal(t, UnitSecond, unit)
	require.Equal(t, true, found)
	require.Equal(t, 1, next)

	unit, found, next = findUnit(
		[]rune("ms"),
		defaultFractionalSeparator,
		defaultUnits,
	)
	require.Equal(t, UnitMillisecond, unit)
	require.Equal(t, true, found)
	require.Equal(t, 2, next)

	unit, found, next = findUnit(
		[]rune("us"),
		defaultFractionalSeparator,
		defaultUnits,
	)
	require.Equal(t, UnitMicrosecond, unit)
	require.Equal(t, true, found)
	require.Equal(t, 2, next)

	unit, found, next = findUnit(
		[]rune("µs"),
		defaultFractionalSeparator,
		defaultUnits,
	)
	require.Equal(t, UnitMicrosecond, unit)
	require.Equal(t, true, found)
	require.Equal(t, 2, next)

	unit, found, next = findUnit(
		[]rune("μs"),
		defaultFractionalSeparator,
		defaultUnits,
	)
	require.Equal(t, UnitMicrosecond, unit)
	require.Equal(t, true, found)
	require.Equal(t, 2, next)

	unit, found, next = findUnit(
		[]rune("ns"),
		defaultFractionalSeparator,
		defaultUnits,
	)
	require.Equal(t, UnitNanosecond, unit)
	require.Equal(t, true, found)
	require.Equal(t, 2, next)
}

func TestFindUnitNotFound(t *testing.T) {
	unit, found, next := findUnit(
		[]rune("u"),
		defaultFractionalSeparator,
		defaultUnits,
	)
	require.Equal(t, UnitUnknown, unit)
	require.Equal(t, false, found)
	require.Equal(t, 0, next)

	unit, found, next = findUnit(
		[]rune("n "),
		defaultFractionalSeparator,
		defaultUnits,
	)
	require.Equal(t, UnitUnknown, unit)
	require.Equal(t, false, found)
	require.Equal(t, 0, next)
}

func TestFindNamedNumber(t *testing.T) {
	number, next, found, unit, err := findNamedNumber(
		[]rune("10d"),
		defaultUnits,
		defaultFractionalSeparator,
	)
	require.NoError(t, err)
	require.Equal(t, []rune("10"), number)
	require.Equal(t, 3, next)
	require.Equal(t, true, found)
	require.Equal(t, UnitDay, unit)

	number, next, found, unit, err = findNamedNumber(
		[]rune("   10d"),
		defaultUnits,
		defaultFractionalSeparator,
	)
	require.NoError(t, err)
	require.Equal(t, []rune("10"), number)
	require.Equal(t, 6, next)
	require.Equal(t, true, found)
	require.Equal(t, UnitDay, unit)

	number, next, found, unit, err = findNamedNumber(
		[]rune("10d2m"),
		defaultUnits,
		defaultFractionalSeparator,
	)
	require.NoError(t, err)
	require.Equal(t, []rune("10"), number)
	require.Equal(t, 3, next)
	require.Equal(t, true, found)
	require.Equal(t, UnitDay, unit)

	number, next, found, unit, err = findNamedNumber(
		[]rune("   10d2m"),
		defaultUnits,
		defaultFractionalSeparator,
	)
	require.NoError(t, err)
	require.Equal(t, []rune("10"), number)
	require.Equal(t, 6, next)
	require.Equal(t, true, found)
	require.Equal(t, UnitDay, unit)

	number, next, found, unit, err = findNamedNumber(
		[]rune("1.10d"),
		defaultUnits,
		defaultFractionalSeparator,
	)
	require.NoError(t, err)
	require.Equal(t, []rune("1.10"), number)
	require.Equal(t, 5, next)
	require.Equal(t, true, found)
	require.Equal(t, UnitDay, unit)

	number, next, found, unit, err = findNamedNumber(
		[]rune("   1.10d"),
		defaultUnits,
		defaultFractionalSeparator,
	)
	require.NoError(t, err)
	require.Equal(t, []rune("1.10"), number)
	require.Equal(t, 8, next)
	require.Equal(t, true, found)
	require.Equal(t, UnitDay, unit)

	number, next, found, unit, err = findNamedNumber(
		[]rune(".10d"),
		defaultUnits,
		defaultFractionalSeparator,
	)
	require.NoError(t, err)
	require.Equal(t, []rune(".10"), number)
	require.Equal(t, 4, next)
	require.Equal(t, true, found)
	require.Equal(t, UnitDay, unit)

	number, next, found, unit, err = findNamedNumber(
		[]rune("   .10d"),
		defaultUnits,
		defaultFractionalSeparator,
	)
	require.NoError(t, err)
	require.Equal(t, []rune(".10"), number)
	require.Equal(t, 7, next)
	require.Equal(t, true, found)
	require.Equal(t, UnitDay, unit)

	number, next, found, unit, err = findNamedNumber(
		[]rune("   .d"),
		defaultUnits,
		defaultFractionalSeparator,
	)
	require.NoError(t, err)
	require.Equal(t, []rune("."), number)
	require.Equal(t, 5, next)
	require.Equal(t, true, found)
	require.Equal(t, UnitDay, unit)

	number, next, found, unit, err = findNamedNumber(
		[]rune(""),
		defaultUnits,
		defaultFractionalSeparator,
	)
	require.NoError(t, err)
	require.Equal(t, []rune(nil), number)
	require.Equal(t, 0, next)
	require.Equal(t, false, found)
	require.Equal(t, UnitUnknown, unit)

	number, next, found, unit, err = findNamedNumber(
		[]rune("  "),
		defaultUnits,
		defaultFractionalSeparator,
	)
	require.NoError(t, err)
	require.Equal(t, []rune(nil), number)
	require.Equal(t, 0, next)
	require.Equal(t, false, found)
	require.Equal(t, UnitUnknown, unit)
}

func TestFindNamedNumberRequireError(t *testing.T) {
	inputs := []string{
		"d",
		"  d",
		"-10d",
		"10",
		"   1 0d",
		"1..10d",
		"   1..10d",
		"..10d",
		"   ..10d",
	}

	for _, input := range inputs {
		number, next, found, unit, err := findNamedNumber(
			[]rune(input),
			defaultUnits,
			defaultFractionalSeparator,
		)
		require.Error(t, err)
		require.Equal(t, []rune(nil), number)
		require.Equal(t, 0, next)
		require.Equal(t, false, found)
		require.Equal(t, UnitUnknown, unit)
	}
}

func TestSplitNumber(t *testing.T) {
	integer, fractional, err := splitNumber([]rune("1.2"), defaultFractionalSeparator)
	require.NoError(t, err)
	require.Equal(t, []rune("1"), integer)
	require.Equal(t, []rune("2"), fractional)

	integer, fractional, err = splitNumber([]rune(".2"), defaultFractionalSeparator)
	require.NoError(t, err)
	require.Equal(t, []rune(""), integer)
	require.Equal(t, []rune("2"), fractional)

	integer, fractional, err = splitNumber([]rune("1."), defaultFractionalSeparator)
	require.NoError(t, err)
	require.Equal(t, []rune("1"), integer)
	require.Equal(t, []rune(""), fractional)

	integer, fractional, err = splitNumber([]rune("."), defaultFractionalSeparator)
	require.NoError(t, err)
	require.Equal(t, []rune(""), integer)
	require.Equal(t, []rune(""), fractional)

	integer, fractional, err = splitNumber([]rune("1"), defaultFractionalSeparator)
	require.NoError(t, err)
	require.Equal(t, []rune("1"), integer)
	require.Equal(t, []rune(nil), fractional)

	integer, fractional, err = splitNumber([]rune("12"), defaultFractionalSeparator)
	require.NoError(t, err)
	require.Equal(t, []rune("12"), integer)
	require.Equal(t, []rune(nil), fractional)

	integer, fractional, err = splitNumber([]rune(""), defaultFractionalSeparator)
	require.NoError(t, err)
	require.Equal(t, []rune(""), integer)
	require.Equal(t, []rune(nil), fractional)
}

func TestSplitNumberRequireError(t *testing.T) {
	inputs := []string{
		" ",
		"1. 1",
		"1 .1",
		" 1.1",
		"1..1",
	}

	for _, input := range inputs {
		integer, fractional, err := splitNumber([]rune(input), defaultFractionalSeparator)
		require.Error(t, err)
		require.Equal(t, []rune(nil), integer)
		require.Equal(t, []rune(nil), fractional)
	}
}

func TestParseDuration(t *testing.T) {
	duration, err := parseDuration(
		namedNumber{[]rune("2.5"), UnitHour},
		defaultNumberBase,
		defaultFractionalSeparator,
		false,
	)
	require.NoError(t, err)
	require.Equal(t, 2*time.Hour+30*time.Minute, duration)

	duration, err = parseDuration(
		namedNumber{[]rune("2.5"), UnitHour},
		defaultNumberBase,
		defaultFractionalSeparator,
		true,
	)
	require.NoError(t, err)
	require.Equal(t, 2*time.Hour+30*time.Minute, duration)
}

func TestParseDurationRequireError(t *testing.T) {
	duration, err := parseDuration(
		namedNumber{[]rune("2,5"), UnitHour},
		defaultNumberBase,
		defaultFractionalSeparator,
		false,
	)
	require.Error(t, err)
	require.Equal(t, time.Duration(0), duration)

	duration, err = parseDuration(
		namedNumber{[]rune("2,5"), UnitHour},
		defaultNumberBase,
		defaultFractionalSeparator,
		true,
	)
	require.Error(t, err)
	require.Equal(t, time.Duration(0), duration)

	duration, err = parseDuration(
		namedNumber{[]rune("2.5"), UnitUnknown},
		defaultNumberBase,
		defaultFractionalSeparator,
		false,
	)
	require.Error(t, err)
	require.Equal(t, time.Duration(0), duration)

	duration, err = parseDuration(
		namedNumber{[]rune("2.5"), UnitUnknown},
		defaultNumberBase,
		defaultFractionalSeparator,
		true,
	)
	require.Error(t, err)
	require.Equal(t, time.Duration(0), duration)
}

func TestParseFractionalDurationUnexpectedUnit(t *testing.T) {
	duration, err := parseFractionalDuration(
		[]rune(""),
		defaultNumberBase,
		UnitUnknown,
	)
	require.Error(t, err)
	require.Equal(t, time.Duration(0), duration)
}
