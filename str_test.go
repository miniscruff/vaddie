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
	{
		Name:          "ascii",
		ValidValues:   []string{"abcd"},
		InvalidValues: []string{"abcd😎"},
		Validation:    StrAscii(),
	},
	{
		Name:          "has prefix",
		ValidValues:   []string{"abcd"},
		InvalidValues: []string{"def"},
		Validation:    StrHasPrefix("abc"),
	},
	{
		Name:          "not has prefix",
		ValidValues:   []string{"abcd"},
		InvalidValues: []string{"def"},
		Validation:    StrNotHasPrefix("def"),
	},
	{
		Name:          "has suffix",
		ValidValues:   []string{"aaawxyz"},
		InvalidValues: []string{"aaabcde"},
		Validation:    StrHasSuffix("xyz"),
	},
	{
		Name:          "not has suffix",
		ValidValues:   []string{"aaabcde"},
		InvalidValues: []string{"aaawxyz"},
		Validation:    StrNotHasSuffix("xyz"),
	},
	{
		Name:          "contains",
		ValidValues:   []string{"things.com"},
		InvalidValues: []string{"without a dot"},
		Validation:    StrContains("."),
	},
	{
		Name:          "not contains",
		ValidValues:   []string{"without a dot"},
		InvalidValues: []string{"with a dot."},
		Validation:    StrNotContains("."),
	},
}

func Test_Strings(t *testing.T) {
	for _, tc := range strTests {
		tc.Run(t)
	}
}
