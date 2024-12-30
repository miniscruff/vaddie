package vaddy

import (
	"testing"
)

type groupTestThing struct {
	X int
	Y int
}

var groupingTests = []GroupTestCase[groupTestThing]{
	{
		Name: "join and all of",
		ValidValues: []groupTestThing{
			{X: 10, Y: 10},
		},
		InvalidValues: []groupTestThing{
			{X: 5, Y: 5},
		},
		Validation: func(v groupTestThing) error {
			return Join(
				AllOf(v.X, "x", OrderedEq(10)),
				AllOf(v.Y, "y", OrderedEq(10)),
			)
		},
	},
	{
		Name: "join and one of",
		ValidValues: []groupTestThing{
			{X: 10, Y: 10},
		},
		InvalidValues: []groupTestThing{
			{X: 5, Y: 5},
		},
		Validation: func(v groupTestThing) error {
			return Join(
				OneOf(v.X, "x", OrderedEq(11), OrderedEq(10)),
				OneOf(v.Y, "y", OrderedEq(15), OrderedEq(10)),
			)
		},
	},
	{
		Name: "with and",
		ValidValues: []groupTestThing{
			{X: 10, Y: 10},
		},
		InvalidValues: []groupTestThing{
			{X: 3, Y: 50},
		},
		Validation: func(v groupTestThing) error {
			return Join(
				AllOf(v.X, "x", And(OrderedGte(5), OrderedLte(15))),
				AllOf(v.Y, "y", And(OrderedGte(5), OrderedLte(15))),
			)
		},
	},
	{
		Name: "with or",
		ValidValues: []groupTestThing{
			{X: 3, Y: 50},
		},
		InvalidValues: []groupTestThing{
			{X: 10, Y: 10},
		},
		Validation: func(v groupTestThing) error {
			return Join(
				AllOf(v.X, "x", Or(OrderedGte(15), OrderedLte(5))),
				AllOf(v.Y, "y", Or(OrderedGte(15), OrderedLte(5))),
			)
		},
	},
}

func Test_Grouping(t *testing.T) {
	for _, tc := range groupingTests {
		tc.Run(t)
	}
}
