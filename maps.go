package vaddie

import "fmt"


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
