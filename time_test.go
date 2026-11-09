package vaddie

import (
	"testing"
	"testing/synctest"
	"time"
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

var relativeTimeTests = []struct {
	Name           string
	Validation     ValidateValue[time.Time]
	ValidOffsets   []time.Duration
	InvalidOffsets []time.Duration
}{
	{
		Name:           "before now",
		Validation:     TimeBeforeNow(),
		ValidOffsets:   []time.Duration{-time.Hour, -time.Nanosecond},
		InvalidOffsets: []time.Duration{0, time.Nanosecond, time.Hour},
	},
	{
		Name:           "after now",
		Validation:     TimeAfterNow(),
		ValidOffsets:   []time.Duration{time.Nanosecond, time.Hour},
		InvalidOffsets: []time.Duration{-time.Hour, -time.Nanosecond, 0},
	},
	{
		Name:           "before from now",
		Validation:     TimeBeforeFromNow(time.Hour),
		ValidOffsets:   []time.Duration{-time.Hour, 0, time.Hour - time.Nanosecond},
		InvalidOffsets: []time.Duration{time.Hour, 2 * time.Hour},
	},
	{
		Name:           "after from now",
		Validation:     TimeAfterFromNow(time.Hour),
		ValidOffsets:   []time.Duration{time.Hour + time.Nanosecond, 2 * time.Hour},
		InvalidOffsets: []time.Duration{-time.Hour, 0, time.Hour},
	},
	{
		Name:           "before from now negative offset",
		Validation:     TimeBeforeFromNow(-time.Hour),
		ValidOffsets:   []time.Duration{-2 * time.Hour, -time.Hour - time.Nanosecond},
		InvalidOffsets: []time.Duration{-time.Hour, 0},
	},
	{
		Name:           "after from now negative offset",
		Validation:     TimeAfterFromNow(-time.Hour),
		ValidOffsets:   []time.Duration{-time.Hour + time.Nanosecond, 0},
		InvalidOffsets: []time.Duration{-2 * time.Hour, -time.Hour},
	},
}

func Test_RelativeTime(t *testing.T) {
	for _, tc := range relativeTimeTests {
		t.Run(tc.Name, func(t *testing.T) {
			// Time inside the bubble only moves on sleep, so offsets are exact.
			synctest.Test(t, func(t *testing.T) {
				now := time.Now()

				for _, offset := range tc.ValidOffsets {
					if err := tc.Validation(now.Add(offset)); err != nil {
						t.Errorf("unexpected error with offset '%v': %v", offset, err)
					}
				}

				for _, offset := range tc.InvalidOffsets {
					if err := tc.Validation(now.Add(offset)); err == nil {
						t.Errorf("expected error with invalid offset '%v'", offset)
					}
				}
			})
		})
	}
}

func Test_RelativeTimeReadsNowOnValidate(t *testing.T) {
	testCases := []struct {
		Name        string
		Validation  ValidateValue[time.Time]
		Offset      time.Duration
		ValidBefore bool
	}{
		{Name: "before now", Validation: TimeBeforeNow(), Offset: 30 * time.Minute},
		{Name: "after now", Validation: TimeAfterNow(), Offset: 30 * time.Minute, ValidBefore: true},
		{Name: "before from now", Validation: TimeBeforeFromNow(time.Hour), Offset: 90 * time.Minute},
		{
			Name:        "after from now",
			Validation:  TimeAfterFromNow(time.Hour),
			Offset:      90 * time.Minute,
			ValidBefore: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			synctest.Test(t, func(t *testing.T) {
				value := time.Now().Add(tc.Offset)

				if err := tc.Validation(value); (err == nil) != tc.ValidBefore {
					t.Errorf("before sleeping: valid want %v, got error: %v", tc.ValidBefore, err)
				}

				time.Sleep(time.Hour)

				if err := tc.Validation(value); (err == nil) == tc.ValidBefore {
					t.Errorf("after sleeping: valid want %v, got error: %v", !tc.ValidBefore, err)
				}
			})
		})
	}
}

func Test_TimeMessages(t *testing.T) {
	// Time inside a bubble starts at this instant.
	start := time.Date(2000, time.January, 1, 0, 0, 0, 0, time.UTC)
	plusTwo := time.FixedZone("UTC+2", 2*60*60)

	const (
		yearEnd  = `"1999-12-31 23:00:00 +0000 UTC"`
		midnight = `"2000-01-01 00:00:00 +0000 UTC"`
		oneAM    = `"2000-01-01 01:00:00 +0000 UTC"`
		nextDay  = `"2000-01-02 00:00:00 +0000 UTC"`
		twoDays  = `"2000-01-03 00:00:00 +0000 UTC"`
		monthAgo = `"1999-12-02 00:00:00 +0000 UTC"`
		inZone   = `"2000-01-01 02:00:00 +0200 UTC+2"`
	)

	testCases := []struct {
		Name       string
		Validation ValidateValue[time.Time]
		Value      time.Time
		Want       string
	}{
		{
			Name:       "eq",
			Validation: TimeEq(start),
			Value:      start.Add(time.Hour),
			Want:       `t times are not equal ( ` + oneAM + ` != ` + midnight + ` )`,
		},
		{
			Name:       "before",
			Validation: TimeBefore(start),
			Value:      start.Add(time.Hour),
			Want:       `t time is not before ( ` + oneAM + ` is not before ` + midnight + ` )`,
		},
		{
			Name:       "before eq",
			Validation: TimeBeforeEq(start),
			Value:      start.Add(time.Hour),
			Want:       `t time is not before or equal ( ` + oneAM + ` is after ` + midnight + ` )`,
		},
		{
			Name:       "after",
			Validation: TimeAfter(start),
			Value:      start.Add(-time.Hour),
			Want:       `t time is not after ( ` + yearEnd + ` is not after ` + midnight + ` )`,
		},
		{
			Name:       "after eq",
			Validation: TimeAfterEq(start),
			Value:      start.Add(-time.Hour),
			Want:       `t time is not after or equal ( ` + yearEnd + ` is before ` + midnight + ` )`,
		},
		{
			Name:       "before now",
			Validation: TimeBeforeNow(),
			Value:      start.Add(time.Hour),
			Want:       `t time is not before now ( ` + oneAM + ` is not before ` + midnight + ` )`,
		},
		{
			Name:       "after now",
			Validation: TimeAfterNow(),
			Value:      start.Add(-time.Hour),
			Want:       `t time is not after now ( ` + yearEnd + ` is not after ` + midnight + ` )`,
		},
		{
			Name:       "before from now",
			Validation: TimeBeforeFromNow(24 * time.Hour),
			Value:      start.Add(48 * time.Hour),
			Want:       `t time is not before 24h0m0s from now ( ` + twoDays + ` is not before ` + nextDay + ` )`,
		},
		{
			Name:       "after from now",
			Validation: TimeAfterFromNow(24 * time.Hour),
			Value:      start.Add(time.Hour),
			Want:       `t time is not after 24h0m0s from now ( ` + oneAM + ` is not after ` + nextDay + ` )`,
		},
		{
			Name:       "before ago",
			Validation: TimeBeforeFromNow(-720 * time.Hour),
			Value:      start.Add(time.Hour),
			Want:       `t time is not before 720h0m0s ago ( ` + oneAM + ` is not before ` + monthAgo + ` )`,
		},
		{
			Name:       "after ago",
			Validation: TimeAfterFromNow(-time.Hour),
			Value:      start.Add(-2 * time.Hour),
			Want: `t time is not after 1h0m0s ago ( "1999-12-31 22:00:00 +0000 UTC" is not after ` +
				yearEnd + ` )`,
		},
		{
			Name:       "bound uses the location of the value",
			Validation: TimeBeforeNow(),
			Value:      start.In(plusTwo),
			Want:       `t time is not before now ( ` + inZone + ` is not before ` + inZone + ` )`,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			synctest.Test(t, func(t *testing.T) {
				err := AllOf(tc.Value, "t", tc.Validation)
				if err == nil {
					t.Fatal("expected an error")
				}

				if err.Error() != tc.Want {
					t.Errorf("message\n got: %v\nwant: %v", err, tc.Want)
				}
			})
		})
	}
}
