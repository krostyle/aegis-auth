package domain_services

import (
	domain_entities_interfaces "github.com/krostyle/auth-systme-go/src/domain/domain_interfaces"
	domain_entities "github.com/krostyle/auth-systme-go/src/domain/entities"
)

type PermissionService struct {
	Permission *domain_entities.Permission
}

func NewPermissionService() domain_entities_interfaces.PermissionInterface {
	permissionService := &PermissionService{
		Permission: &domain_entities.Permission{},
	}
	return permissionService
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
