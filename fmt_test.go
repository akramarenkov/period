package period

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestFormatFractional(t *testing.T) {
	formated, err := formatFractional(
		1,
		defaultNumberBase,
		defaultFormatFractionalSize,
		defaultFractionalSeparator,
	)
	require.NoError(t, err)
	require.Equal(t, ".000000001", formated)

	formated, err = formatFractional(
		10,
		defaultNumberBase,
		defaultFormatFractionalSize,
		defaultFractionalSeparator,
	)
	require.NoError(t, err)
	require.Equal(t, ".00000001", formated)

	formated, err = formatFractional(
		100,
		defaultNumberBase,
		defaultFormatFractionalSize,
		defaultFractionalSeparator,
	)
	require.NoError(t, err)
	require.Equal(t, ".0000001", formated)

	formated, err = formatFractional(
		1000,
		defaultNumberBase,
		defaultFormatFractionalSize,
		defaultFractionalSeparator,
	)
	require.NoError(t, err)
	require.Equal(t, ".000001", formated)

	formated, err = formatFractional(
		10000,
		defaultNumberBase,
		defaultFormatFractionalSize,
		defaultFractionalSeparator,
	)
	require.NoError(t, err)
	require.Equal(t, ".00001", formated)

	formated, err = formatFractional(
		100000,
		defaultNumberBase,
		defaultFormatFractionalSize,
		defaultFractionalSeparator,
	)
	require.NoError(t, err)
	require.Equal(t, ".0001", formated)

	formated, err = formatFractional(
		1000000,
		defaultNumberBase,
		defaultFormatFractionalSize,
		defaultFractionalSeparator,
	)
	require.NoError(t, err)
	require.Equal(t, ".001", formated)

	formated, err = formatFractional(
		10000000,
		defaultNumberBase,
		defaultFormatFractionalSize,
		defaultFractionalSeparator,
	)
	require.NoError(t, err)
	require.Equal(t, ".01", formated)

	formated, err = formatFractional(
		100000000,
		defaultNumberBase,
		defaultFormatFractionalSize,
		defaultFractionalSeparator,
	)
	require.NoError(t, err)
	require.Equal(t, ".1", formated)

	formated, err = formatFractional(
		100000000,
		defaultNumberBase,
		defaultFormatFractionalSize,
		defaultFractionalSeparator,
	)
	require.NoError(t, err)
	require.Equal(t, ".1", formated)

	formated, err = formatFractional(
		1000000000,
		defaultNumberBase,
		defaultFormatFractionalSize,
		defaultFractionalSeparator,
	)
	require.NoError(t, err)
	require.Equal(t, "", formated)

	formated, err = formatFractional(
		1000000001,
		defaultNumberBase,
		defaultFormatFractionalSize,
		defaultFractionalSeparator,
	)
	require.NoError(t, err)
	require.Equal(t, ".000000001", formated)

	formated, err = formatFractional(
		1000000010,
		defaultNumberBase,
		defaultFormatFractionalSize,
		defaultFractionalSeparator,
	)
	require.NoError(t, err)
	require.Equal(t, ".00000001", formated)

	formated, err = formatFractional(
		1000000100,
		defaultNumberBase,
		defaultFormatFractionalSize,
		defaultFractionalSeparator,
	)
	require.NoError(t, err)
	require.Equal(t, ".0000001", formated)

	formated, err = formatFractional(
		1000001000,
		defaultNumberBase,
		defaultFormatFractionalSize,
		defaultFractionalSeparator,
	)
	require.NoError(t, err)
	require.Equal(t, ".000001", formated)

	formated, err = formatFractional(
		1000010000,
		defaultNumberBase,
		defaultFormatFractionalSize,
		defaultFractionalSeparator,
	)
	require.NoError(t, err)
	require.Equal(t, ".00001", formated)

	formated, err = formatFractional(
		1000100000,
		defaultNumberBase,
		defaultFormatFractionalSize,
		defaultFractionalSeparator,
	)
	require.NoError(t, err)
	require.Equal(t, ".0001", formated)

	formated, err = formatFractional(
		1001000000,
		defaultNumberBase,
		defaultFormatFractionalSize,
		defaultFractionalSeparator,
	)
	require.NoError(t, err)
	require.Equal(t, ".001", formated)

	formated, err = formatFractional(
		1010000000,
		defaultNumberBase,
		defaultFormatFractionalSize,
		defaultFractionalSeparator,
	)
	require.NoError(t, err)
	require.Equal(t, ".01", formated)

	formated, err = formatFractional(
		1100000000,
		defaultNumberBase,
		defaultFormatFractionalSize,
		defaultFractionalSeparator,
	)
	require.NoError(t, err)
	require.Equal(t, ".1", formated)

	formated, err = formatFractional(
		2000000000,
		defaultNumberBase,
		defaultFormatFractionalSize,
		defaultFractionalSeparator,
	)
	require.NoError(t, err)
	require.Equal(t, "", formated)

	formated, err = formatFractional(
		2987654321,
		defaultNumberBase,
		defaultFormatFractionalSize,
		defaultFractionalSeparator,
	)
	require.NoError(t, err)
	require.Equal(t, ".987654321", formated)

	formated, err = formatFractional(
		2123456789,
		defaultNumberBase,
		defaultFormatFractionalSize,
		defaultFractionalSeparator,
	)
	require.NoError(t, err)
	require.Equal(t, ".123456789", formated)

	formated, err = formatFractional(
		2123406789,
		defaultNumberBase,
		defaultFormatFractionalSize,
		defaultFractionalSeparator,
	)
	require.NoError(t, err)
	require.Equal(t, ".123406789", formated)

	formated, err = formatFractional(
		2987604321,
		defaultNumberBase,
		defaultFormatFractionalSize,
		defaultFractionalSeparator,
	)
	require.NoError(t, err)
	require.Equal(t, ".987604321", formated)

	formated, err = formatFractional(
		2987604320,
		defaultNumberBase,
		defaultFormatFractionalSize,
		defaultFractionalSeparator,
	)
	require.NoError(t, err)
	require.Equal(t, ".98760432", formated)
}

func TestFormatFractionalRequireError(t *testing.T) {
	formated, err := formatFractional(
		2987604320,
		0,
		defaultFormatFractionalSize,
		defaultFractionalSeparator,
	)
	require.Error(t, err)
	require.Equal(t, "", formated)

	formated, err = formatFractional(
		2987604320,
		9223372036854775808,
		0,
		defaultFractionalSeparator,
	)
	require.Error(t, err)
	require.Equal(t, "", formated)

	formated, err = formatFractional(
		2987604320,
		defaultNumberBase,
		19,
		defaultFractionalSeparator,
	)
	require.Error(t, err)
	require.Equal(t, "", formated)

	formated, err = formatFractional(
		2987604320,
		defaultNumberBase,
		20,
		defaultFractionalSeparator,
	)
	require.Error(t, err)
	require.Equal(t, "", formated)

	formated, err = formatFractional(
		2987604320,
		defaultNumberBase,
		21,
		defaultFractionalSeparator,
	)
	require.Error(t, err)
	require.Equal(t, "", formated)
}
