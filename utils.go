package period

import "golang.org/x/exp/constraints"

func sumSigned[Type constraints.Signed](one Type, two Type) (Type, bool) {
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

func sumUnsigned[Type constraints.Unsigned](one Type, two Type) (Type, bool) {
	var zero Type

	sum := one + two

	if sum < one {
		return zero, true
	}

	return sum, false
}
