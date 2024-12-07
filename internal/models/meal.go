package models

import (
	"time"

	"github.com/google/uuid"
)

type Meal struct {
	ID               uuid.UUID `json:"id"`
	EntityID         uuid.UUID `json:"entity_id"`
	MealTime         time.Time `json:"meal_time"`
	MealType         string    `json:"meal_type"`
	Foods            string    `json:"foods"`
	Beverages        string    `json:"beverages"`
	Medications      string    `json:"medications"`
	Symptoms         string    `json:"symptoms"`
	Mood             string    `json:"mood"`
	WeatherID        uuid.UUID `json:"weather_id"`
	Exercise         bool      `json:"exercise"`
	ExerciseDuration int       `json:"exercise_duration"`
	SleepDuration    int       `json:"sleep_duration"`
	MenstrualPhase   string    `json:"menstrual_phase"`
	Notes            string    `json:"notes"`
	CreatedAt        time.Time `json:"created_at"`
}
