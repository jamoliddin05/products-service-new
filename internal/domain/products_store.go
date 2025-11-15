package domain

import (
	"github.com/google/uuid"
	"time"
)

type ProductsStore struct {
	ID        uint      `json:"-" gorm:"primaryKey;autoIncrement"`
	UserID    uuid.UUID `json:"user_id" gorm:"type:uuid;not null;index"`
	Name      string    `json:"name" gorm:"not null;"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
}
