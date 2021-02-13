package models

import (
	"github.com/google/uuid"
)

type Bank struct {
	ID   uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4()"`
	Name string    `json:"name"`
	Code int       `json:"code"`
}
