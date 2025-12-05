package vaddie

import (
	"testing"
)

var cmpTests = []TestCase[string]{}

func Test_Cmp(t *testing.T) {
	for _, tc := range orderedIntTests {
		tc.Run(t)
	}
}
