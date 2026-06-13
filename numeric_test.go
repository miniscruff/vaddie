package vaddie

import (
	"testing"
)

var numericIntTests = []TestCase[int]{
	{
		Name:          "mult",
		ValidValues:   []int{15, 0, 12},
		InvalidValues: []int{13, 11},
		Validation:    NumericMultipleOf(3),
	},
	{
		Name:          "mult zero",
		InvalidValues: []int{13},
		Validation:    NumericMultipleOf(0),
	},
}

func Test_Numeric(t *testing.T) {
	for _, tc := range numericIntTests {
		tc.Run(t)
	}
}
