// Package then has assertions for our tests.
package then

import (
	"fmt"
	"strings"
	"testing"
)

// Valid checks that value passes validate.
func Valid[T any](t testing.TB, validate func(T) error, value T) {
	t.Helper()

	if err := validate(value); err != nil {
		t.Errorf("expected %q to be valid: %v", fmt.Sprint(value), strings.TrimSpace(err.Error()))
	}
}

// Invalid checks that value fails validate.
func Invalid[T any](t testing.TB, validate func(T) error, value T) {
	t.Helper()

	if err := validate(value); err == nil {
		t.Errorf("expected %q to be invalid", fmt.Sprint(value))
	}
}
