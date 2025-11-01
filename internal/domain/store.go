package domain

import (
	"github.com/google/uuid"
	"time"
)

// Store represents seller's store
type Store struct {
	ID        uint      `json:"id" gorm:"primaryKey;autoIncrement"`
	UserID    uuid.UUID `json:"user_id" gorm:"type:uuid;not null;index"`
	Name      string    `json:"name" gorm:"not null;"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
