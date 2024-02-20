package entity

import "time"

type Permission struct {
	ID          string    `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func NewPermission(id, name, description string) *Permission {
	return &Permission{
		ID:          id,
		Name:        name,
		Description: description,
	}
}
