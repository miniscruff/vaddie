package vaddie

import (
	"testing"
)

var cmpTests = []TestCase[string]{
	{
		Name:          "eq",
		ValidValues:   []string{"a"},
		InvalidValues: []string{"b", "c", "d"},
		Validation:    ComparableEq("a"),
	},
	{
		Name:          "ne",
		ValidValues:   []string{"a", "b", "c"},
		InvalidValues: []string{"d"},
		Validation:    ComparableNe("d"),
	},
	{
		Name:          "contains",
		ValidValues:   []string{"a", "b", "c"},
		InvalidValues: []string{"d"},
		Validation:    ComparableContains("a", "b", "c"),
	},
}

func Test_Cmp(t *testing.T) {
	for _, tc := range cmpTests {
		tc.Run(t)
	}
}
