package tuktuk

func distance(raw float64) float64 {
	d := 0.0
	for raw > 0 {
		raw -= 0.5
		d += 0.5
	}

	return d
}
