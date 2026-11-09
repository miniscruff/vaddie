package vaddie

import (
	"fmt"
	"time"
)

// TimeEq validates our time is the same instant as the provided time.
// Times in different locations are equal if they represent the same instant.
func TimeEq(value time.Time) ValidateValue[time.Time] {
	return func(t time.Time) error {
		if !t.Equal(value) {
			return &ValidationError{
				Message: "times are not equal",
				Help:    fmt.Sprintf("'%v' != '%v'", t, value),
			}
		}

		return nil
	}
}

// TimeBefore validates our time comes strictly before the provided time.
func TimeBefore(before time.Time) ValidateValue[time.Time] {
	return func(t time.Time) error {
		if !t.Before(before) {
			return &ValidationError{
				Message: "time is not before",
				Help:    fmt.Sprintf("'%v' is not before '%v'", t, before),
			}
		}

		return nil
	}
}

// TimeBeforeEq validates our time comes before or is the same instant as the provided time.
func TimeBeforeEq(before time.Time) ValidateValue[time.Time] {
	return func(t time.Time) error {
		if t.After(before) {
			return &ValidationError{
				Message: "time is after",
				Help:    fmt.Sprintf("'%v' is after '%v'", t, before),
			}
		}

		return nil
	}
}

// TimeAfter validates our time comes strictly after the provided time.
func TimeAfter(after time.Time) ValidateValue[time.Time] {
	return func(t time.Time) error {
		if !t.After(after) {
			return &ValidationError{
				Message: "time is not after",
				Help:    fmt.Sprintf("'%v' is not after '%v'", t, after),
			}
		}

		return nil
	}
}

// TimeAfterEq validates our time comes after or is the same instant as the provided time.
func TimeAfterEq(after time.Time) ValidateValue[time.Time] {
	return func(t time.Time) error {
		if t.Before(after) {
			return &ValidationError{
				Message: "time is before",
				Help:    fmt.Sprintf("'%v' is before '%v'", t, after),
			}
		}

		return nil
	}
}
