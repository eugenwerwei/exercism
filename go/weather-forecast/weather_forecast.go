// Package weather provides utility for determining 
// the weather condition at a specific location.
package weather

// CurrentCondition stores a representration of the condition as a string.
var CurrentCondition string
// CurrentLocation stores a representation of the location as a string.
var CurrentLocation string

// Forecast returnn a string sentence about the weather condition at a specific location.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
