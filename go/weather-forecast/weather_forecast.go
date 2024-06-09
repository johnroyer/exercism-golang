// Package weather: test.
package weather

// CurrentCondition used by.
var CurrentCondition string

// CurrentLocation used by.
var CurrentLocation string

// Forecast will rrturn the weather by location.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
