package models

import (
	"time"

	"github.com/google/uuid"
)

type MealSymptomLink struct {
	ID        uuid.UUID `json:"id"`
	MealID    uuid.UUID `json:"meal_id"`
	Symptom   string    `json:"symptom"`
	CreatedAt time.Time `json:"created_at"`
}
