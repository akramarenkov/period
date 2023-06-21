package period

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSafeSumInt8(t *testing.T) {
	sum, overflow := safeSum[int8](0, 0)
	require.Equal(t, int8(0), sum)
	require.Equal(t, false, overflow)

	sum, overflow = safeSum[int8](126, 0)
	require.Equal(t, int8(126), sum)
	require.Equal(t, false, overflow)

	sum, overflow = safeSum[int8](126, 1)
	require.Equal(t, int8(127), sum)
	require.Equal(t, false, overflow)

	sum, overflow = safeSum[int8](127, 1)
	require.Equal(t, int8(0), sum)
	require.Equal(t, true, overflow)

	sum, overflow = safeSum[int8](127, 2)
	require.Equal(t, int8(0), sum)
	require.Equal(t, true, overflow)

	sum, overflow = safeSum[int8](127, 3)
	require.Equal(t, int8(0), sum)
	require.Equal(t, true, overflow)

	sum, overflow = safeSum[int8](127, 125)
	require.Equal(t, int8(0), sum)
	require.Equal(t, true, overflow)

	sum, overflow = safeSum[int8](127, 126)
	require.Equal(t, int8(0), sum)
	require.Equal(t, true, overflow)

	sum, overflow = safeSum[int8](127, 127)
	require.Equal(t, int8(0), sum)
	require.Equal(t, true, overflow)

	sum, overflow = safeSum[int8](-127, 0)
	require.Equal(t, int8(-127), sum)
	require.Equal(t, false, overflow)

	sum, overflow = safeSum[int8](-127, -1)
	require.Equal(t, int8(-128), sum)
	require.Equal(t, false, overflow)

	sum, overflow = safeSum[int8](-128, -1)
	require.Equal(t, int8(0), sum)
	require.Equal(t, true, overflow)

	sum, overflow = safeSum[int8](-128, -2)
	require.Equal(t, int8(0), sum)
	require.Equal(t, true, overflow)

	sum, overflow = safeSum[int8](-128, -3)
	require.Equal(t, int8(0), sum)
	require.Equal(t, true, overflow)

	sum, overflow = safeSum[int8](-128, -126)
	require.Equal(t, int8(0), sum)
	require.Equal(t, true, overflow)

	sum, overflow = safeSum[int8](-128, -127)
	require.Equal(t, int8(0), sum)
	require.Equal(t, true, overflow)

	sum, overflow = safeSum[int8](-128, -128)
	require.Equal(t, int8(0), sum)
	require.Equal(t, true, overflow)
}

func TestSafeSumUint8(t *testing.T) {
	sum, overflow := safeSum[uint8](0, 0)
	require.Equal(t, uint8(0), sum)
	require.Equal(t, false, overflow)

	sum, overflow = safeSum[uint8](254, 0)
	require.Equal(t, uint8(254), sum)
	require.Equal(t, false, overflow)

	sum, overflow = safeSum[uint8](254, 1)
	require.Equal(t, uint8(255), sum)
	require.Equal(t, false, overflow)

	sum, overflow = safeSum[uint8](255, 1)
	require.Equal(t, uint8(0), sum)
	require.Equal(t, true, overflow)

	sum, overflow = safeSum[uint8](255, 2)
	require.Equal(t, uint8(0), sum)
	require.Equal(t, true, overflow)

	sum, overflow = safeSum[uint8](255, 3)
	require.Equal(t, uint8(0), sum)
	require.Equal(t, true, overflow)

	sum, overflow = safeSum[uint8](255, 253)
	require.Equal(t, uint8(0), sum)
	require.Equal(t, true, overflow)

	sum, overflow = safeSum[uint8](255, 254)
	require.Equal(t, uint8(0), sum)
	require.Equal(t, true, overflow)

	sum, overflow = safeSum[uint8](255, 255)
	require.Equal(t, uint8(0), sum)
	require.Equal(t, true, overflow)
}
