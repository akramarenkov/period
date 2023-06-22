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

	sum, overflow = safeSum[int8](0, 126)
	require.Equal(t, int8(126), sum)
	require.Equal(t, false, overflow)

	sum, overflow = safeSum[int8](126, 1)
	require.Equal(t, int8(127), sum)
	require.Equal(t, false, overflow)

	sum, overflow = safeSum[int8](1, 126)
	require.Equal(t, int8(127), sum)
	require.Equal(t, false, overflow)

	sum, overflow = safeSum[int8](127, 1)
	require.Equal(t, int8(0), sum)
	require.Equal(t, true, overflow)

	sum, overflow = safeSum[int8](1, 127)
	require.Equal(t, int8(0), sum)
	require.Equal(t, true, overflow)

	sum, overflow = safeSum[int8](127, 2)
	require.Equal(t, int8(0), sum)
	require.Equal(t, true, overflow)

	sum, overflow = safeSum[int8](2, 127)
	require.Equal(t, int8(0), sum)
	require.Equal(t, true, overflow)

	sum, overflow = safeSum[int8](127, 3)
	require.Equal(t, int8(0), sum)
	require.Equal(t, true, overflow)

	sum, overflow = safeSum[int8](3, 127)
	require.Equal(t, int8(0), sum)
	require.Equal(t, true, overflow)

	sum, overflow = safeSum[int8](127, 125)
	require.Equal(t, int8(0), sum)
	require.Equal(t, true, overflow)

	sum, overflow = safeSum[int8](125, 127)
	require.Equal(t, int8(0), sum)
	require.Equal(t, true, overflow)

	sum, overflow = safeSum[int8](127, 126)
	require.Equal(t, int8(0), sum)
	require.Equal(t, true, overflow)

	sum, overflow = safeSum[int8](126, 127)
	require.Equal(t, int8(0), sum)
	require.Equal(t, true, overflow)

	sum, overflow = safeSum[int8](127, 127)
	require.Equal(t, int8(0), sum)
	require.Equal(t, true, overflow)

	sum, overflow = safeSum[int8](-127, 0)
	require.Equal(t, int8(-127), sum)
	require.Equal(t, false, overflow)

	sum, overflow = safeSum[int8](0, -127)
	require.Equal(t, int8(-127), sum)
	require.Equal(t, false, overflow)

	sum, overflow = safeSum[int8](-127, -1)
	require.Equal(t, int8(-128), sum)
	require.Equal(t, false, overflow)

	sum, overflow = safeSum[int8](-1, -127)
	require.Equal(t, int8(-128), sum)
	require.Equal(t, false, overflow)

	sum, overflow = safeSum[int8](-128, -1)
	require.Equal(t, int8(0), sum)
	require.Equal(t, true, overflow)

	sum, overflow = safeSum[int8](-1, -128)
	require.Equal(t, int8(0), sum)
	require.Equal(t, true, overflow)

	sum, overflow = safeSum[int8](-128, -2)
	require.Equal(t, int8(0), sum)
	require.Equal(t, true, overflow)

	sum, overflow = safeSum[int8](-2, -128)
	require.Equal(t, int8(0), sum)
	require.Equal(t, true, overflow)

	sum, overflow = safeSum[int8](-128, -3)
	require.Equal(t, int8(0), sum)
	require.Equal(t, true, overflow)

	sum, overflow = safeSum[int8](-3, -128)
	require.Equal(t, int8(0), sum)
	require.Equal(t, true, overflow)

	sum, overflow = safeSum[int8](-128, -126)
	require.Equal(t, int8(0), sum)
	require.Equal(t, true, overflow)

	sum, overflow = safeSum[int8](-126, -128)
	require.Equal(t, int8(0), sum)
	require.Equal(t, true, overflow)

	sum, overflow = safeSum[int8](-128, -127)
	require.Equal(t, int8(0), sum)
	require.Equal(t, true, overflow)

	sum, overflow = safeSum[int8](-127, -128)
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

	sum, overflow = safeSum[uint8](0, 254)
	require.Equal(t, uint8(254), sum)
	require.Equal(t, false, overflow)

	sum, overflow = safeSum[uint8](254, 1)
	require.Equal(t, uint8(255), sum)
	require.Equal(t, false, overflow)

	sum, overflow = safeSum[uint8](1, 254)
	require.Equal(t, uint8(255), sum)
	require.Equal(t, false, overflow)

	sum, overflow = safeSum[uint8](255, 1)
	require.Equal(t, uint8(0), sum)
	require.Equal(t, true, overflow)

	sum, overflow = safeSum[uint8](1, 255)
	require.Equal(t, uint8(0), sum)
	require.Equal(t, true, overflow)

	sum, overflow = safeSum[uint8](255, 2)
	require.Equal(t, uint8(0), sum)
	require.Equal(t, true, overflow)

	sum, overflow = safeSum[uint8](2, 255)
	require.Equal(t, uint8(0), sum)
	require.Equal(t, true, overflow)

	sum, overflow = safeSum[uint8](255, 3)
	require.Equal(t, uint8(0), sum)
	require.Equal(t, true, overflow)

	sum, overflow = safeSum[uint8](3, 255)
	require.Equal(t, uint8(0), sum)
	require.Equal(t, true, overflow)

	sum, overflow = safeSum[uint8](255, 253)
	require.Equal(t, uint8(0), sum)
	require.Equal(t, true, overflow)

	sum, overflow = safeSum[uint8](253, 255)
	require.Equal(t, uint8(0), sum)
	require.Equal(t, true, overflow)

	sum, overflow = safeSum[uint8](255, 254)
	require.Equal(t, uint8(0), sum)
	require.Equal(t, true, overflow)

	sum, overflow = safeSum[uint8](254, 255)
	require.Equal(t, uint8(0), sum)
	require.Equal(t, true, overflow)

	sum, overflow = safeSum[uint8](255, 255)
	require.Equal(t, uint8(0), sum)
	require.Equal(t, true, overflow)

	sum, overflow = safeSum[uint8](128, 128)
	require.Equal(t, uint8(0), sum)
	require.Equal(t, true, overflow)

	sum, overflow = safeSum[uint8](127, 129)
	require.Equal(t, uint8(0), sum)
	require.Equal(t, true, overflow)

	sum, overflow = safeSum[uint8](129, 127)
	require.Equal(t, uint8(0), sum)
	require.Equal(t, true, overflow)
}

func TestSafeSlowMultiplicationInt8(t *testing.T) {
	product, overflow := safeSlowMultiplication[int8](0, 0)
	require.Equal(t, int8(0), product)
	require.Equal(t, false, overflow)

	product, overflow = safeSlowMultiplication[int8](2, 0)
	require.Equal(t, int8(0), product)
	require.Equal(t, false, overflow)

	product, overflow = safeSlowMultiplication[int8](0, 2)
	require.Equal(t, int8(0), product)
	require.Equal(t, false, overflow)

	product, overflow = safeSlowMultiplication[int8](-2, 0)
	require.Equal(t, int8(0), product)
	require.Equal(t, false, overflow)

	product, overflow = safeSlowMultiplication[int8](0, -2)
	require.Equal(t, int8(0), product)
	require.Equal(t, false, overflow)

	product, overflow = safeSlowMultiplication[int8](3, 0)
	require.Equal(t, int8(0), product)
	require.Equal(t, false, overflow)

	product, overflow = safeSlowMultiplication[int8](0, 3)
	require.Equal(t, int8(0), product)
	require.Equal(t, false, overflow)

	product, overflow = safeSlowMultiplication[int8](-3, 0)
	require.Equal(t, int8(0), product)
	require.Equal(t, false, overflow)

	product, overflow = safeSlowMultiplication[int8](0, -3)
	require.Equal(t, int8(0), product)
	require.Equal(t, false, overflow)

	product, overflow = safeSlowMultiplication[int8](2, -3)
	require.Equal(t, int8(-6), product)
	require.Equal(t, false, overflow)

	product, overflow = safeSlowMultiplication[int8](-3, 2)
	require.Equal(t, int8(-6), product)
	require.Equal(t, false, overflow)

	product, overflow = safeSlowMultiplication[int8](-3, -2)
	require.Equal(t, int8(6), product)
	require.Equal(t, false, overflow)

	product, overflow = safeSlowMultiplication[int8](127, 1)
	require.Equal(t, int8(127), product)
	require.Equal(t, false, overflow)

	product, overflow = safeSlowMultiplication[int8](-128, 1)
	require.Equal(t, int8(-128), product)
	require.Equal(t, false, overflow)

	product, overflow = safeSlowMultiplication[int8](127, 2)
	require.Equal(t, int8(0), product)
	require.Equal(t, true, overflow)

	product, overflow = safeSlowMultiplication[int8](-128, 2)
	require.Equal(t, int8(0), product)
	require.Equal(t, true, overflow)

	product, overflow = safeSlowMultiplication[int8](64, 3)
	require.Equal(t, int8(0), product)
	require.Equal(t, true, overflow)

	product, overflow = safeSlowMultiplication[int8](-64, 3)
	require.Equal(t, int8(0), product)
	require.Equal(t, true, overflow)

	product, overflow = safeSlowMultiplication[int8](-128, -128)
	require.Equal(t, int8(0), product)
	require.Equal(t, true, overflow)
}

func TestSafeSlowMultiplicationUint8(t *testing.T) {
	product, overflow := safeSlowMultiplication[uint8](0, 0)
	require.Equal(t, uint8(0), product)
	require.Equal(t, false, overflow)

	product, overflow = safeSlowMultiplication[uint8](2, 0)
	require.Equal(t, uint8(0), product)
	require.Equal(t, false, overflow)

	product, overflow = safeSlowMultiplication[uint8](0, 2)
	require.Equal(t, uint8(0), product)
	require.Equal(t, false, overflow)

	product, overflow = safeSlowMultiplication[uint8](3, 0)
	require.Equal(t, uint8(0), product)
	require.Equal(t, false, overflow)

	product, overflow = safeSlowMultiplication[uint8](0, 3)
	require.Equal(t, uint8(0), product)
	require.Equal(t, false, overflow)

	product, overflow = safeSlowMultiplication[uint8](2, 3)
	require.Equal(t, uint8(6), product)
	require.Equal(t, false, overflow)

	product, overflow = safeSlowMultiplication[uint8](3, 2)
	require.Equal(t, uint8(6), product)
	require.Equal(t, false, overflow)

	product, overflow = safeSlowMultiplication[uint8](127, 2)
	require.Equal(t, uint8(254), product)
	require.Equal(t, false, overflow)

	product, overflow = safeSlowMultiplication[uint8](255, 1)
	require.Equal(t, uint8(255), product)
	require.Equal(t, false, overflow)

	product, overflow = safeSlowMultiplication[uint8](128, 2)
	require.Equal(t, uint8(0), product)
	require.Equal(t, true, overflow)

	product, overflow = safeSlowMultiplication[uint8](127, 3)
	require.Equal(t, uint8(0), product)
	require.Equal(t, true, overflow)

	product, overflow = safeSlowMultiplication[uint8](255, 2)
	require.Equal(t, uint8(0), product)
	require.Equal(t, true, overflow)

	product, overflow = safeSlowMultiplication[uint8](255, 255)
	require.Equal(t, uint8(0), product)
	require.Equal(t, true, overflow)
}

func TestSafeMultiplicationUint8(t *testing.T) {
	variants := 0

	for first := 0; first <= 255; first++ {
		for second := 0; second <= 255; second++ {
			first8 := uint8(first)
			second8 := uint8(second)

			product, overflow := safeMultiplication[uint8](first8, second8)
			referenceProduct, referenceOverflow := safeSlowMultiplication[uint8](first8, second8)

			require.Equal(t, referenceProduct, product)
			require.Equal(t, referenceOverflow, overflow)

			variants++
		}
	}

	require.Equal(t, 65536, variants)
}

func TestSafeMultiplicationInt8(t *testing.T) {
	variants := 0

	for first := -128; first <= 127; first++ {
		for second := -128; second <= 127; second++ {
			first8 := int8(first)
			second8 := int8(second)

			product, overflow := safeMultiplication[int8](first8, second8)
			referenceProduct, referenceOverflow := safeSlowMultiplication[int8](first8, second8)

			require.Equal(t, referenceProduct, product)
			require.Equal(t, referenceOverflow, overflow, "%v * %v = %v (%v)", first, second, product, referenceProduct)

			variants++
		}
	}

	require.Equal(t, 65536, variants)
}
