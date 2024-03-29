package model

import (
	"time"

	"gorm.io/gorm"
)

type Permission struct {
	gorm.Model
	ID          string         `gorm:"primaryKey" unique:"true" json:"id" validate:"required"`
	Name        string         `gorm:"unique;not null" json:"name" validate:"required"`
	Description string         `json:"description" validate:"required"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `json:"deleted_at"`
	Roles       []*Role        `gorm:"many2many:role_permissions;" json:"roles"`
}
