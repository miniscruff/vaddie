package vaddie

import (
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

	t.Run(c.Name+"_valid", func(t *testing.T) {
		for _, validValue := range c.ValidValues {
			if err := c.Validation(validValue); err != nil {
				t.Errorf("unexpected error with value '%v': %v", validValue, err)
			}
		}
	})

	t.Run(c.Name+"_invalid", func(t *testing.T) {
		for _, invalidValue := range c.InvalidValues {
			if err := c.Validation(invalidValue); err == nil {
				t.Errorf("expected error with invalid value '%v': %v", invalidValue, err)
			}
		}
	})
}

type SliceTestCase[T any] struct {
	Name          string
	ValidValues   [][]T
	InvalidValues [][]T
	Validation    ValidateSlice[T]
}

func (c SliceTestCase[T]) Run(t *testing.T) {
	t.Helper()

	t.Run(c.Name+"_valid", func(t *testing.T) {
		for _, validValue := range c.ValidValues {
			if err := c.Validation(validValue); err != nil {
				t.Errorf("unexpected error with value '%v': %v", validValue, err)
			}
		}
	})

	t.Run(c.Name+"_invalid", func(t *testing.T) {
		for _, invalidValue := range c.InvalidValues {
			if err := c.Validation(invalidValue); err == nil {
				t.Errorf("expected error with invalid value '%v': %v", invalidValue, err)
			}
		}
	})
}

type TestRunnable interface {
	Run(t *testing.T)
}

type GroupTestCase[T any] struct {
	Name          string
	ValidValues   []T
	InvalidValues []T
	Validation    func(v T) error
}

var _ TestRunnable = GroupTestCase[int]{}

func (c GroupTestCase[T]) Run(t *testing.T) {
	t.Helper()

	t.Run(c.Name+"_valid", func(t *testing.T) {
		for _, validValue := range c.ValidValues {
			if err := c.Validation(validValue); err != nil {
				t.Errorf("unexpected error with value '%v': %v", validValue, err)
			}
		}
	})

	t.Run(c.Name+"_invalid", func(t *testing.T) {
		for _, invalidValue := range c.InvalidValues {
			if err := c.Validation(invalidValue); err == nil {
				t.Errorf("expected error with invalid value '%v': %v", invalidValue, err)
			}
		}
	})
}

func toPtr[T any](value T) *T {
	return &value
}
