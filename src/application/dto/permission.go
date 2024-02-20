package dto

import "time"

type PermissionCreateDTO struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

type PermissionUpdateDTO struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

type PermissionGetDTO struct {
	ID          string    `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type PermissionListDTO struct {
	Permissions []*PermissionGetDTO `json:"permissions"`
}
