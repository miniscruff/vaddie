package vaddie

import (
	"testing"
	"testing/synctest"
	"time"

	"github.com/miniscruff/vaddie/then"
)

var timeTests = []TestCase[time.Time]{
	{
		Name: "equals",
		ValidValues: []time.Time{
			time.Date(2025, time.January, 10, 5, 0, 0, 0, time.UTC),
			time.Date(2025, time.January, 10, 7, 0, 0, 0, time.FixedZone("UTC+2", 2*60*60)),
		},
		InvalidValues: []time.Time{
			time.Date(2025, time.January, 15, 5, 0, 0, 0, time.UTC),
			time.Date(2025, time.January, 10, 5, 0, 0, 0, time.FixedZone("UTC+2", 2*60*60)),
		},
		Validation: TimeEq(time.Date(2025, time.January, 10, 5, 0, 0, 0, time.UTC)),
	},
	{
		Name: "before",
		ValidValues: []time.Time{
			time.Date(2025, time.January, 10, 5, 0, 0, 0, time.UTC),
		},
		InvalidValues: []time.Time{
			time.Date(2025, time.January, 20, 5, 0, 0, 0, time.UTC),
			time.Date(2025, time.January, 25, 5, 0, 0, 0, time.UTC),
		},
		Validation: TimeBefore(time.Date(2025, time.January, 20, 5, 0, 0, 0, time.UTC)),
	},
	{
		Name: "before or equal",
		ValidValues: []time.Time{
			time.Date(2025, time.January, 10, 5, 0, 0, 0, time.UTC),
			time.Date(2025, time.January, 20, 5, 0, 0, 0, time.UTC),
		},
		InvalidValues: []time.Time{
			time.Date(2025, time.January, 25, 5, 0, 0, 0, time.UTC),
		},
		Validation: TimeBeforeEq(time.Date(2025, time.January, 20, 5, 0, 0, 0, time.UTC)),
	},
	{
		Name: "after",
		ValidValues: []time.Time{
			time.Date(2025, time.January, 25, 5, 0, 0, 0, time.UTC),
		},
		InvalidValues: []time.Time{
			time.Date(2025, time.January, 10, 5, 0, 0, 0, time.UTC),
			time.Date(2025, time.January, 20, 5, 0, 0, 0, time.UTC),
		},
		Validation: TimeAfter(time.Date(2025, time.January, 20, 5, 0, 0, 0, time.UTC)),
	},
	{
		Name: "after or equal",
		ValidValues: []time.Time{
			time.Date(2025, time.January, 20, 5, 0, 0, 0, time.UTC),
			time.Date(2025, time.January, 25, 5, 0, 0, 0, time.UTC),
		},
		InvalidValues: []time.Time{
			time.Date(2025, time.January, 10, 5, 0, 0, 0, time.UTC),
		},
		Validation: TimeAfterEq(time.Date(2025, time.January, 20, 5, 0, 0, 0, time.UTC)),
	},
}

func Test_Time(t *testing.T) {
	for _, tc := range timeTests {
		tc.Run(t)
	}
}

func relativeTimeTests(now time.Time) []TestCase[time.Time] {
	return []TestCase[time.Time]{
		{
			Name: "before now",
			ValidValues: []time.Time{
				now.Add(-time.Hour),
				now.Add(-time.Nanosecond),
			},
			InvalidValues: []time.Time{
				now,
				now.Add(time.Nanosecond),
				now.Add(time.Hour),
			},
			Validation: TimeBeforeNow(),
		},
		{
			Name: "after now",
			ValidValues: []time.Time{
				now.Add(time.Nanosecond),
				now.Add(time.Hour),
			},
			InvalidValues: []time.Time{
				now.Add(-time.Hour),
				now.Add(-time.Nanosecond),
				now,
			},
			Validation: TimeAfterNow(),
		},
		{
			Name: "before from now",
			ValidValues: []time.Time{
				now.Add(-time.Hour),
				now,
				now.Add(time.Hour - time.Nanosecond),
			},
			InvalidValues: []time.Time{
				now.Add(time.Hour),
				now.Add(2 * time.Hour),
			},
			Validation: TimeBeforeFromNow(time.Hour),
		},
		{
			Name: "after from now",
			ValidValues: []time.Time{
				now.Add(time.Hour + time.Nanosecond),
				now.Add(2 * time.Hour),
			},
			InvalidValues: []time.Time{
				now.Add(-time.Hour),
				now,
				now.Add(time.Hour),
			},
			Validation: TimeAfterFromNow(time.Hour),
		},
		{
			Name: "before from now negative offset",
			ValidValues: []time.Time{
				now.Add(-2 * time.Hour),
				now.Add(-time.Hour - time.Nanosecond),
			},
			InvalidValues: []time.Time{
				now.Add(-time.Hour),
				now,
			},
			Validation: TimeBeforeFromNow(-time.Hour),
		},
		{
			Name: "after from now negative offset",
			ValidValues: []time.Time{
				now.Add(-time.Hour + time.Nanosecond),
				now,
			},
			InvalidValues: []time.Time{
				now.Add(-2 * time.Hour),
				now.Add(-time.Hour),
			},
			Validation: TimeAfterFromNow(-time.Hour),
		},
	}
}

func Test_RelativeTime(t *testing.T) {
	// Time inside the bubble only moves on sleep, so boundaries are exact.
	synctest.Test(t, func(t *testing.T) {
		for _, tc := range relativeTimeTests(time.Now()) {
			tc.Run(t)
		}
	})
}

func Test_RelativeTimeReadsNowOnValidate(t *testing.T) {
	synctest.Test(t, func(t *testing.T) {
		value := time.Now().Add(90 * time.Minute)
		beforeNow := TimeBeforeNow()
		afterNow := TimeAfterNow()
		beforeFromNow := TimeBeforeFromNow(time.Hour)
		afterFromNow := TimeAfterFromNow(time.Hour)

		then.Invalid(t, "before now", beforeNow, value)
		then.Valid(t, "after now", afterNow, value)
		then.Invalid(t, "before from now", beforeFromNow, value)
		then.Valid(t, "after from now", afterFromNow, value)

		time.Sleep(2 * time.Hour)

		then.Valid(t, "before now", beforeNow, value)
		then.Invalid(t, "after now", afterNow, value)
		then.Valid(t, "before from now", beforeFromNow, value)
		then.Invalid(t, "after from now", afterFromNow, value)
	})
}
