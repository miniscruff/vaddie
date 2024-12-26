package vaddy

import "testing"

var strTests = TestContainer[string]{
	validCases: []TestCase[string]{
		{
			Name:       "not empty",
			Value:      "a",
			Validation: StrNotEmpty(),
		},
		{
			Name:       "min",
			Value:      "abcdefg",
			Validation: StrMin(5),
		},
		{
			Name:       "max",
			Value:      "abc",
			Validation: StrMax(5),
		},
	},
	invalidCases: []TestCase[string]{
		{
			Name:       "empty",
			Value:      "",
			Validation: StrNotEmpty(),
		},
		{
			Name:       "min",
			Value:      "ab",
			Validation: StrMin(3),
		},
		{
			Name:       "max",
			Value:      "abcde",
			Validation: StrMax(3),
		},
	},
}

func TestValues(t *testing.T) {
	strTests.Run(t, "string")
}
