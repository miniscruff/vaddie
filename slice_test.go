package vaddie

import (
	"testing"
)

var sliceIntTests = []SliceTestCase[int]{
	{
		Name:          "eq",
		ValidValues:   [][]int{{15, 18}},
		InvalidValues: [][]int{{7}},
		Validation:    SliceMinLength[int](2),
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

var sliceGroupTests = []GroupTestCase[[]sliceTestThing]{
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

func Test_DiveValid(t *testing.T) {
	validValues := []string{"a", "b", "c"}

	err := All(validValues, "things", Dive(StrLetters()))
	if err != nil {
		t.Errorf("unexpected error for valid values: %v", err)
	}
}

func Test_DiveInvalid(t *testing.T) {
	validValues := []string{"a", "b", "c", "5"}

	err := All(validValues, "things", Dive(StrLetters()))
	if err == nil {
		t.Errorf("expected error for invalid values: %v", err)
	}
}
