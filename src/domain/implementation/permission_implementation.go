package domain_entities_implementations

import (
	domain_entities "github.com/krostyle/auth-systme-go/src/domain/entities"
	domain_entities_interfaces "github.com/krostyle/auth-systme-go/src/domain/interfaces"
)

type PermissionManager struct {
	Permission *domain_entities.Permission
}

func NewPermissionEntity() domain_entities_interfaces.PermissionInterface {
	permissionManager := &PermissionManager{
		Permission: &domain_entities.Permission{},
	}
	return permissionManager
}

func (p *PermissionManager) GetID() string {
	return p.Permission.ID
}

func (p *PermissionManager) GetName() string {
	return p.Permission.Name
}

func (p *PermissionManager) SetID(id string) {
	p.Permission.ID = id
}

func (p *PermissionManager) SetName(name string) {
	p.Permission.Name = name
}
