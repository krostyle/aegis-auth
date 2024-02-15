package entity

import "time"

type Role struct {
	ID          string       `json:"id"`
	Name        string       `json:"name"`
	Permissions []Permission `json:"permissions"`
	CreatedAt   time.Time    `json:"created_at"`
	UpdatedAt   time.Time    `json:"updated_at"`
}
