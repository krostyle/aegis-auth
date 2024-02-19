package entity

import "time"

type Role struct {
	ID          string       `json:"id"`
	Name        string       `json:"name"`
	Permissions []Permission `json:"permissions"`
	CreatedAt   time.Time    `json:"created_at"`
	UpdatedAt   time.Time    `json:"updated_at"`
}

func NewRole(id, name string, permissions []Permission) *Role {
	return &Role{
		ID:   id,
		Name: name,
	}
}

func (r *Role) AddPermission(p *Permission) {
	r.Permissions = append(r.Permissions, *p)
}

func (r *Role) RemovePermission(p *Permission) {
	for i, perm := range r.Permissions {
		if perm.ID == p.ID {
			r.Permissions = append(r.Permissions[:i], r.Permissions[i+1:]...)
			break
		}
	}
}

func (r *Role) HasPermission(p *Permission) bool {
	for _, perm := range r.Permissions {
		if perm.ID == p.ID {
			return true
		}
	}
	return false
}
