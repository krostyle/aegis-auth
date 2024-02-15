package model

import (
	"time"

	"gorm.io/gorm"
)

type Permission struct {
	gorm.Model
	ID        string         `json:"id"`
	Name      string         `json:"name"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
}
