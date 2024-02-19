package entity

import "time"

type Permission struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func NewPermission(id, name string) *Permission {
	return &Permission{
		ID:   id,
		Name: name,
	}
}
