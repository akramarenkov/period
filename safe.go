package period

import (
	"errors"

	"golang.org/x/exp/constraints"
)

var (
	ErrValueOverflow = errors.New("value overflow")
)

func safeSumInt[Type constraints.Integer](first Type, second Type) (Type, error) {
	sum := first + second

	switch {
	case first > 0 && second > 0:
		if sum < first {
			return 0, ErrValueOverflow
		}
	case first < 0 && second < 0:
		if sum > first {
			return 0, ErrValueOverflow
		}
	}

	return sum, nil
}

func safeProductInt[Type constraints.Integer](first Type, second Type) (Type, error) {
	if second == 0 {
		return 0, nil
	}

	if isMaxNegative(first) && second < 0 {
		return 0, ErrValueOverflow
	}

	product := first * second

	if product/second != first {
		return 0, ErrValueOverflow
	}

	return product, nil
}

func isMaxNegative[Type constraints.Integer](number Type) bool {
	if number >= 0 {
		return false
	}

	number--

	return number >= 0
}

func isMaxPositive[Type constraints.Integer](number Type) bool {
	if number <= 0 {
		return false
	}

	number++

	return number <= 0
}

func safeFloatToInt[Float constraints.Float, Integer constraints.Integer](
	float Float,
) (Integer, error) {
	converted := Integer(float)
	reverted := Float(converted)

	if reverted > float && isMaxNegative(converted) {
		return 0, ErrValueOverflow
	}

	if reverted < float && isMaxPositive(converted) {
		return 0, ErrValueOverflow
	}

	if reverted > float+1 {
		return 0, ErrValueOverflow
	}

	if reverted < float-1 {
		return 0, ErrValueOverflow
	}

	return converted, nil
}

func safeInvertInt[Type constraints.Integer](number Type) (Type, error) {
	if isMaxNegative(number) {
		return 0, ErrValueOverflow
	}

	return -number, nil
}

func safePowUint[Type constraints.Unsigned](base Type, exponent Type) (Type, error) {
	if base == 0 {
		return 0, nil
	}

	if base == 1 {
		return 1, nil
	}

	if exponent == 0 {
		return 1, nil
	}

	if exponent == 1 {
		return base, nil
	}

	powered := base

	for stage := Type(1); stage < exponent; stage++ {
		product, err := safeProductInt(powered, base)
		if err != nil {
			return 0, err
		}

		powered = product
	}

	return powered, nil
}

func safeUintToInt[
	Unsigned constraints.Unsigned,
	Signed constraints.Signed,
](number Unsigned) (Signed, error) {
	converted := Signed(number)

	if converted < 0 {
		return 0, ErrValueOverflow
	}

	reverted := Unsigned(converted)

	if reverted != number {
		return 0, ErrValueOverflow
	}

	return converted, nil
}
