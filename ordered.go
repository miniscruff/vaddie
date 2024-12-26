package vaddy

import (
	"cmp"
	"fmt"
)

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
