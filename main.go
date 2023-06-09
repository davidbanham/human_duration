// Package human_duration provides human readable output of
// time.Duration
package human_duration

import (
	"fmt"
	"math"
	"regexp"
	"strings"
	"time"
)

// Available precisions
const (
	Second = "second"
	Minute = "minute"
	Hour   = "hour"
	Day    = "day"
	Week   = "week"
	Year   = "year"
)

func precisionToDuration(precision string) time.Duration {
	switch precision {
	case Second:
		return time.Second
	case Minute:
		return time.Minute
	case Hour:
		return time.Hour
	case Day:
		return time.Hour * 24
	case Week:
		return time.Hour * 24 * 7
	case Year:
		return time.Hour * 24 * 365
	default:
		return time.Nanosecond
	}
}

// String converts duration to human readable format, according to precision.
func String(duration time.Duration, precision string) string {
	return StringCeiling(duration, precision, "")
}

func StringCeiling(duration time.Duration, precision, ceiling string) string {
	return StringCeilingPadded(duration, precision, ceiling, false)
}

type chunk struct {
	singularName string
	amount       int64
}

func StringCeilingPadded(duration time.Duration, precision, ceiling string, padded bool) string {
	years := int64(duration.Hours() / 24 / 365)
	weeks := int64(math.Mod(float64(int64(duration.Hours()/24/7)), 52))
	days := int64(math.Mod(float64(int64(duration.Hours()/24)), 365)) - (weeks * 7)
	hours := int64(math.Mod(duration.Hours(), 24))
	minutes := int64(math.Mod(duration.Minutes(), 60))
	seconds := int64(math.Mod(duration.Seconds(), 60))

	switch ceiling {
	case Second:
		seconds = int64(duration.Seconds())
		minutes = 0
		hours = 0
		days = 0
		years = 0
	case Minute:
		minutes = int64(duration.Minutes())
		hours = 0
		days = 0
		years = 0
	case Hour:
		hours = int64(duration.Hours())
		days = 0
		years = 0
	case Day:
		days = int64(float64(int64(duration.Hours() / 24)))
		years = 0
		weeks = 0
	case Week:
		weeks = int64(float64(int64(duration.Hours() / 24 / 7)))
		years = 0
	}

	chunks := []chunk{
		{"year", years},
		{"week", weeks},
		{"day", days},
		{"hour", hours},
		{"minute", minutes},
		{"second", seconds},
	}

	parts := []string{}
	preciseEnough := false
	isLeading := true

	unpaddedNumberFormat := "%d"
	paddedNumberFormat := "%02d"

	hitSomething := false
	var targetChunk chunk

	for _, chunk := range chunks {
		if preciseEnough {
			continue
		}

		if chunk.amount != 0 {
			hitSomething = true
		}

		if chunk.singularName == precision || chunk.singularName+"s" == precision {
			targetChunk = chunk
			preciseEnough = true
		}

		numberFormat := unpaddedNumberFormat
		if chunk.amount > 0 && isLeading {
			isLeading = false
		} else if padded {
			numberFormat = paddedNumberFormat
		}

		switch chunk.amount {
		case 0:
			continue
		case 1:
			parts = append(parts, fmt.Sprintf(numberFormat+" %s", chunk.amount, chunk.singularName))
		default:
			parts = append(parts, fmt.Sprintf(numberFormat+" %ss", chunk.amount, chunk.singularName))
		}
	}

	if !hitSomething {
		return "less than a " + targetChunk.singularName
	}

	return strings.Join(parts, " ")
}

// ShortString converts duration to a shortened human readable format, according to precision.
func ShortString(duration time.Duration, precision string) string {
	str := String(duration, precision)
	str = shorten(str)
	return str
}

func shorten(str string) string {
	str = strings.Replace(str, " ", "", -1)
	str = strings.Replace(str, "years", "y", 1)
	str = strings.Replace(str, "year", "y", 1)
	str = strings.Replace(str, "days", "d", 1)
	str = strings.Replace(str, "day", "d", 1)
	str = strings.Replace(str, "weeks", "w", 1)
	str = strings.Replace(str, "week", "w", 1)
	str = strings.Replace(str, "hours", "h", 1)
	str = strings.Replace(str, "hour", "h", 1)
	str = strings.Replace(str, "minutes", "m", 1)
	str = strings.Replace(str, "minute", "m", 1)
	str = strings.Replace(str, "seconds", "s", 1)
	str = strings.Replace(str, "second", "s", 1)
	return str
}

var trailingColon = regexp.MustCompile(`:$`)

// Timestamp converts duration to a common timestamp format, often used for videos.
func Timestamp(interval time.Duration, precision string) string {
	if precisionToDuration(precision) > time.Hour {
		precision = "hours"
	}

	str := shorten(StringCeilingPadded(interval, precision, "hour", true))
	str = strings.Replace(str, "h", ":", 1)
	str = strings.Replace(str, "m", ":", 1)
	str = strings.Replace(str, "s", "", 1)
	str = trailingColon.ReplaceAllString(str, "")
	if !strings.Contains(str, ":") {
		str = "0:" + str
	}

	return str
}
