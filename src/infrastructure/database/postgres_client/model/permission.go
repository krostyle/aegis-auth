package model

import "gorm.io/gorm"

type Permission struct {
	gorm.Model
	ID   string `json:"id"`
	Name string `json:"name"`
}
