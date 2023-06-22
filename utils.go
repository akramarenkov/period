package period

import (
	"golang.org/x/exp/constraints"
)

func safeSum[Type constraints.Integer](first Type, second Type) (Type, bool) {
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

func safeSlowMultiplication[Type constraints.Integer](first Type, second Type) (Type, bool) {
	var zero Type

	product := zero

	multiplier := second

	if second < zero {
		multiplier = -second
	}

	for id := zero; id < multiplier; id++ {
		sum, overflow := safeSum(product, first)
		if overflow {
			return zero, true
		}

		product = sum
	}

	if second < zero {
		product = -product
	}

	return product, false
}

func safeMultiplication[Type constraints.Integer](first Type, second Type) (Type, bool) {
	var zero Type

	product := first * second

	if second == zero {
		return product, false
	}

	if product/second != first {
		return zero, true
	}

	return product, false
}
