package vaddie

import (
	"testing"
	"time"
)

var timeTests = []TestCase[time.Time]{
	{
		Name:          "equals",
		ValidValues:   []time.Time{time.Date(2025, time.January, 10, 5, 0, 0, 0, time.UTC)},
		InvalidValues: []time.Time{time.Date(2025, time.January, 15, 5, 0, 0, 0, time.UTC)},
		Validation:    TimeEq(time.Date(2025, time.January, 10, 5, 0, 0, 0, time.UTC)),
	},
	{
		Name:          "before",
		ValidValues:   []time.Time{time.Date(2025, time.January, 10, 5, 0, 0, 0, time.UTC)},
		InvalidValues: []time.Time{time.Date(2025, time.January, 25, 5, 0, 0, 0, time.UTC)},
		Validation:    TimeBefore(time.Date(2025, time.January, 20, 5, 0, 0, 0, time.UTC)),
	},
	{
		Name:          "after",
		ValidValues:   []time.Time{time.Date(2025, time.January, 25, 5, 0, 0, 0, time.UTC)},
		InvalidValues: []time.Time{time.Date(2025, time.January, 10, 5, 0, 0, 0, time.UTC)},
		Validation:    TimeAfter(time.Date(2025, time.January, 20, 5, 0, 0, 0, time.UTC)),
	},
}

func Test_Time(t *testing.T) {
	for _, tc := range timeTests {
		tc.Run(t)
	}
}
