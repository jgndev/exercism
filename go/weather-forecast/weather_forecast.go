// Package weather provides a foreast function for a city.
package weather

var (
	// CurrentCondition describes the current weather condition.
	CurrentCondition string
	// CurrentLocation is the name of the location for the forecast.
	CurrentLocation string
)

// Forecast takes a city and condition and returns a string of concatenated output
// of the form 'city - current weather condition:  condition'.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
