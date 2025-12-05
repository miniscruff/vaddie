package vaddie

import (
	"testing"
	"time"
)

var orderedIntTests = []TestCase[int]{
	{
		Name:          "eq",
		ValidValues:   []int{15},
		InvalidValues: []int{13},
		Validation:    OrderedEq(15),
	},
	{
		Name:          "ne",
		ValidValues:   []int{12},
		InvalidValues: []int{15},
		Validation:    OrderedNe(15),
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

var durationTests = []TestCase[time.Duration]{
	{
		Name:          "equals",
		ValidValues:   []time.Duration{5 * time.Second},
		InvalidValues: []time.Duration{10 * time.Second},
		Validation:    OrderedEq(5 * time.Second),
	},
	{
		Name:          "lte",
		ValidValues:   []time.Duration{10 * time.Second},
		InvalidValues: []time.Duration{30 * time.Second},
		Validation:    OrderedLte(15 * time.Second),
	},
	{
		Name:          "gte",
		ValidValues:   []time.Duration{30 * time.Second},
		InvalidValues: []time.Duration{10 * time.Second},
		Validation:    OrderedGte(15 * time.Second),
	},
}

func Test_Duration(t *testing.T) {
	for _, tc := range durationTests {
		tc.Run(t)
	}
}
