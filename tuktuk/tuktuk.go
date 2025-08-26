package tuktuk

import "math"

func distance(raw float64) float64 {
	return math.Ceil(raw*2) / 2
}

func waitTime(sec int) float64 {
	return math.Ceil(float64(sec) / 60)
}
