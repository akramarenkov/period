package period

import (
	"errors"
	"unsafe"

	"github.com/akramarenkov/safe"
)

var (
	ErrNumberBaseIsZero = errors.New("number base is zero")
)

func formatFractional(
	number int64,
	numberBase uint,
	fractionalSize uint,
	fractionalSeparator byte,
) (string, error) {
	buffer := make([]byte, 1+fractionalSize)

	digits := buffer[1:]

	buffer[0] = fractionalSeparator

	if numberBase == 0 {
		return "", ErrNumberBaseIsZero
	}

	powered, err := safe.PowUnsigned(numberBase, fractionalSize)
	if err != nil {
		return "", ErrValueOverflow // For backward compatibility
	}

	divisor, err := safe.UnsignedToSigned[uint, int64](powered)
	if err != nil {
		return "", ErrValueOverflow // For backward compatibility
	}

	base, err := safe.UnsignedToSigned[uint, int64](numberBase)
	if err != nil {
		return "", ErrValueOverflow // For backward compatibility
	}

	// cut off an integer part of the number
	integer := number / divisor
	number -= integer * divisor

	for id := range digits {
		divisor /= base

		digit := number / divisor
		number -= digit * divisor

		// we can't get this error because above we cut off integer part of the number
		symbol, _ := digitToSymbol(digit)

		digits[id] = symbol
	}

	formated := unsafe.String(unsafe.SliceData(buffer), len(buffer))

	cleared := clearFractional(formated, fractionalSeparator)

	return cleared, nil
}
