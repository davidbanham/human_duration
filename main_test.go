package human_duration

import (
	"fmt"
	"testing"
	"time"
)

type fixture struct {
	duration  time.Duration
	precision string
	result    string
}

func ExampleString() {
	duration := time.Hour*24*365 + time.Hour*8 + time.Minute*33 + time.Second*24
	fmt.Println(String(duration, Second))

	start, _ := time.Parse(time.RFC3339, "2012-11-01T22:08:41+00:00")
	stop, _ := time.Parse(time.RFC3339, "2013-12-04T23:09:42+00:00")

	fmt.Println(String(stop.Sub(start), Second))

	// Output: 1 year 8 hours 33 minutes 24 seconds
	// 1 year 33 days 1 hour 1 minute 1 second
}

func TestString(t *testing.T) {
	day := time.Hour * 24
	year := day * 365

	data := []fixture{
		{
			duration:  day + time.Minute*4 + time.Second*8,
			precision: "second",
			result:    "1 day 4 minutes 8 seconds",
		},
		{
			duration:  day + time.Minute*4 + time.Second*8,
			precision: "minute",
			result:    "1 day 4 minutes",
		},
		{
			duration:  day + time.Minute*4 + time.Second*8,
			precision: "day",
			result:    "1 day",
		},
		{
			duration:  year*4 + day*2,
			precision: "second",
			result:    "4 years 2 days",
		},
		{
			duration:  time.Minute * 60,
			precision: "hour",
			result:    "1 hour",
		},
		{
			duration:  time.Minute * 60,
			precision: "minute",
			result:    "1 hour",
		},
		{
			duration:  time.Minute * 61,
			precision: "minute",
			result:    "1 hour 1 minute",
		},
		{
			duration:  (time.Minute * 61) + (time.Second * 10),
			precision: "minute",
			result:    "1 hour 1 minute",
		},
		{
			duration:  (time.Minute * 61) + (time.Second * 10),
			precision: "second",
			result:    "1 hour 1 minute 10 seconds",
		},
		{
			duration:  time.Hour * 24,
			precision: "day",
			result:    "1 day",
		},
		{
			duration:  time.Hour * 49,
			precision: "day",
			result:    "2 days",
		},
		{
			duration:  time.Hour * 49,
			precision: "hour",
			result:    "2 days 1 hour",
		},
		{
			duration:  time.Hour*49 + time.Second,
			precision: "foobarlalala",
			result:    "2 days 1 hour 1 second",
		},
		{
			duration:  time.Hour*49 + time.Second,
			precision: "",
			result:    "2 days 1 hour 1 second",
		},
		{
			duration:  time.Minute * 61,
			precision: "hours",
			result:    "1 hour",
		},
		{
			duration:  year + day + time.Hour*2,
			precision: "hours",
			result:    "1 year 1 day 2 hours",
		},
	}

	for _, fixture := range data {
		f := fixture
		t.Run(f.result+" "+f.duration.String(), func(t *testing.T) {
			t.Parallel()
			result := String(f.duration, f.precision)
			if result != f.result {
				t.Errorf("got %s, want %s", result, f.result)
			}
		})
	}
}
