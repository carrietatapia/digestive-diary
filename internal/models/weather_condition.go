package models

import (
	"time"

	"github.com/google/uuid"
)

type WeatherCondition struct {
	ID               uuid.UUID `json:"id"`
	Timestamp        time.Time `json:"timestamp"`
	Temperature      float64   `json:"temperature"`
	WeatherCondition string    `json:"weather_condition"`
	Humidity         float64   `json:"humidity"`
	WindSpeed        float64   `json:"wind_speed"`
	CreatedAt        time.Time `json:"created_at"`
}
