package period

import "golang.org/x/exp/constraints"

func safeSum[Type constraints.Integer](one Type, two Type) (Type, bool) {
	var zero Type

	sum := one + two

	switch {
	case one > zero && two > zero:
		if sum < one {
			return zero, true
		}
	case one < zero && two < zero:
		if sum > one {
			return zero, true
		}
	}

	return sum, false
}
