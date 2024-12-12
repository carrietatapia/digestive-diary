package models

import (
	"time"
)

type StoolConsistency struct {
	ID          string    `json:"id"`
	TypeNumber  int       `json:"type_number"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
}
