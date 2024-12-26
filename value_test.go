package vaddy

import (
	"strings"
	"testing"
)

type TestContainer[T any] struct {
	validCases   []TestCase[T]
	invalidCases []TestCase[T]
}

func (c TestContainer[T]) Run(t *testing.T, typeName string) {
	t.Helper()
	t.Parallel()

	for _, tc := range c.validCases {
		t.Run(
			strings.Join([]string{typeName, "valid", tc.Name}, "_"),
			func(t *testing.T) {
				if err := tc.Validation(tc.Value); err != nil {
					t.Errorf("unexpected error: %v", err)
				}
			},
		)
	}

	for _, tc := range c.invalidCases {
		t.Run(
			strings.Join([]string{typeName, "invalid", tc.Name}, "_"),
			func(t *testing.T) {
				if err := tc.Validation(tc.Value); err == nil {
					t.Error("expected error but got nil")
				}
			},
		)
	}
}

type TestCase[T any] struct {
	Name       string
	Value      T
	Validation ValidateValue[T]
}
