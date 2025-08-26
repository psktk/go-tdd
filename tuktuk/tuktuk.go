package tuktuk

import "math"

func distance(raw float64) float64 {
	return math.Ceil(raw*2) / 2
}

func waitTime(raw float64) float64 {
	return math.Ceil(raw)
}
