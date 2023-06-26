package period

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestIsValidDefaultUnitsTable(t *testing.T) {
	require.NoError(t, IsValidUnitsTable(defaultUnits))
}

func TestIsValidUnitsTableInvalidUnit(t *testing.T) {
	units := UnitsTable{
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

	require.Error(t, IsValidUnitsTable(units))
}

func TestIsValidUnitsTableMissingUnit(t *testing.T) {
	units := UnitsTable{
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

	require.Error(t, IsValidUnitsTable(units))
}

func TestIsValidUnitsTableMissingUnitModifier(t *testing.T) {
	units := UnitsTable{
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

	require.Error(t, IsValidUnitsTable(units))
}

func TestIsValidUnitsTableEmptyUnitModifier(t *testing.T) {
	units := UnitsTable{
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

	require.Error(t, IsValidUnitsTable(units))
}

func TestIsValidUnitsTableModifierIsNotUnique(t *testing.T) {
	units := UnitsTable{
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

	require.Error(t, IsValidUnitsTable(units))
}

func TestGetDurationDimension(t *testing.T) {
	dimension, err := getDurationDimension(UnitHour)
	require.NoError(t, err)
	require.Equal(t, time.Hour, dimension)

	dimension, err = getDurationDimension(UnitMinute)
	require.NoError(t, err)
	require.Equal(t, time.Minute, dimension)

	dimension, err = getDurationDimension(UnitSecond)
	require.NoError(t, err)
	require.Equal(t, time.Second, dimension)

	dimension, err = getDurationDimension(UnitMillisecond)
	require.NoError(t, err)
	require.Equal(t, time.Millisecond, dimension)

	dimension, err = getDurationDimension(UnitMicrosecond)
	require.NoError(t, err)
	require.Equal(t, time.Microsecond, dimension)

	dimension, err = getDurationDimension(UnitNanosecond)
	require.NoError(t, err)
	require.Equal(t, time.Nanosecond, dimension)
}

func TestGetDurationDimensionRequireError(t *testing.T) {
	dimension, err := getDurationDimension(UnitYear)
	require.Error(t, err)
	require.Equal(t, time.Duration(0), dimension)
}
