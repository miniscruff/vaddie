package vaddy

import (
	"fmt"
	"testing"
)

type TestCase[T any] struct {
	Name          string
	ValidValues   []T
	InvalidValues []T
	Validation    ValidateValue[T]
}

func (c TestCase[T]) Run(t *testing.T) {
	t.Helper()

	var emptyT T
	typeName := fmt.Sprintf("%T", emptyT)

	t.Run(typeName+"valid", func(t *testing.T) {
		for _, validValue := range c.ValidValues {
			if err := c.Validation(validValue); err != nil {
				t.Errorf("unexpected error with value '%v': %v", validValue, err)
			}
		}
	})

	t.Run(typeName+"invalid", func(t *testing.T) {
		for _, invalidValue := range c.InvalidValues {
			if err := c.Validation(invalidValue); err == nil {
				t.Errorf("expected error with invalid value '%v': %v", invalidValue, err)
			}
		}
	})
}
