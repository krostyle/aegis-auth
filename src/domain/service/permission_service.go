package service

import (
	"github.com/krostyle/auth-systme-go/src/domain/contract"
	"github.com/krostyle/auth-systme-go/src/domain/entity"
)

type PermissionService struct {
	Permission *entity.Permission
}

func NewPermissionService() contract.PermissionInterface {
	permissionService := &PermissionService{
		Permission: &entity.Permission{},
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
