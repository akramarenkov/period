package period

import (
	"errors"
	"time"
)

var (
	ErrEmptyUnitModifier       = errors.New("unit modifier is empty")
	ErrInvalidUnit             = errors.New("invalid unit")
	ErrMissingUnit             = errors.New("missing unit")
	ErrMissingUnitModifier     = errors.New("unit modifier is missing")
	ErrUnexpectedUnit          = errors.New("unexpected unit")
	ErrUnitModifierIsNotUnique = errors.New("unit modifier is not unique")
)

type Unit int

const (
	UnitUnknown Unit = iota
	UnitYear
	UnitMonth
	UnitDay
	UnitHour
	UnitMinute
	UnitSecond
	UnitMillisecond
	UnitMicrosecond
	UnitNanosecond
)

const (
	validUnitsQuantity = 9
)

// Units table for custom parsing and converting to string.
//
// Must contains all Unit constants (except UnitUnknown) and
// at least one modifier for each unit of measure.
// First modifier for unit is a default modifier that used when converting to string.
//
// Default units:
//
//   - y      - years;
//   - mo     - months;
//   - d      - days;
//   - h      - hours;
//   - m      - minutes;
//   - s      - seconds;
//   - ms     - milliseconds;
//   - us, µs - microseconds;
//   - ns     - nanoseconds.
type UnitsTable map[Unit][]string

// Validates units table.
func IsValidUnitsTable(units UnitsTable) error {
	unitsQuantity := 0
	uniqueModifiers := make(map[string]struct{}, len(units))

	for unit, modifiers := range units {
		if err := isValidUnit(unit); err != nil {
			return err
		}

		unitsQuantity++

		if err := isValidModifiers(modifiers, uniqueModifiers); err != nil {
			return err
		}
	}

	if unitsQuantity != validUnitsQuantity {
		return ErrMissingUnit
	}

	return nil
}

func isValidModifiers(modifiers []string, uniqueModifiers map[string]struct{}) error {
	if len(modifiers) == 0 {
		return ErrMissingUnitModifier
	}

	for _, modifier := range modifiers {
		if len(modifier) == 0 {
			return ErrEmptyUnitModifier
		}

		if _, exists := uniqueModifiers[modifier]; exists {
			return ErrUnitModifierIsNotUnique
		}

		uniqueModifiers[modifier] = struct{}{}
	}

	return nil
}

func isValidUnit(unit Unit) error {
	switch unit {
	case UnitYear:
	case UnitMonth:
	case UnitDay:
	case UnitHour:
	case UnitMinute:
	case UnitSecond:
	case UnitMillisecond:
	case UnitMicrosecond:
	case UnitNanosecond:
	default:
		return ErrInvalidUnit
	}

	return nil
}

func isYMDUnit(unit Unit) bool {
	switch unit {
	case UnitYear:
	case UnitMonth:
	case UnitDay:
	default:
		return false
	}

	return true
}

func getDurationDimension(unit Unit) (time.Duration, error) {
	switch unit {
	case UnitHour:
		return time.Hour, nil
	case UnitMinute:
		return time.Minute, nil
	case UnitSecond:
		return time.Second, nil
	case UnitMillisecond:
		return time.Millisecond, nil
	case UnitMicrosecond:
		return time.Microsecond, nil
	case UnitNanosecond:
		return time.Nanosecond, nil
	}

	return 0, ErrUnexpectedUnit
}
