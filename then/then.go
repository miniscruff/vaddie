// Package then has assertions for our tests.
package then

import (
	"fmt"
	"strings"
	"testing"
)

// Valid checks that value passes validate, name identifies the check on failure.
func Valid[T any](t testing.TB, name string, validate func(T) error, value T) {
	t.Helper()

	if err := validate(value); err != nil {
		t.Errorf("%s: expected %q to be valid: %v", name, fmt.Sprint(value), strings.TrimSpace(err.Error()))
	}
}

// Invalid checks that value fails validate, name identifies the check on failure.
func Invalid[T any](t testing.TB, name string, validate func(T) error, value T) {
	t.Helper()

	if err := validate(value); err == nil {
		t.Errorf("%s: expected %q to be invalid", name, fmt.Sprint(value))
	}
}

// Equals compares two values.
func Equals[T comparable](t testing.TB, expected, actual T) {
	t.Helper()

	if expected != actual {
		t.Errorf("expected '%v' to equal '%v'", expected, actual)
	}
}

// Nil compares a value to nil.
func Nil(t testing.TB, value any) {
	t.Helper()

	if value != nil {
		t.Errorf("expected '%v' to be nil", value)
	}
}
