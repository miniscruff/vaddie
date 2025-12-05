package vaddie

import (
	"fmt"
	"slices"
)

// ComparableEq validates that two comparable values are equal.
func ComparableEq[T comparable](eq T) ValidateValue[T] {
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

// ComparableNe validates that two comparable values are not equal.
func ComparableNe[T comparable](eq T) ValidateValue[T] {
	return func(value T) error {
		if value == eq {
			return &ValidationError{
				Message: "values are equal",
				Help:    fmt.Sprintf("'%v' == '%v'", value, eq),
			}
		}

		return nil
	}
}

// ComparableContains validates a comparable value is part of a slice.
func ComparableContains[T comparable](values ...T) ValidateValue[T] {
	return func(value T) error {
		if !slices.Contains(values, value) {
			return &ValidationError{
				Message: "value is not part of slice",
				Help:    fmt.Sprintf("'%v' not in [%v]", value, values),
			}
		}

		return nil
	}
}
