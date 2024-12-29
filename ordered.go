package vaddy

import (
	"cmp"
	"fmt"
)

// OrderedGt validates that an ordered value as greater than a minimum value.
func OrderedGt[T cmp.Ordered](minValue T) ValidateValue[T] {
	return func(value T) error {
		if value <= minValue {
			return &ValidationError{
				Message: "too low",
				Help:    fmt.Sprintf("'%v' <= '%v'", value, minValue),
			}
		}

		return nil
	}
}

// OrderedGte validates that an ordered value as greater than or equal to a minimum value.
func OrderedGte[T cmp.Ordered](minValue T) ValidateValue[T] {
	return func(value T) error {
		if value < minValue {
			// return fmt.Errorf("value too low (%v < %v)", value, minValue)
			return &ValidationError{
				Message: "too low",
				Help:    fmt.Sprintf("'%v' < '%v'", value, minValue),
			}
		}

		return nil
	}
}

// OrderedLt validates that an ordered value as less than to a maximum value.
func OrderedLt[T cmp.Ordered](maxValue T) ValidateValue[T] {
	return func(value T) error {
		if value >= maxValue {
			return &ValidationError{
				Message: "too high",
				Help:    fmt.Sprintf("'%v' >= '%v'", value, maxValue),
			}
		}

		return nil
	}
}

// OrderedLt validates that an ordered value as less than or equal to a maximum value.
func OrderedLte[T cmp.Ordered](maxValue T) ValidateValue[T] {
	return func(value T) error {
		if value > maxValue {
			return &ValidationError{
				Message: "too high",
				Help:    fmt.Sprintf("'%v' > '%v'", value, maxValue),
			}
		}

		return nil
	}
}

// OrderedEq validates that an ordered value is equal to another value.
func OrderedEq[T cmp.Ordered](eq T) ValidateValue[T] {
	return func(value T) error {
		if value != eq {
			return &ValidationError{
				Message: "values are not equal",
				Help:    fmt.Sprintf("'%v' != '%v'", value, eq),
			}
		}

		return nil
	}
}
