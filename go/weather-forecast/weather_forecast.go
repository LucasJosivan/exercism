// The weather package provides tools to get
// the weather of your location.
package weather

// CurrentCondition represents your current
// weather condition.
var CurrentCondition string

// CurrentLocation represents your current
// location.
var CurrentLocation string

// Forecast gets the city and condition and return
// a string with the current location and current
// condition.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
