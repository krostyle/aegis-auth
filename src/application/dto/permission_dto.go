package dto

import "time"

type PermissionCreateDTO struct {
	Name string `json:"name"`
}

type PermissionUpdateDTO struct {
	Name string `json:"name"`
}

type PermissionDTO struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
