package entity

import (
	"time"
)

type User struct {
	ID           string    `json:"id"`
	Name         string    `json:"name"`
	Lastname     string    `json:"lastname"`
	Email        string    `json:"email"`
	PasswordHash string    `json:"-"`
	Roles        []Role    `json:"roles"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

func NewUser(id, name, lastname, email, password string) *User {
	return &User{
		ID:           id,
		Name:         name,
		Lastname:     lastname,
		Email:        email,
		PasswordHash: password,
	}
}

func (u *User) AddRole(r *Role) {
	u.Roles = append(u.Roles, *r)
}

func (u *User) RemoveRole(r *Role) {

	for i, role := range u.Roles {
		if role.ID == r.ID {
			u.Roles = append(u.Roles[:i], u.Roles[i+1:]...)
			break
		}
	}
}

func (u *User) HasRole(r *Role) bool {
	for _, role := range u.Roles {
		if role.ID == r.ID {
			return true
		}
	}
	return false
}
