package period

import "golang.org/x/exp/constraints"

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
