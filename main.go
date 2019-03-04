package human_duration

import (
	"fmt"
	"math"
	"strings"
	"time"
)

func String(duration time.Duration, precision string) string {
	years := int64(duration.Hours() / 24 / 365)
	days := int64(duration.Hours() / 24)
	hours := int64(math.Mod(duration.Hours(), 24))
	minutes := int64(math.Mod(duration.Minutes(), 60))
	seconds := int64(math.Mod(duration.Seconds(), 60))

	chunks := []struct {
		singularName string
		amount       int64
	}{
		{"year", years},
		{"day", days},
		{"hour", hours},
		{"minute", minutes},
		{"second", seconds},
	}

	parts := []string{}

	preciseEnough := false

	for _, chunk := range chunks {
		if preciseEnough {
			continue
		}
		if chunk.singularName == precision || chunk.singularName+"s" == precision {
			preciseEnough = true
		}
		switch chunk.amount {
		case 0:
			continue
		case 1:
			parts = append(parts, fmt.Sprintf("%d %s", chunk.amount, chunk.singularName))
		default:
			parts = append(parts, fmt.Sprintf("%d %ss", chunk.amount, chunk.singularName))
		}
	}
	return strings.Join(parts, " ")
}
