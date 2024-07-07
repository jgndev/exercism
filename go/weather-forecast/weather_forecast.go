// Package weather provides a Forecast method and two public variables for
// CurrentCondition and CurrentLocation respectively.
package weather

// CurrentCondition is a string representing the current conditions.
var CurrentCondition string

// CurrentLocation is a string representing the current location.
var CurrentLocation string

// Forecast takes a city and condition as strings and returns a string that is
// the result of CurrentLocation + message + CurrentCondition.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
