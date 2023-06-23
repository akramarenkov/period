package period

import (
	"golang.org/x/exp/constraints"
)

func safeSumInt[Type constraints.Integer](first Type, second Type) (Type, bool) {
	var zero Type

	sum := first + second

	switch {
	case first > zero && second > zero:
		if sum < first {
			return zero, true
		}
	case first < zero && second < zero:
		if sum > first {
			return zero, true
		}
	}

	return sum, false
}

func safeProductInt[Type constraints.Integer](first Type, second Type) (Type, bool) {
	var zero Type

	if second == zero {
		return zero, false
	}

	if isMaxNegative(first) && second < zero {
		return zero, true
	}

	product := first * second

	if product/second != first {
		return zero, true
	}

	return product, false
}

func isMaxNegative[Type constraints.Integer](number Type) bool {
	var zero Type

	if number >= zero {
		return false
	}

	number--

	return number >= zero
}

func isMaxPositive[Type constraints.Integer](number Type) bool {
	var zero Type

	if number <= zero {
		return false
	}

	number++

	return number <= zero
}

func safeFloatToInt[Float constraints.Float, Integer constraints.Integer](
	float Float,
) (Integer, bool) {
	var zero Integer

	converted := Integer(float)
	reverted := Float(converted)

	if reverted > float && isMaxNegative(converted) {
		return zero, true
	}

	if reverted < float && isMaxPositive(converted) {
		return zero, true
	}

	if reverted > float+1 {
		return zero, true
	}

	if reverted < float-1 {
		return zero, true
	}

	return converted, false
}
