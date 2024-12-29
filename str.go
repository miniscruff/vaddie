package vaddy

import (
	"fmt"
	"unicode"
)

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

func StrLetters() ValidateValue[string] {
	return func(value string) error {
		for i, v := range value {
			if !unicode.IsLetter(v) {
				return &ValidationError{
					Message: "non-letter rune",
					Index:   &i,
					Help:    string(v),
				}
			}
		}

		return nil
	}
}
