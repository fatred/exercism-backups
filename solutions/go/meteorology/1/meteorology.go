package meteorology

import "fmt"

type TemperatureUnit int

const (
	Celsius    TemperatureUnit = 0
	Fahrenheit TemperatureUnit = 1
)

func (tu TemperatureUnit) String() string {
    units := []string{"°C", "°F"}
    return units[tu]
}

type Temperature struct {
	degree int
	unit   TemperatureUnit
}

func (tmp Temperature) String() string {
    return fmt.Sprintf("%v %v", tmp.degree, tmp.unit)
}

type SpeedUnit int

const (
	KmPerHour    SpeedUnit = 0
	MilesPerHour SpeedUnit = 1
)

func (su SpeedUnit) String() string {
    units := []string{"km/h", "mph"}
    return units[su]
}

type Speed struct {
	magnitude int
	unit      SpeedUnit
}

func (spd Speed) String() string {
    return fmt.Sprintf("%v %v", spd.magnitude, spd.unit)
}

type MeteorologyData struct {
	location      string
	temperature   Temperature
	windDirection string
	windSpeed     Speed
	humidity      int
}

func (md MeteorologyData) String() string {
    return fmt.Sprintf("%v: %v, Wind %v at %v, %v%% Humidity", md.location, md.temperature, md.windDirection, md.windSpeed, md.humidity)
}