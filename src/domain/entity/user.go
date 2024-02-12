package entity

import "time"

type User struct {
	ID           string    `json:"id"`
	Name         string    `json:"name"`
	Lastname     string    `json:"lastname"`
	Email        string    `json:"email"`
	PasswordHash string    `json:"-"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
	Roles        []Role    `json:"roles"`
}
