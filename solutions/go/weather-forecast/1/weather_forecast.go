// Package weather provides the forecast for a location.
package weather

// CurrentCondition is a string defining the weather conditions for the city.
var CurrentCondition string
// CurrentLocation is a string defining the city for the weather conditions.
var CurrentLocation string

/* Forecast accepts the city and the conditions as string and returns a string with the location and observed conditions.
*/
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
