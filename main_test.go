package human_duration

import (
	"testing"
	"time"
)

type fixture struct {
	duration  time.Duration
	precision string
	result    string
}

func TestString(t *testing.T) {
	t.Parallel()

	day := time.Hour * 24
	year := day * 365

	data := []fixture{
		fixture{
			duration:  time.Hour*25 + time.Minute*4 + time.Second*8,
			precision: "second",
			result:    "1 day 4 minutes 8 seconds",
		},
		fixture{
			duration:  time.Hour*25 + time.Minute*4 + time.Second*8,
			precision: "minute",
			result:    "1 day 4 minutes",
		},
		fixture{
			duration:  time.Hour*25 + time.Minute*4 + time.Second*8,
			precision: "day",
			result:    "1 day",
		},
		fixture{
			duration:  year*4 + day*2,
			precision: "second",
			result:    "4 years 2 days",
		},
		fixture{
			duration:  time.Minute * 60,
			precision: "hour",
			result:    "1 hour",
		},
		fixture{
			duration:  time.Minute * 60,
			precision: "minute",
			result:    "1 hour",
		},
		fixture{
			duration:  time.Minute * 61,
			precision: "minute",
			result:    "1 hour 1 minute",
		},
		fixture{
			duration:  (time.Minute * 61) + (time.Second * 10),
			precision: "minute",
			result:    "1 hour 1 minute",
		},
		fixture{
			duration:  (time.Minute * 61) + (time.Second * 10),
			precision: "second",
			result:    "1 hour 1 minute 10 seconds",
		},
		fixture{
			duration:  time.Hour * 23,
			precision: "day",
			result:    "1 day",
		},
		fixture{
			duration:  time.Hour * 49,
			precision: "day",
			result:    "2 days",
		},
		fixture{
			duration:  time.Hour * 49,
			precision: "hour",
			result:    "2 days 1 hour",
		},
	}

	for _, fixture := range data {
		t.Run(fixture.result+fixture.duration.String(), func(t *testing.T) {
			t.Parallel()

			result := String(fixture.duration, fixture.precision)
			if result != fixture.result {
				t.Errorf("got %s, want %s", result, fixture.result)
			}
		})
	}
}
