package dto

import "time"

type RoleCreateDTO struct {
	Name string `json:"name"`
}

type RoleUpdateDTO struct {
	Name string `json:"name"`
}

type RoleGetDTO struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type RoleListDTO struct {
	Roles []*RoleGetDTO `json:"roles"`
}
