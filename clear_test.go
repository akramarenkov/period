package period

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestClearInteger(t *testing.T) {
	require.Equal(t, "", clearInteger(""))
	require.Equal(t, "0", clearInteger("0"))
	require.Equal(t, "0", clearInteger("00"))
	require.Equal(t, "1", clearInteger("1"))
	require.Equal(t, "12", clearInteger("12"))
	require.Equal(t, "12", clearInteger("012"))
	require.Equal(t, "12", clearInteger("0012"))
	require.Equal(t, "10012", clearInteger("10012"))
	require.Equal(t, "10012", clearInteger("010012"))
	require.Equal(t, "100120", clearInteger("0100120"))
	require.Equal(t, "100120", clearInteger("00100120"))
	require.Equal(t, "1001200", clearInteger("001001200"))
}

func TestClearFractional(t *testing.T) {
	trailed := clearFractional("0.001001001", defaultFractionalSeparator)
	require.Equal(t, "0.001001001", trailed)

	trailed = clearFractional("0.001001010", defaultFractionalSeparator)
	require.Equal(t, "0.00100101", trailed)

	trailed = clearFractional("0.001001100", defaultFractionalSeparator)
	require.Equal(t, "0.0010011", trailed)

	trailed = clearFractional("0.001001000", defaultFractionalSeparator)
	require.Equal(t, "0.001001", trailed)

	trailed = clearFractional("0.001010000", defaultFractionalSeparator)
	require.Equal(t, "0.00101", trailed)

	trailed = clearFractional("0.001100000", defaultFractionalSeparator)
	require.Equal(t, "0.0011", trailed)

	trailed = clearFractional("0.001000000", defaultFractionalSeparator)
	require.Equal(t, "0.001", trailed)

	trailed = clearFractional("0.010000000", defaultFractionalSeparator)
	require.Equal(t, "0.01", trailed)

	trailed = clearFractional("0.100000000", defaultFractionalSeparator)
	require.Equal(t, "0.1", trailed)

	trailed = clearFractional("0.000000000", defaultFractionalSeparator)
	require.Equal(t, "", trailed)

	trailed = clearFractional("1.000000000", defaultFractionalSeparator)
	require.Equal(t, "1", trailed)

	trailed = clearFractional("", defaultFractionalSeparator)
	require.Equal(t, "", trailed)
}
