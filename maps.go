package vaddie

import (
	"fmt"
	"maps"
	"slices"
)

// MapMinLength validates our map has at least a minimum set of properties.
func MapMinLength[K comparable, V any](minLength int) ValidateValue[map[K]V] {
	return func(value map[K]V) error {
		l := len(value)
		if l < minLength {
			return &ValidationError{
				Message: "not long enough",
				Help:    fmt.Sprintf("%d < %d", l, minLength),
			}
		}

		return nil
	}
}

// MapMaxLength validates our map has no more than a maximum set of properties.
func MapMaxLength[K comparable, V any](maxLength int) ValidateValue[map[K]V] {
	return func(value map[K]V) error {
		l := len(value)
		if l > maxLength {
			return &ValidationError{
				Message: "too long",
				Help:    fmt.Sprintf("%d > %d", l, maxLength),
			}
		}

		return nil
	}
}

// MapRequiredKeys validates our map has a set of required keys.
// Values of the keys are not validated separately other than existing.
func MapRequiredKeys[K comparable, V any](reqs ...K) ValidateValue[map[K]V] {
	return func(value map[K]V) error {
		reqLookup := make(map[K]struct{}, len(reqs))
		for _, v := range reqs {
			reqLookup[v] = struct{}{}
		}

		for k := range value {
			delete(reqLookup, k)
		}

		if len(reqLookup) != 0 {
			return &ValidationError{
				Message: "required keys not found in map",
				Help:    fmt.Sprintf("%v are missing from map", slices.Collect(maps.Keys(reqLookup))),
			}
		}

		return nil
	}
}
