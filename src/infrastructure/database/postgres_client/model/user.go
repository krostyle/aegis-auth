package model

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID           string    `gorm:"primaryKey" unique:"true" json:"id" validate:"required"`
	Name         string    `gorm:"not null" json:"name" validate:"required"`
	Lastname     string    `gorm:"not null" json:"lastname" validate:"required"`
	Email        string    `gorm:"unique;not null" json:"email" validate:"required,email"`
	PasswordHash string    `json:"-"`
	Roles        []*Role   `gorm:"many2many:user_roles;" json:"roles"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}
