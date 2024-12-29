package vaddy

import (
	"fmt"
	"unicode"
)

// StrNotEmpty validates that a given string is not empty.
func StrNotEmpty() ValidateValue[string] {
	return func(value string) error {
		if value == "" {
			return &ValidationError{
				Message: "is empty",
			}
		}

		return nil
	}
}

// StrMin validates our value is at least a minimum length.
func StrMin(minLength int) ValidateValue[string] {
	return func(value string) error {
		length := len(value)
		if length < minLength {
			return &ValidationError{
				Message: "too short",
				Help:    fmt.Sprintf("%d < %d", length, minLength),
			}
		}

		return nil
	}
}

// StrMax validates our value is no more then a maximum length.
func StrMax(maxLength int) ValidateValue[string] {
	return func(value string) error {
		length := len(value)
		if length > maxLength {
			return &ValidationError{
				Message: "too long",
				Help:    fmt.Sprintf("%d > %d", length, maxLength),
			}
		}

		return nil
	}
}

// StrLetters validates every rune is a letter.
func StrLetters() ValidateValue[string] {
	return func(value string) error {
		for i, v := range value {
			if !unicode.IsLetter(v) {
				return &ValidationError{
					Message: "non-letter rune",
					Index:   &i, // TODO: index should be for slices, we should move this to help
					Help:    string(v),
				}
			}
		}

		return nil
	}
}
