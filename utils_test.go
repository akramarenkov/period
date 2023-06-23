package period

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSafeSumInt8(t *testing.T) {
	sum, overflow := safeSumInt[int8](0, 0)
	require.Equal(t, int8(0), sum)
	require.Equal(t, false, overflow)

	sum, overflow = safeSumInt[int8](126, 0)
	require.Equal(t, int8(126), sum)
	require.Equal(t, false, overflow)

	sum, overflow = safeSumInt[int8](0, 126)
	require.Equal(t, int8(126), sum)
	require.Equal(t, false, overflow)

	sum, overflow = safeSumInt[int8](126, 1)
	require.Equal(t, int8(127), sum)
	require.Equal(t, false, overflow)

	sum, overflow = safeSumInt[int8](1, 126)
	require.Equal(t, int8(127), sum)
	require.Equal(t, false, overflow)

	sum, overflow = safeSumInt[int8](-127, 0)
	require.Equal(t, int8(-127), sum)
	require.Equal(t, false, overflow)

	sum, overflow = safeSumInt[int8](0, -127)
	require.Equal(t, int8(-127), sum)
	require.Equal(t, false, overflow)

	sum, overflow = safeSumInt[int8](-127, -1)
	require.Equal(t, int8(-128), sum)
	require.Equal(t, false, overflow)

	sum, overflow = safeSumInt[int8](-1, -127)
	require.Equal(t, int8(-128), sum)
	require.Equal(t, false, overflow)
}

func TestSafeSumOverflowInt8(t *testing.T) {
	sum, overflow := safeSumInt[int8](127, 1)
	require.Equal(t, int8(0), sum)
	require.Equal(t, true, overflow)

	sum, overflow = safeSumInt[int8](1, 127)
	require.Equal(t, int8(0), sum)
	require.Equal(t, true, overflow)

	sum, overflow = safeSumInt[int8](127, 2)
	require.Equal(t, int8(0), sum)
	require.Equal(t, true, overflow)

	sum, overflow = safeSumInt[int8](2, 127)
	require.Equal(t, int8(0), sum)
	require.Equal(t, true, overflow)

	sum, overflow = safeSumInt[int8](127, 3)
	require.Equal(t, int8(0), sum)
	require.Equal(t, true, overflow)

	sum, overflow = safeSumInt[int8](3, 127)
	require.Equal(t, int8(0), sum)
	require.Equal(t, true, overflow)

	sum, overflow = safeSumInt[int8](127, 125)
	require.Equal(t, int8(0), sum)
	require.Equal(t, true, overflow)

	sum, overflow = safeSumInt[int8](125, 127)
	require.Equal(t, int8(0), sum)
	require.Equal(t, true, overflow)

	sum, overflow = safeSumInt[int8](127, 126)
	require.Equal(t, int8(0), sum)
	require.Equal(t, true, overflow)

	sum, overflow = safeSumInt[int8](126, 127)
	require.Equal(t, int8(0), sum)
	require.Equal(t, true, overflow)

	sum, overflow = safeSumInt[int8](127, 127)
	require.Equal(t, int8(0), sum)
	require.Equal(t, true, overflow)

	sum, overflow = safeSumInt[int8](-128, -1)
	require.Equal(t, int8(0), sum)
	require.Equal(t, true, overflow)

	sum, overflow = safeSumInt[int8](-1, -128)
	require.Equal(t, int8(0), sum)
	require.Equal(t, true, overflow)

	sum, overflow = safeSumInt[int8](-128, -2)
	require.Equal(t, int8(0), sum)
	require.Equal(t, true, overflow)

	sum, overflow = safeSumInt[int8](-2, -128)
	require.Equal(t, int8(0), sum)
	require.Equal(t, true, overflow)

	sum, overflow = safeSumInt[int8](-128, -3)
	require.Equal(t, int8(0), sum)
	require.Equal(t, true, overflow)

	sum, overflow = safeSumInt[int8](-3, -128)
	require.Equal(t, int8(0), sum)
	require.Equal(t, true, overflow)

	sum, overflow = safeSumInt[int8](-128, -126)
	require.Equal(t, int8(0), sum)
	require.Equal(t, true, overflow)

	sum, overflow = safeSumInt[int8](-126, -128)
	require.Equal(t, int8(0), sum)
	require.Equal(t, true, overflow)

	sum, overflow = safeSumInt[int8](-128, -127)
	require.Equal(t, int8(0), sum)
	require.Equal(t, true, overflow)

	sum, overflow = safeSumInt[int8](-127, -128)
	require.Equal(t, int8(0), sum)
	require.Equal(t, true, overflow)

	sum, overflow = safeSumInt[int8](-128, -128)
	require.Equal(t, int8(0), sum)
	require.Equal(t, true, overflow)
}

func TestSafeSumUint8(t *testing.T) {
	sum, overflow := safeSumInt[uint8](0, 0)
	require.Equal(t, uint8(0), sum)
	require.Equal(t, false, overflow)

	sum, overflow = safeSumInt[uint8](254, 0)
	require.Equal(t, uint8(254), sum)
	require.Equal(t, false, overflow)

	sum, overflow = safeSumInt[uint8](0, 254)
	require.Equal(t, uint8(254), sum)
	require.Equal(t, false, overflow)

	sum, overflow = safeSumInt[uint8](254, 1)
	require.Equal(t, uint8(255), sum)
	require.Equal(t, false, overflow)

	sum, overflow = safeSumInt[uint8](1, 254)
	require.Equal(t, uint8(255), sum)
	require.Equal(t, false, overflow)
}

func TestSafeSumOverflowUint8(t *testing.T) {
	sum, overflow := safeSumInt[uint8](255, 1)
	require.Equal(t, uint8(0), sum)
	require.Equal(t, true, overflow)

	sum, overflow = safeSumInt[uint8](1, 255)
	require.Equal(t, uint8(0), sum)
	require.Equal(t, true, overflow)

	sum, overflow = safeSumInt[uint8](255, 2)
	require.Equal(t, uint8(0), sum)
	require.Equal(t, true, overflow)

	sum, overflow = safeSumInt[uint8](2, 255)
	require.Equal(t, uint8(0), sum)
	require.Equal(t, true, overflow)

	sum, overflow = safeSumInt[uint8](255, 3)
	require.Equal(t, uint8(0), sum)
	require.Equal(t, true, overflow)

	sum, overflow = safeSumInt[uint8](3, 255)
	require.Equal(t, uint8(0), sum)
	require.Equal(t, true, overflow)

	sum, overflow = safeSumInt[uint8](255, 253)
	require.Equal(t, uint8(0), sum)
	require.Equal(t, true, overflow)

	sum, overflow = safeSumInt[uint8](253, 255)
	require.Equal(t, uint8(0), sum)
	require.Equal(t, true, overflow)

	sum, overflow = safeSumInt[uint8](255, 254)
	require.Equal(t, uint8(0), sum)
	require.Equal(t, true, overflow)

	sum, overflow = safeSumInt[uint8](254, 255)
	require.Equal(t, uint8(0), sum)
	require.Equal(t, true, overflow)

	sum, overflow = safeSumInt[uint8](255, 255)
	require.Equal(t, uint8(0), sum)
	require.Equal(t, true, overflow)

	sum, overflow = safeSumInt[uint8](128, 128)
	require.Equal(t, uint8(0), sum)
	require.Equal(t, true, overflow)

	sum, overflow = safeSumInt[uint8](127, 129)
	require.Equal(t, uint8(0), sum)
	require.Equal(t, true, overflow)

	sum, overflow = safeSumInt[uint8](129, 127)
	require.Equal(t, uint8(0), sum)
	require.Equal(t, true, overflow)
}

func TestIsMaxNegative(t *testing.T) {
	require.Equal(t, true, isMaxNegative[int8](-128))
	require.Equal(t, false, isMaxNegative[int8](-127))
	require.Equal(t, false, isMaxNegative[int8](-126))
	require.Equal(t, false, isMaxNegative[int8](-125))
	require.Equal(t, false, isMaxNegative[int8](-3))
	require.Equal(t, false, isMaxNegative[int8](-2))
	require.Equal(t, false, isMaxNegative[int8](-1))
	require.Equal(t, false, isMaxNegative[int8](0))
	require.Equal(t, false, isMaxNegative[int8](1))
	require.Equal(t, false, isMaxNegative[int8](2))
	require.Equal(t, false, isMaxNegative[int8](3))
	require.Equal(t, false, isMaxNegative[int8](125))
	require.Equal(t, false, isMaxNegative[int8](126))
	require.Equal(t, false, isMaxNegative[int8](127))

	require.Equal(t, false, isMaxNegative[uint8](0))
	require.Equal(t, false, isMaxNegative[uint8](1))
	require.Equal(t, false, isMaxNegative[uint8](255))
}

func TestSafeProductInt8(t *testing.T) {
	product, overflow := safeProductInt[int8](0, 0)
	require.Equal(t, int8(0), product)
	require.Equal(t, false, overflow)

	product, overflow = safeProductInt[int8](2, 0)
	require.Equal(t, int8(0), product)
	require.Equal(t, false, overflow)

	product, overflow = safeProductInt[int8](0, 2)
	require.Equal(t, int8(0), product)
	require.Equal(t, false, overflow)

	product, overflow = safeProductInt[int8](3, 0)
	require.Equal(t, int8(0), product)
	require.Equal(t, false, overflow)

	product, overflow = safeProductInt[int8](0, 3)
	require.Equal(t, int8(0), product)
	require.Equal(t, false, overflow)

	product, overflow = safeProductInt[int8](-2, 0)
	require.Equal(t, int8(0), product)
	require.Equal(t, false, overflow)

	product, overflow = safeProductInt[int8](0, -2)
	require.Equal(t, int8(0), product)
	require.Equal(t, false, overflow)

	product, overflow = safeProductInt[int8](-3, 0)
	require.Equal(t, int8(0), product)
	require.Equal(t, false, overflow)

	product, overflow = safeProductInt[int8](0, -3)
	require.Equal(t, int8(0), product)
	require.Equal(t, false, overflow)

	product, overflow = safeProductInt[int8](2, 3)
	require.Equal(t, int8(6), product)
	require.Equal(t, false, overflow)

	product, overflow = safeProductInt[int8](3, 2)
	require.Equal(t, int8(6), product)
	require.Equal(t, false, overflow)

	product, overflow = safeProductInt[int8](-2, 3)
	require.Equal(t, int8(-6), product)
	require.Equal(t, false, overflow)

	product, overflow = safeProductInt[int8](3, -2)
	require.Equal(t, int8(-6), product)
	require.Equal(t, false, overflow)

	product, overflow = safeProductInt[int8](2, -3)
	require.Equal(t, int8(-6), product)
	require.Equal(t, false, overflow)

	product, overflow = safeProductInt[int8](-3, 2)
	require.Equal(t, int8(-6), product)
	require.Equal(t, false, overflow)

	product, overflow = safeProductInt[int8](-2, -3)
	require.Equal(t, int8(6), product)
	require.Equal(t, false, overflow)

	product, overflow = safeProductInt[int8](-3, -2)
	require.Equal(t, int8(6), product)
	require.Equal(t, false, overflow)

	product, overflow = safeProductInt[int8](1, -127)
	require.Equal(t, int8(-127), product)
	require.Equal(t, false, overflow)

	product, overflow = safeProductInt[int8](-127, 1)
	require.Equal(t, int8(-127), product)
	require.Equal(t, false, overflow)

	product, overflow = safeProductInt[int8](127, -1)
	require.Equal(t, int8(-127), product)
	require.Equal(t, false, overflow)

	product, overflow = safeProductInt[int8](-1, 127)
	require.Equal(t, int8(-127), product)
	require.Equal(t, false, overflow)

	product, overflow = safeProductInt[int8](-127, -1)
	require.Equal(t, int8(127), product)
	require.Equal(t, false, overflow)

	product, overflow = safeProductInt[int8](-1, -127)
	require.Equal(t, int8(127), product)
	require.Equal(t, false, overflow)

	product, overflow = safeProductInt[int8](-128, 1)
	require.Equal(t, int8(-128), product)
	require.Equal(t, false, overflow)

	product, overflow = safeProductInt[int8](1, -128)
	require.Equal(t, int8(-128), product)
	require.Equal(t, false, overflow)
}

func TestSafeProductOverflowInt8(t *testing.T) {
	product, overflow := safeProductInt[int8](127, 2)
	require.Equal(t, int8(0), product)
	require.Equal(t, true, overflow)

	product, overflow = safeProductInt[int8](2, 127)
	require.Equal(t, int8(0), product)
	require.Equal(t, true, overflow)

	product, overflow = safeProductInt[int8](64, 2)
	require.Equal(t, int8(0), product)
	require.Equal(t, true, overflow)

	product, overflow = safeProductInt[int8](2, 64)
	require.Equal(t, int8(0), product)
	require.Equal(t, true, overflow)

	product, overflow = safeProductInt[int8](127, 127)
	require.Equal(t, int8(0), product)
	require.Equal(t, true, overflow)

	product, overflow = safeProductInt[int8](-127, 2)
	require.Equal(t, int8(0), product)
	require.Equal(t, true, overflow)

	product, overflow = safeProductInt[int8](2, -127)
	require.Equal(t, int8(0), product)
	require.Equal(t, true, overflow)

	product, overflow = safeProductInt[int8](-2, 127)
	require.Equal(t, int8(0), product)
	require.Equal(t, true, overflow)

	product, overflow = safeProductInt[int8](127, -2)
	require.Equal(t, int8(0), product)
	require.Equal(t, true, overflow)

	product, overflow = safeProductInt[int8](-127, -2)
	require.Equal(t, int8(0), product)
	require.Equal(t, true, overflow)

	product, overflow = safeProductInt[int8](-2, -127)
	require.Equal(t, int8(0), product)
	require.Equal(t, true, overflow)

	product, overflow = safeProductInt[int8](127, -127)
	require.Equal(t, int8(0), product)
	require.Equal(t, true, overflow)

	product, overflow = safeProductInt[int8](-127, 127)
	require.Equal(t, int8(0), product)
	require.Equal(t, true, overflow)

	product, overflow = safeProductInt[int8](-127, -127)
	require.Equal(t, int8(0), product)
	require.Equal(t, true, overflow)

	product, overflow = safeProductInt[int8](-128, -1)
	require.Equal(t, int8(0), product)
	require.Equal(t, true, overflow)

	product, overflow = safeProductInt[int8](-1, -128)
	require.Equal(t, int8(0), product)
	require.Equal(t, true, overflow)

	product, overflow = safeProductInt[int8](-128, -128)
	require.Equal(t, int8(0), product)
	require.Equal(t, true, overflow)
}

func TestSafeProductUint8(t *testing.T) {
	product, overflow := safeProductInt[uint8](0, 0)
	require.Equal(t, uint8(0), product)
	require.Equal(t, false, overflow)

	product, overflow = safeProductInt[uint8](2, 0)
	require.Equal(t, uint8(0), product)
	require.Equal(t, false, overflow)

	product, overflow = safeProductInt[uint8](0, 2)
	require.Equal(t, uint8(0), product)
	require.Equal(t, false, overflow)

	product, overflow = safeProductInt[uint8](3, 0)
	require.Equal(t, uint8(0), product)
	require.Equal(t, false, overflow)

	product, overflow = safeProductInt[uint8](0, 3)
	require.Equal(t, uint8(0), product)
	require.Equal(t, false, overflow)

	product, overflow = safeProductInt[uint8](3, 2)
	require.Equal(t, uint8(6), product)
	require.Equal(t, false, overflow)

	product, overflow = safeProductInt[uint8](2, 3)
	require.Equal(t, uint8(6), product)
	require.Equal(t, false, overflow)

	product, overflow = safeProductInt[uint8](85, 3)
	require.Equal(t, uint8(255), product)
	require.Equal(t, false, overflow)

	product, overflow = safeProductInt[uint8](3, 85)
	require.Equal(t, uint8(255), product)
	require.Equal(t, false, overflow)

	product, overflow = safeProductInt[uint8](255, 1)
	require.Equal(t, uint8(255), product)
	require.Equal(t, false, overflow)

	product, overflow = safeProductInt[uint8](1, 255)
	require.Equal(t, uint8(255), product)
	require.Equal(t, false, overflow)
}

func TestSafeProductOverflowUint8(t *testing.T) {
	product, overflow := safeProductInt[uint8](255, 2)
	require.Equal(t, uint8(0), product)
	require.Equal(t, true, overflow)

	product, overflow = safeProductInt[uint8](2, 255)
	require.Equal(t, uint8(0), product)
	require.Equal(t, true, overflow)

	product, overflow = safeProductInt[uint8](128, 2)
	require.Equal(t, uint8(0), product)
	require.Equal(t, true, overflow)

	product, overflow = safeProductInt[uint8](2, 128)
	require.Equal(t, uint8(0), product)
	require.Equal(t, true, overflow)

	product, overflow = safeProductInt[uint8](86, 3)
	require.Equal(t, uint8(0), product)
	require.Equal(t, true, overflow)

	product, overflow = safeProductInt[uint8](3, 86)
	require.Equal(t, uint8(0), product)
	require.Equal(t, true, overflow)

	product, overflow = safeProductInt[uint8](64, 4)
	require.Equal(t, uint8(0), product)
	require.Equal(t, true, overflow)

	product, overflow = safeProductInt[uint8](4, 64)
	require.Equal(t, uint8(0), product)
	require.Equal(t, true, overflow)

	product, overflow = safeProductInt[uint8](255, 254)
	require.Equal(t, uint8(0), product)
	require.Equal(t, true, overflow)

	product, overflow = safeProductInt[uint8](254, 255)
	require.Equal(t, uint8(0), product)
	require.Equal(t, true, overflow)

	product, overflow = safeProductInt[uint8](255, 255)
	require.Equal(t, uint8(0), product)
	require.Equal(t, true, overflow)
}
