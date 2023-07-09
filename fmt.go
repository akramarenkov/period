package period

import (
	"errors"

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
	formated := make([]rune, 1+fractionalSize)

	digits := formated[1:]

	formated[0] = rune(fractionalSeparator)

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

	cleared := clearFractional(string(formated), fractionalSeparator)

	return string(cleared), nil
}
