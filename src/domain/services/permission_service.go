package domain_entities_implementations

import (
	domain_entities "github.com/krostyle/auth-systme-go/src/domain/entities"
	domain_entities_interfaces "github.com/krostyle/auth-systme-go/src/domain/interfaces"
)

type PermissionService struct {
	Permission *domain_entities.Permission
}

func NewPermissionEntity() domain_entities_interfaces.PermissionInterface {
	permissionManager := &PermissionService{
		Permission: &domain_entities.Permission{},
	}
	return permissionManager
}

func (p *PermissionService) GetID() string {
	return p.Permission.ID
}

func (p *PermissionService) GetName() string {
	return p.Permission.Name
}

func (p *PermissionService) SetID(id string) {
	p.Permission.ID = id
}

func (p *PermissionService) SetName(name string) {
	p.Permission.Name = name
}
