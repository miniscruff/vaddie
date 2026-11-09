package vaddie

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

// TimeBeforeNow validates our time comes before the current time.
// The current time is read on each validation.
func TimeBeforeNow() ValidateValue[time.Time] {
	return TimeBeforeFromNow(0)
}

// TimeAfterNow validates our time comes after the current time.
// The current time is read on each validation.
func TimeAfterNow() ValidateValue[time.Time] {
	return TimeAfterFromNow(0)
}

// TimeBeforeFromNow validates our time comes before the current time plus the offset.
// A negative offset looks into the past.
// The current time is read on each validation.
func TimeBeforeFromNow(offset time.Duration) ValidateValue[time.Time] {
	return func(t time.Time) error {
		return TimeBefore(time.Now().Add(offset))(t)
	}
}

// TimeAfterFromNow validates our time comes after the current time plus the offset.
// A negative offset looks into the past.
// The current time is read on each validation.
func TimeAfterFromNow(offset time.Duration) ValidateValue[time.Time] {
	return func(t time.Time) error {
		return TimeAfter(time.Now().Add(offset))(t)
	}
}
