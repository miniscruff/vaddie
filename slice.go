package vaddy

import (
	"fmt"
)

// ValidateSlice can be used to validate a slice of values.
type ValidateSlice[T any] func(values []T) error

// SliceMinLength validates that a slice has at least a minimum amount of values.
func SliceMinLength[T any](minLength int) ValidateSlice[T] {
	return func(values []T) error {
		l := len(values)
		if l < minLength {
			return &ValidationError{
				Message: "not long enough",
				Help:    fmt.Sprintf("%d < %d", l, minLength),
			}
		}

		return nil
	}
}

// All can be used to validate all the items of a slice.
// If T implements the [Validator] interface, each value will also run that validation.
func All[T any](values []T, key string, validateSlice ...ValidateSlice[T]) error {
	errs := make([]error, 0)

	for i, value := range values {
		if v, isValidator := (any(value)).(Validator); isValidator {
			if err := v.Validate(); err != nil {
				errs = append(errs, expandErrorKeyIndex(err, key, i))
			}
		}
	}

	for _, validation := range validateSlice {
		err := validation(values)
		if err != nil {
			errs = append(errs, expandErrorKey(err, key))
		}
	}

	return Join(errs...)
}

// Dive can be used to dive into a slice validating the values within.
func Dive[T any](validateValues ...ValidateValue[T]) ValidateSlice[T] {
	return func(values []T) error {
		errs := make([]error, 0)

		for i, value := range values {
			for _, validator := range validateValues {
				if err := validator(value); err != nil {
					errs = append(errs, expandErrorIndex(err, i))
				}
			}
		}

		return Join(errs...)
	}
}
