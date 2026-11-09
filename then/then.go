// Package then has assertions for our tests.
package then

import "testing"

// Nil compares a value to nil.
func Nil(t testing.TB, value any) {
	t.Helper()

	if value != nil {
		t.Logf("expected '%v' to be nil", value)
		t.FailNow()
	}
}

// NotNil compares a value is not nil.
func NotNil(t testing.TB, value any) {
	t.Helper()

	if value == nil {
		t.Logf("expected '%v' not to be nil", value)
		t.FailNow()
	}
}
