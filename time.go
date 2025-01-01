package vaddy

import (
	"fmt"
	"time"
)

// TimeEq validates our time is equal to the provided time.
func TimeEq(value time.Time) ValidateValue[time.Time] {
	return func(t time.Time) error {
		if value != t {
			return &ValidationError{
				Message: "times are not equal",
				Help:    fmt.Sprintf("'%v' != '%v'", t, value),
			}
		}

		return nil
	}
}

// TimeBefore validates our time comes before the provided time.
func TimeBefore(before time.Time) ValidateValue[time.Time] {
	return func(t time.Time) error {
		if t.After(before) {
			return &ValidationError{
				Message: "time is not before",
				Help:    fmt.Sprintf("'%v' is after '%v'", t, before),
			}
		}

		return nil
	}
}

// TimeAfter validates our time comes after the provided time.
func TimeAfter(after time.Time) ValidateValue[time.Time] {
	return func(t time.Time) error {
		if t.Before(after) {
			return &ValidationError{
				Message: "time is not after",
				Help:    fmt.Sprintf("'%v' is before '%v'", t, after),
			}
		}

		return nil
	}
}
