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
	// 1 year 4 weeks 5 days 1 hour 1 minute 1 second

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
		{
			duration:  day * 14,
			precision: "week",
			result:    "2 weeks",
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

func ExampleStringCeiling() {
	duration := time.Hour*24 + time.Hour*2 + time.Minute*33 + time.Second*24
	fmt.Println(StringCeiling(duration, Second, Hour))

	// Output: 26 hours 33 minutes 24 seconds
}

func ExampleShortString() {
	day := time.Hour * 24
	year := day * 365

	duration := 2*year + 2*day + 2*time.Minute + 2*time.Second

	fmt.Println(ShortString(duration, Second))

	// Output: 2y2d2m2s
}

func TestShortString(t *testing.T) {
	day := time.Hour * 24
	year := day * 365

	data := []fixture{
		{
			duration:  year + day + time.Minute + time.Second,
			precision: "second",
			result:    "1y1d1m1s",
		},
		{
			duration:  2*year + 2*day + 2*time.Minute + 2*time.Second,
			precision: "second",
			result:    "2y2d2m2s",
		},
		{
			duration:  2*year + 16*day + 2*time.Minute + 2*time.Second,
			precision: "second",
			result:    "2y2w2d2m2s",
		},
	}

	for _, fixture := range data {
		f := fixture
		t.Run(f.result+" "+f.duration.String(), func(t *testing.T) {
			t.Parallel()
			result := ShortString(f.duration, f.precision)
			if result != f.result {
				t.Errorf("got %s, want %s", result, f.result)
			}
		})
	}
}

func ExampleTimestamp() {
	duration := (25 * time.Hour) + (20 * time.Minute) + (14 * time.Second)

	fmt.Println(Timestamp(duration, "second"))
	fmt.Println(Timestamp(duration, "minute"))

	// Output: 25:20:14
	// 25:20
}

func TestTimestamp(t *testing.T) {
	data := []fixture{
		{
			duration:  time.Minute + time.Second,
			precision: Second,
			result:    "1:01",
		},
		{
			duration:  (20 * time.Minute) + (14 * time.Second),
			precision: Second,
			result:    "20:14",
		},
		{
			duration:  time.Hour + (20 * time.Minute) + (14 * time.Second),
			precision: Second,
			result:    "1:20:14",
		},
		{
			duration:  (25 * time.Hour) + (20 * time.Minute) + (14 * time.Second),
			precision: Second,
			result:    "25:20:14",
		},
	}

	for _, fixture := range data {
		f := fixture
		t.Run(f.result+" "+f.duration.String(), func(t *testing.T) {
			t.Parallel()
			result := Timestamp(f.duration, f.precision)
			if result != f.result {
				t.Errorf("got %s, want %s", result, f.result)
			}
		})
	}
}
