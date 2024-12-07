package models

import (
	"time"

	"github.com/google/uuid"
)

type Stool struct {
	ID               uuid.UUID `json:"id"`
	EntityID         uuid.UUID `json:"entity_id"`
	ObservedAt       time.Time `json:"observed_at"`
	ConsistencyID    uuid.UUID `json:"consistency_id"`
	Color            string    `json:"color"`
	Volume           float64   `json:"volume"`
	Mood             string    `json:"mood"`
	WeatherID        uuid.UUID `json:"weather_id"`
	Exercise         bool      `json:"exercise"`
	ExerciseDuration int       `json:"exercise_duration"`
	SleepDuration    int       `json:"sleep_duration"`
	MenstrualPhase   string    `json:"menstrual_phase"`
	Notes            string    `json:"notes"`
	CreatedAt        time.Time `json:"created_at"`
}
