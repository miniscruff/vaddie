package vaddie

import (
	"testing"
)

var sliceIntTests = []TestCase[[]int]{
	{
		Name:          "min",
		ValidValues:   [][]int{{15, 18}},
		InvalidValues: [][]int{{7}},
		Validation:    SliceMinLength[int](2),
	},
	{
		Name:          "max",
		ValidValues:   [][]int{{15, 18}},
		InvalidValues: [][]int{{7, 3, 18, 34}},
		Validation:    SliceMaxLength[int](2),
	},
	{
		Name:          "unique",
		ValidValues:   [][]int{{15, 18}},
		InvalidValues: [][]int{{7, 3, 18, 7}},
		Validation:    SliceUnique[int](),
	},
	{
		Name:          "contains",
		ValidValues:   [][]int{{15, 18, 24}},
		InvalidValues: [][]int{{7, 3, 18, 7}},
		Validation:    SliceContains(24),
	},
	{
		Name:        "min contains",
		ValidValues: [][]int{{15, 18, 21, 21}},
		InvalidValues: [][]int{
			{7, 3, 18, 7, 21},
			{7},
		},
		Validation: SliceMinContains(21, 2),
	},
	{
		Name:          "max contains",
		ValidValues:   [][]int{{15, 18, 23}},
		InvalidValues: [][]int{{23, 7, 3, 18, 7, 23, 23}},
		Validation:    SliceMaxContains(23, 2),
	},
}

func Test_Slice(t *testing.T) {
	for _, tc := range sliceIntTests {
		tc.Run(t)
	}
}

type sliceTestThing struct {
	X int
	Y int
}

func (t sliceTestThing) Validate() error {
	return Join(
		AllOf(t.X, "x", OrderedGte(5)),
		AllOf(t.Y, "y", OrderedLte(5)),
	)
}

var sliceGroupTests = []TestCase[[]sliceTestThing]{
	{
		Name: "all",
		ValidValues: [][]sliceTestThing{
			{
				{X: 10, Y: 3},
			},
			{},
			nil,
		},
		InvalidValues: [][]sliceTestThing{
			{
				{X: 10, Y: 10},
				{X: 3, Y: 3},
			},
		},
		Validation: func(t []sliceTestThing) error {
			return All(t, "things")
		},
	},
	{
		Name: "all with extra validation",
		ValidValues: [][]sliceTestThing{
			{
				{X: 10, Y: 3},
				{X: 10, Y: 3},
				{X: 10, Y: 3},
			},
		},
		InvalidValues: [][]sliceTestThing{
			{
				{X: 10, Y: 3},
			},
		},
		Validation: func(t []sliceTestThing) error {
			return All(t, "things", SliceMinLength[sliceTestThing](3))
		},
	},
}

func Test_SliceGroups(t *testing.T) {
	for _, tc := range sliceGroupTests {
		tc.Run(t)
	}
}

func Test_Dive(t *testing.T) {
	TestCase[[]string]{
		Name:          "dive",
		ValidValues:   [][]string{{"a", "b", "c"}},
		InvalidValues: [][]string{{"a", "b", "c", "5"}},
		Validation: func(values []string) error {
			return All(values, "things", Dive(StrLetters()))
		},
	}.Run(t)
}
