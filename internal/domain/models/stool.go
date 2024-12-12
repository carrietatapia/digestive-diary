package models

import (
	"time"
)

type Stool struct {
	ID               string    `json:"id"`
	EntityID         string    `json:"entity_id"`
	ObservedAt       time.Time `json:"observed_at"`
	ConsistencyID    string    `json:"consistency_id"`
	Color            string    `json:"color"`
	Volume           float64   `json:"volume"`
	Mood             string    `json:"mood"`
	WeatherID        string    `json:"weather_id"`
	Exercise         bool      `json:"exercise"`
	ExerciseDuration int       `json:"exercise_duration"`
	SleepDuration    int       `json:"sleep_duration"`
	MenstrualPhase   string    `json:"menstrual_phase"`
	Notes            string    `json:"notes"`
	CreatedAt        time.Time `json:"created_at"`
}
