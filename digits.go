package period

import (
	"errors"

	"golang.org/x/exp/constraints"
)

const (
	digitZero  = 0
	digitOne   = 1
	digitTwo   = 2
	digitThree = 3
	digitFour  = 4
	digitFive  = 5
	digitSix   = 6
	digitSeven = 7
	digitEight = 8
	digitNine  = 9
)

var (
	ErrUnexpectedDigit = errors.New("unexpected digit")
)

func symbolToDigit(symbol rune) (uint, error) {
	switch symbol {
	case '0':
		return digitZero, nil
	case '1':
		return digitOne, nil
	case '2':
		return digitTwo, nil
	case '3':
		return digitThree, nil
	case '4':
		return digitFour, nil
	case '5':
		return digitFive, nil
	case '6':
		return digitSix, nil
	case '7':
		return digitSeven, nil
	case '8':
		return digitEight, nil
	case '9':
		return digitNine, nil
	}

	return 0, ErrUnexpectedDigit
}

func digitToSymbol[Type constraints.Integer](digit Type) (byte, error) {
	switch digit {
	case digitZero:
		return '0', nil
	case digitOne:
		return '1', nil
	case digitTwo:
		return '2', nil
	case digitThree:
		return '3', nil
	case digitFour:
		return '4', nil
	case digitFive:
		return '5', nil
	case digitSix:
		return '6', nil
	case digitSeven:
		return '7', nil
	case digitEight:
		return '8', nil
	case digitNine:
		return '9', nil
	}

	return 0, ErrUnexpectedDigit
}
