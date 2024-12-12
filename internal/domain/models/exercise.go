package models

import (
	"time"
)

type Exercise struct {
	ID           string    `json:"id"`
	EntityID     string    `json:"entity_id"`
	ExerciseType string    `json:"exercise_type"`
	Duration     int       `json:"duration"`
	Intensity    string    `json:"intensity"`
	PerformedAt  time.Time `json:"performed_at"`
	CreatedAt    time.Time `json:"created_at"`
}
