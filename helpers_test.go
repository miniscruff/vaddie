package vaddie

import (
	"testing"

	"github.com/miniscruff/vaddie/then"
)

type TestCase[T any] struct {
	Name          string
	ValidValues   []T
	InvalidValues []T
	Validation    func(T) error
}

func (c TestCase[T]) Run(t *testing.T) {
	t.Helper()

	for _, value := range c.ValidValues {
		then.Valid(t, c.Name, c.Validation, value)
	}

	for _, value := range c.InvalidValues {
		then.Invalid(t, c.Name, c.Validation, value)
	}
}
