package vaddy

import (
	"testing"
)

var strTests = []TestCase[string]{
	{
		Name:          "not empty",
		ValidValues:   []string{"a"},
		InvalidValues: []string{""},
		Validation:    StrNotEmpty(),
	},
	{
		Name:          "min",
		ValidValues:   []string{"abcdefg"},
		InvalidValues: []string{"ab"},
		Validation:    StrMin(5),
	},
	{
		Name:          "max",
		ValidValues:   []string{"abc"},
		InvalidValues: []string{"abcdefg"},
		Validation:    StrMax(5),
	},
	{
		Name:          "letters",
		ValidValues:   []string{"abcd"},
		InvalidValues: []string{"abcd1"},
		Validation:    StrLetters(),
	},
}

func Test_Strings(t *testing.T) {
	for _, tc := range strTests {
		tc.Run(t)
	}
}
