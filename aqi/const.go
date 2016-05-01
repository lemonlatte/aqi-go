package aqi

type AQICategory string

const (
	PM10  AQICategory = "PM10"
	PM2_5 AQICategory = "PM2.5"
	CO    AQICategory = "CO"
	O3    AQICategory = "O3"
	NO2   AQICategory = "NO2"
	SO2   AQICategory = "SO2"
)
