package domain_entities_implementations

import (
	domain_entities "github.com/krostyle/auth-systme-go/src/domain/entities"
	domain_entities_interfaces "github.com/krostyle/auth-systme-go/src/domain/interfaces"
)

type RoleService struct {
	Role *domain_entities.Role
}

func NewRoleEntity() domain_entities_interfaces.RoleInterface {
	roleManager := &RoleService{
		Role: &domain_entities.Role{},
	}
	return roleManager
}

func (r *RoleService) GetID() string {
	return r.Role.ID
}

func (r *RoleService) GetName() string {
	return r.Role.Name
}

func (r *RoleService) GetPermissions() []domain_entities_interfaces.PermissionInterface {
	permissionsInterfaces := make([]domain_entities_interfaces.PermissionInterface, len(r.Role.Permissions))
	for i := range r.Role.Permissions {
		permissionsInterfaces[i] = &PermissionService{Permission: &r.Role.Permissions[i]}
	}
	return permissionsInterfaces
}

func (r *RoleService) SetID(id string) {
	r.Role.ID = id
}

func (r *RoleService) SetName(name string) {
	r.Role.Name = name
}

func (r *RoleService) SetPermissions(permissions []domain_entities_interfaces.PermissionInterface) {
	permissionsEntities := make([]domain_entities.Permission, len(permissions))
	for i, perm := range permissions {
		permissionsEntities[i] = *perm.(*PermissionService).Permission
	}
	r.Role.Permissions = permissionsEntities
}
