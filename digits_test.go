package period

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSymbolToDigit(t *testing.T) {
	symbol, err := symbolToDigit('0')
	require.NoError(t, err)
	require.Equal(t, uint(0), symbol)

	symbol, err = symbolToDigit('1')
	require.NoError(t, err)
	require.Equal(t, uint(1), symbol)

	symbol, err = symbolToDigit('2')
	require.NoError(t, err)
	require.Equal(t, uint(2), symbol)

	symbol, err = symbolToDigit('3')
	require.NoError(t, err)
	require.Equal(t, uint(3), symbol)

	symbol, err = symbolToDigit('4')
	require.NoError(t, err)
	require.Equal(t, uint(4), symbol)

	symbol, err = symbolToDigit('5')
	require.NoError(t, err)
	require.Equal(t, uint(5), symbol)

	symbol, err = symbolToDigit('6')
	require.NoError(t, err)
	require.Equal(t, uint(6), symbol)

	symbol, err = symbolToDigit('7')
	require.NoError(t, err)
	require.Equal(t, uint(7), symbol)

	symbol, err = symbolToDigit('8')
	require.NoError(t, err)
	require.Equal(t, uint(8), symbol)

	symbol, err = symbolToDigit('9')
	require.NoError(t, err)
	require.Equal(t, uint(9), symbol)

	symbol, err = symbolToDigit(rune(0))
	require.Error(t, err)
	require.Equal(t, uint(0), symbol)
}

func TestDigitToSymbol(t *testing.T) {
	symbol, err := digitToSymbol(0)
	require.NoError(t, err)
	require.Equal(t, '0', symbol)

	symbol, err = digitToSymbol(1)
	require.NoError(t, err)
	require.Equal(t, '1', symbol)

	symbol, err = digitToSymbol(2)
	require.NoError(t, err)
	require.Equal(t, '2', symbol)

	symbol, err = digitToSymbol(3)
	require.NoError(t, err)
	require.Equal(t, '3', symbol)

	symbol, err = digitToSymbol(4)
	require.NoError(t, err)
	require.Equal(t, '4', symbol)

	symbol, err = digitToSymbol(5)
	require.NoError(t, err)
	require.Equal(t, '5', symbol)

	symbol, err = digitToSymbol(6)
	require.NoError(t, err)
	require.Equal(t, '6', symbol)

	symbol, err = digitToSymbol(7)
	require.NoError(t, err)
	require.Equal(t, '7', symbol)

	symbol, err = digitToSymbol(8)
	require.NoError(t, err)
	require.Equal(t, '8', symbol)

	symbol, err = digitToSymbol(9)
	require.NoError(t, err)
	require.Equal(t, '9', symbol)

	symbol, err = digitToSymbol(10)
	require.Error(t, err)
	require.Equal(t, rune(0), symbol)
}
