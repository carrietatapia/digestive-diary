package models

import (
	"time"

	"github.com/google/uuid"
)

type StoolConsistency struct {
	ID          uuid.UUID `json:"id"`
	TypeNumber  int       `json:"type_number"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
}
