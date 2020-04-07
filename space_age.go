package space

import "math"

// Planet is the type for the planet input
type Planet string

// Age converts the age according with the input planet
func Age(seconds float64, planet Planet) float64 {
	return secondsToYears(seconds, planetYearConverter(planet))
}

func planetYearConverter(planet Planet) float64 {
	switch planet {
	case "Mercury":
		return 0.2408467
	case "Venus":
		return 0.61519726
	case "Mars":
		return 1.8808158
	case "Jupiter":
		return 11.862615
	case "Saturn":
		return 29.447498
	case "Uranus":
		return 84.016846
	case "Neptune":
		return 164.79132
	default:
		return 1
	}
}

func secondsToYears(seconds float64, converter float64) float64 {
	return math.Round((seconds/60/60/24/365.25)/converter*100) / 100
}
