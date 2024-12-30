package vaddy

import (
	"testing"
)

var orderedIntTests = []TestCase[int]{
	{
		Name:          "eq",
		ValidValues:   []int{15},
		InvalidValues: []int{13},
		Validation:    OrderedEq(15),
	},
	{
		Name:          "gt",
		ValidValues:   []int{16, 17},
		InvalidValues: []int{13, 8},
		Validation:    OrderedGt(15),
	},
	{
		Name:          "gte",
		ValidValues:   []int{15, 18},
		InvalidValues: []int{12, 6},
		Validation:    OrderedGte(15),
	},
	{
		Name:          "lt",
		ValidValues:   []int{14, 8},
		InvalidValues: []int{15, 20},
		Validation:    OrderedLt(15),
	},
	{
		Name:          "lte",
		ValidValues:   []int{15, 3},
		InvalidValues: []int{27, 50},
		Validation:    OrderedLte(15),
	},
}

func Test_Ordered(t *testing.T) {
	for _, tc := range orderedIntTests {
		tc.Run(t)
	}
}
