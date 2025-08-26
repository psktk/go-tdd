package tuktuk

import "math"

func distance(raw float64) float64 {
	return math.Ceil(raw*2) / 2
}

func waitMins(sec int) float64 {
	return math.Ceil(float64(sec) / 60)
}

func fare(km float64, seconds int) float64 {
	fare := 4.0*distance(km) + 1.0*waitMins(seconds)
	return math.Max(fare, 35.0)
}
