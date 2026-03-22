// Package weather returns weather forecast for a specific city.
package weather

var (
	// CurrentCondition represents the current weather condition.
	CurrentCondition string
	// CurrentLocation represents the current location for the forecast.
	CurrentLocation  string
)

// Forecast returns a formatted string with the weather forecast for a specific city.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}