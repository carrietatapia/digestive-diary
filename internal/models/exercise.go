package models

import (
	"time"

	"github.com/google/uuid"
)

type Exercise struct {
	ID           uuid.UUID `json:"id"`
	EntityID     uuid.UUID `json:"entity_id"`
	ExerciseType string    `json:"exercise_type"`
	Duration     int       `json:"duration"`
	Intensity    string    `json:"intensity"`
	PerformedAt  time.Time `json:"performed_at"`
	CreatedAt    time.Time `json:"created_at"`
}
