package vaddy

import (
	"fmt"
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
