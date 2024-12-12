package models

import (
	"time"
)

type Meal struct {
	ID               string    `json:"id"`
	EntityID         string    `json:"entity_id"`
	MealTime         time.Time `json:"meal_time"`
	MealType         string    `json:"meal_type"`
	Foods            string    `json:"foods"`
	Beverages        string    `json:"beverages"`
	Medications      string    `json:"medications"`
	Symptoms         string    `json:"symptoms"`
	Mood             string    `json:"mood"`
	WeatherID        string    `json:"weather_id"`
	Exercise         bool      `json:"exercise"`
	ExerciseDuration int       `json:"exercise_duration"`
	SleepDuration    int       `json:"sleep_duration"`
	MenstrualPhase   string    `json:"menstrual_phase"`
	Notes            string    `json:"notes"`
	CreatedAt        time.Time `json:"created_at"`
}
