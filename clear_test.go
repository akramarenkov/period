package period

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestClearInteger(t *testing.T) {
	require.Equal(t, "", string(clearInteger([]rune(""))))
	require.Equal(t, "0", string(clearInteger([]rune("0"))))
	require.Equal(t, "0", string(clearInteger([]rune("00"))))
	require.Equal(t, "1", string(clearInteger([]rune("1"))))
	require.Equal(t, "12", string(clearInteger([]rune("12"))))
	require.Equal(t, "12", string(clearInteger([]rune("012"))))
	require.Equal(t, "12", string(clearInteger([]rune("0012"))))
	require.Equal(t, "10012", string(clearInteger([]rune("10012"))))
	require.Equal(t, "10012", string(clearInteger([]rune("010012"))))
	require.Equal(t, "100120", string(clearInteger([]rune("0100120"))))
	require.Equal(t, "100120", string(clearInteger([]rune("00100120"))))
	require.Equal(t, "1001200", string(clearInteger([]rune("001001200"))))
}

func TestClearFractional(t *testing.T) {
	trailed := clearFractional([]rune("0.001001001"), defaultFractionalSeparator)
	require.Equal(t, "0.001001001", string(trailed))

	trailed = clearFractional([]rune("0.001001010"), defaultFractionalSeparator)
	require.Equal(t, "0.00100101", string(trailed))

	trailed = clearFractional([]rune("0.001001100"), defaultFractionalSeparator)
	require.Equal(t, "0.0010011", string(trailed))

	trailed = clearFractional([]rune("0.001001000"), defaultFractionalSeparator)
	require.Equal(t, "0.001001", string(trailed))

	trailed = clearFractional([]rune("0.001010000"), defaultFractionalSeparator)
	require.Equal(t, "0.00101", string(trailed))

	trailed = clearFractional([]rune("0.001100000"), defaultFractionalSeparator)
	require.Equal(t, "0.0011", string(trailed))

	trailed = clearFractional([]rune("0.001000000"), defaultFractionalSeparator)
	require.Equal(t, "0.001", string(trailed))

	trailed = clearFractional([]rune("0.010000000"), defaultFractionalSeparator)
	require.Equal(t, "0.01", string(trailed))

	trailed = clearFractional([]rune("0.100000000"), defaultFractionalSeparator)
	require.Equal(t, "0.1", string(trailed))

	trailed = clearFractional([]rune("0.000000000"), defaultFractionalSeparator)
	require.Equal(t, "", string(trailed))

	trailed = clearFractional([]rune("1.000000000"), defaultFractionalSeparator)
	require.Equal(t, "1", string(trailed))

	trailed = clearFractional([]rune(""), defaultFractionalSeparator)
	require.Equal(t, "", string(trailed))
}
