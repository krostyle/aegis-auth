package domain_services

import (
	"fmt"

	domain_entities_interfaces "github.com/krostyle/auth-systme-go/src/domain/domain_interfaces"
	domain_entities "github.com/krostyle/auth-systme-go/src/domain/entities"
	domain_errors "github.com/krostyle/auth-systme-go/src/domain/errors"
)

type RoleService struct {
	Role *domain_entities.Role
}

func NewRoleService() domain_entities_interfaces.RoleInterface {
	roleService := &RoleService{
		Role: &domain_entities.Role{},
	}
	return roleService
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

func validatePermission(permission domain_entities_interfaces.PermissionInterface) error {
	if permission.GetID() == "" || permission.GetName() == "" {
		return domain_errors.ErrInvalidPermission
	}
	return nil
}

func (r *RoleService) AddPermission(permission domain_entities_interfaces.PermissionInterface) error {
	if err := validatePermission(permission); err != nil {
		return fmt.Errorf("invalid permission: %w", err)
	}

	if r.HasPermission(permission) {
		return domain_errors.ErrPermissionExists
	}

	r.Role.Permissions = append(r.Role.Permissions, *permission.(*PermissionService).Permission) //Repasar Sintaxis (Type Assertion y Dereferenciación)
	return nil
}

func (r *RoleService) RemovePermission(permission domain_entities_interfaces.PermissionInterface) error {
	for i, existingPermission := range r.Role.Permissions {
		if existingPermission.ID == permission.GetID() {
			r.Role.Permissions = append(r.Role.Permissions[:i], r.Role.Permissions[i+1:]...)
			return nil
		}
	}
	return domain_errors.ErrPermissionNotFound
}

func (r *RoleService) HasPermission(permission domain_entities_interfaces.PermissionInterface) bool {
	for _, existingPerm := range r.Role.Permissions {
		if existingPerm.ID == permission.GetID() {
			return true
		}
	}
	return false
}
