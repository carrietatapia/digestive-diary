package models

import (
	"time"

	"github.com/google/uuid"
)

type MenstrualCycle struct {
	ID             uuid.UUID `json:"id"`
	EntityID       uuid.UUID `json:"entity_id"`
	CycleStartDate time.Time `json:"cycle_start_date"`
	Phase          string    `json:"phase"`
	Symptoms       string    `json:"symptoms"`
	CreatedAt      time.Time `json:"created_at"`
}
