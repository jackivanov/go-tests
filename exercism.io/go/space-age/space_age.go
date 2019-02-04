package space

// Planet is planet's name string
type Planet string

// YearOnEarth a year on Earh in seconds
const YearOnEarth float64 = 31557600

// Age returns Earth-years old age (float64) for given planet (Planet) and age in seconds (float64)
func Age(age float64, planet Planet) float64 {
	// AgeOnPlanets a map of planets and Earth-years orbital period
	AgeOnPlanets := map[Planet]float64{
		"Earth":   1,
		"Mercury": 0.2408467,
		"Venus":   0.61519726,
		"Mars":    1.8808158,
		"Jupiter": 11.862615,
		"Saturn":  29.447498,
		"Uranus":  84.016846,
		"Neptune": 164.79132,
	}

	return age / (YearOnEarth * AgeOnPlanets[planet])
}
