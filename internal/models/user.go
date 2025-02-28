package models

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID        uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	Firstname string    `gorm:"not null" validate:"required,min=2,max=50"`
	Lastname  string    `gorm:"not null" validate:"required,min=2,max=50"`
	Email     string    `gorm:"uniqueIndex;not null" validate:"required,email"`
	Password  string    `gorm:"not null" validate:"required,min=8,max=32"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
