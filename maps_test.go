package vaddie

import (
	"testing"
)

var mapStringIntTests = []TestCase[map[string]int]{
	{
		Name: "min",
		ValidValues: []map[string]int{
			{
				"a": 5,
				"b": 3,
				"c": 1,
			},
		},
		InvalidValues: []map[string]int{
			{
				"a": 5,
			},
		},
		Validation: MapMinLength[string, int](3),
	},
	{
		Name: "max",
		ValidValues: []map[string]int{
			{
				"a": 5,
			},
		},
		InvalidValues: []map[string]int{
			{
				"a": 5,
				"b": 3,
				"c": 1,
			},
		},
		Validation: MapMaxLength[string, int](2),
	},
	{
		Name: "required",
		ValidValues: []map[string]int{
			{
				"a": 5,
				"b": 3,
			},
		},
		InvalidValues: []map[string]int{
			{
				"a": 5,
			},
		},
		Validation: MapRequiredKeys[string, int]("a", "b"),
	},
}

func Test_MapStringInt(t *testing.T) {
	for _, tc := range mapStringIntTests {
		tc.Run(t)
	}
}
