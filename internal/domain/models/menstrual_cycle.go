package models

import (
	"time"
)

type MenstrualCycle struct {
	ID             string    `json:"id"`
	EntityID       string    `json:"entity_id"`
	CycleStartDate time.Time `json:"cycle_start_date"`
	Phase          string    `json:"phase"`
	Symptoms       string    `json:"symptoms"`
	CreatedAt      time.Time `json:"created_at"`
}
