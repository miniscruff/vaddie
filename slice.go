package vaddie

import (
	"fmt"
	"slices"
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

// SliceMaxLength validates that a slice has no more than a maximum amount of values.
func SliceMaxLength[T any](maxLength int) ValidateSlice[T] {
	return func(values []T) error {
		l := len(values)
		if l > maxLength {
			return &ValidationError{
				Message: "too long",
				Help:    fmt.Sprintf("%d > %d", l, maxLength),
			}
		}

		return nil
	}
}

// SliceUnique validates that all items in the slice are unique.
func SliceUnique[T comparable]() ValidateSlice[T] {
	return func(values []T) error {
		counts := make(map[T]int, len(values))
		for _, v := range values {
			counts[v]++
		}

		dupes := make([]T, 0)
		for v, count := range counts {
			if count > 1 {
				dupes = append(dupes, v)
			}
		}

		if len(dupes) > 0 {
			return &ValidationError{
				Message: "value found in slice more than once",
				Help:    fmt.Sprintf("%v repeated", dupes),
			}
		}

		return nil
	}
}

// SliceContains validates that the value is contained in the slice at least once.
func SliceContains[T comparable](v T) ValidateSlice[T] {
	return func(values []T) error {
		if !slices.Contains(values, v) {
			return &ValidationError{
				Message: "value not found in slice",
				Help:    fmt.Sprintf("%v missing", v),
			}
		}

		return nil
	}
}

// SliceMinContains validates that the value is contained in the slice at least a minimum set of times.
func SliceMinContains[T comparable](v T, minCount int) ValidateSlice[T] {
	return func(values []T) error {
		if len(values) < minCount {
			return &ValidationError{
				Message: "values is not long enough to contain enough values",
				Help:    fmt.Sprintf("expecting %d of value %v but only has %d values", minCount, v, len(values)),
			}
		}

		count := 0
		for _, sv := range values {
			if sv == v {
				count++
			}
		}

		if count < minCount {
			return &ValidationError{
				Message: "value not too many times",
				Help:    fmt.Sprintf("%v found %d < %d", v, count, minCount),
			}
		}

		return nil
	}
}

// SliceMaxContains validates that the value is contained in the slice at most a maximum set of times.
func SliceMaxContains[T comparable](v T, maxCount int) ValidateSlice[T] {
	return func(values []T) error {
		count := 0
		for _, sv := range values {
			if sv == v {
				count++
			}
		}

		if count > maxCount {
			return &ValidationError{
				Message: "value not found enough times",
				Help:    fmt.Sprintf("%v found %d > %d", v, count, maxCount),
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
