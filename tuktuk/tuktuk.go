package tuktuk

import (
	"fmt"
	"math"
)

func distance(km float64) float64 {
	return math.Ceil(km*2) / 2
}

func waitMins(sec int) float64 {
	return math.Ceil(float64(sec) / 60)
}

func fare(km float64, seconds int) float64 {
	fare := 4.0*distance(km) + 1.0*waitMins(seconds)
	return math.Max(fare, 35.0)
}

func PrintFare(km float64, seconds int) string {
	return fmt.Sprintf("Ride Fare: ฿%.2f for Distance: %.1f km, Waiting Time: %.0f minutes",
		fare(km, seconds),
		distance(km),
		waitMins(seconds))
}
