package tuktuk

import "math"

func distance(raw float64) float64 {
	return math.Ceil(raw*2) / 2
}

func waitTime(sec int) float64 {
	return math.Ceil(float64(sec) / 60)
}

func fare(km float64, seconds int) float64 {
	fare := 4.0*distance(km) + 1.0*waitTime(seconds)
	return math.Max(fare, 35.0)
}
