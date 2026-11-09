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
				Help:    fmt.Sprintf("%q != %q", t, value),
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
				Help:    fmt.Sprintf("%q is not before %q", t, before),
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
				Message: "time is not before or equal",
				Help:    fmt.Sprintf("%q is after %q", t, before),
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
				Help:    fmt.Sprintf("%q is not after %q", t, after),
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
				Message: "time is not after or equal",
				Help:    fmt.Sprintf("%q is before %q", t, after),
			}
		}

		return nil
	}
}

// TimeBeforeNow validates our time comes before the current time.
// The current time is read on each validation.
func TimeBeforeNow() ValidateValue[time.Time] {
	return func(t time.Time) error {
		now := time.Now()
		if !t.Before(now) {
			return &ValidationError{
				Message: "time is not before now",
				Help:    fmt.Sprintf("%q is not before %q", t, displayBound(now, t)),
			}
		}

		return nil
	}
}

// TimeAfterNow validates our time comes after the current time.
// The current time is read on each validation.
func TimeAfterNow() ValidateValue[time.Time] {
	return func(t time.Time) error {
		now := time.Now()
		if !t.After(now) {
			return &ValidationError{
				Message: "time is not after now",
				Help:    fmt.Sprintf("%q is not after %q", t, displayBound(now, t)),
			}
		}

		return nil
	}
}

// TimeBeforeFromNow validates our time comes before the current time plus the offset.
// A negative offset looks into the past.
// The current time is read on each validation.
func TimeBeforeFromNow(offset time.Duration) ValidateValue[time.Time] {
	return func(t time.Time) error {
		bound := time.Now().Add(offset)
		if !t.Before(bound) {
			return &ValidationError{
				Message: "time is not before " + describeOffset(offset),
				Help:    fmt.Sprintf("%q is not before %q", t, displayBound(bound, t)),
			}
		}

		return nil
	}
}

// TimeAfterFromNow validates our time comes after the current time plus the offset.
// A negative offset looks into the past.
// The current time is read on each validation.
func TimeAfterFromNow(offset time.Duration) ValidateValue[time.Time] {
	return func(t time.Time) error {
		bound := time.Now().Add(offset)
		if !t.After(bound) {
			return &ValidationError{
				Message: "time is not after " + describeOffset(offset),
				Help:    fmt.Sprintf("%q is not after %q", t, displayBound(bound, t)),
			}
		}

		return nil
	}
}

// describeOffset words a relative bound for an error message.
func describeOffset(offset time.Duration) string {
	if offset < 0 {
		return (-offset).String() + " ago"
	}

	return offset.String() + " from now"
}

// displayBound prepares a bound read from the clock to be printed next to value,
// in the same location, without the monotonic reading or sub second noise.
func displayBound(bound, value time.Time) time.Time {
	return bound.In(value.Location()).Truncate(time.Second)
}
