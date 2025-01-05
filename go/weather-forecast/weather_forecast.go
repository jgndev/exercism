// Package weather is a package for getting the current weather conditions.
package weather

// CurrentCondition is a string representing the current weather conditions.
var CurrentCondition string

// CurrentLocation is a string representing the weather location.
var CurrentLocation string

// Forecast takes a city and a condition of type string and returns a string
// with the CurrentLocation and CurrentCondition.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
