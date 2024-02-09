package domain_entities_implementations

import (
	domain_entities "github.com/krostyle/auth-systme-go/src/domain/entities"
	domain_entities_interfaces "github.com/krostyle/auth-systme-go/src/domain/interfaces"
)

type RoleManager struct {
	Role *domain_entities.Role
}

func NewRoleEntity() domain_entities_interfaces.RoleInterface {
	roleManager := &RoleManager{
		Role: &domain_entities.Role{},
	}
	return roleManager
}

func (r *RoleManager) GetID() string {
	return r.Role.ID
}

func (r *RoleManager) GetName() string {
	return r.Role.Name
}

func (r *RoleManager) GetPermissions() []domain_entities_interfaces.PermissionInterface {
	permissionsInterfaces := make([]domain_entities_interfaces.PermissionInterface, len(r.Role.Permissions))
	for i, perm := range r.Role.Permissions {
		permissionsInterfaces[i] = &PermissionManager{Permission: &perm}
	}
	return permissionsInterfaces
}

func (r *RoleManager) SetID(id string) {
	r.Role.ID = id
}

func (r *RoleManager) SetName(name string) {
	r.Role.Name = name
}

func (r *RoleManager) SetPermissions(permissions []domain_entities_interfaces.PermissionInterface) {
	permissionsEntities := make([]domain_entities.Permission, len(permissions))
	for i, perm := range permissions {
		permissionsEntities[i] = *perm.(*PermissionManager).Permission
	}
	r.Role.Permissions = permissionsEntities
}
