package models

import (
	"time"

	"github.com/google/uuid"
)

type Base struct {
	ID        uuid.UUID `json:"id" gorm:"type:uuid;default:uuid_generate_v4()"`
	CreatedAt time.Time `json:"created_at" gorm:"column:created_at;type:timestamptz();autoCreateTime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"column:updated_at;type:timestamptz(6);autoUpdateTime"`
}
