package vaddie

import "testing"

var ifTests = []GroupTestCase[int]{
	{
		Name:          "if true",
		ValidValues:   []int{13, 18, 5},
		InvalidValues: []int{3, 1},
		Validation: func(v int) error {
			return IfExp(true,
				AllOf(v, "v", OrderedGte(5)),
			)
		},
	},
	{
		Name:        "if false",
		ValidValues: []int{2, 5, 18},
		Validation: func(v int) error {
			return IfExp(false,
				AllOf(v, "v", OrderedGte(5)),
			)
		},
	},
}

func Test_Ifs(t *testing.T) {
	for _, tc := range ifTests {
		tc.Run(t)
	}
}
