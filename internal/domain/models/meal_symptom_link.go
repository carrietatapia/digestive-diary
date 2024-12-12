package models

import (
	"time"
)

type MealSymptomLink struct {
	ID        string    `json:"id"`
	MealID    string    `json:"meal_id"`
	Symptom   string    `json:"symptom"`
	CreatedAt time.Time `json:"created_at"`
}
