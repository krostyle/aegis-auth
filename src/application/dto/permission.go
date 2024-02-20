package dto

import "time"

type PermissionCreateDTO struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type PermissionUpdateDTO struct {
	Name string `json:"name"`
}

type PermissionGetDTO struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type PermissionListDTO struct {
	Permissions []*PermissionGetDTO `json:"permissions"`
}