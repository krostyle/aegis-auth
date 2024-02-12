package service

import (
	"fmt"

	"github.com/krostyle/auth-systme-go/src/domain/domain_errors"
	"github.com/krostyle/auth-systme-go/src/domain/entity"
	"github.com/krostyle/auth-systme-go/src/domain/interfaces"
)

type RoleService struct {
	Role *entity.Role
}

func NewRoleService() interfaces.RoleInterface {
	roleService := &RoleService{
		Role: &entity.Role{},
	}
	return roleService
}

func (r *RoleService) GetID() string {
	return r.Role.ID
}

func (r *RoleService) GetName() string {
	return r.Role.Name
}

func (r *RoleService) GetPermissions() []interfaces.PermissionInterface {
	permissionsInterfaces := make([]interfaces.PermissionInterface, len(r.Role.Permissions))
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

func validatePermission(permission interfaces.PermissionInterface) error {
	if permission.GetID() == "" || permission.GetName() == "" {
		return domain_errors.ErrInvalidPermission

	}
	return nil
}

func (r *RoleService) AddPermission(permission interfaces.PermissionInterface) error {
	if err := validatePermission(permission); err != nil {
		return fmt.Errorf("invalid permission: %w", err)
	}

	if r.HasPermission(permission) {
		return domain_errors.ErrPermissionExists
	}

	r.Role.Permissions = append(r.Role.Permissions, *permission.(*PermissionService).Permission) //Repasar Sintaxis (Type Assertion y Dereferenciación)
	return nil
}

func (r *RoleService) RemovePermission(permission interfaces.PermissionInterface) error {
	for i, existingPermission := range r.Role.Permissions {
		if existingPermission.ID == permission.GetID() {
			r.Role.Permissions = append(r.Role.Permissions[:i], r.Role.Permissions[i+1:]...)
			return nil
		}
	}

	return domain_errors.ErrPermissionNotFound
}

func (r *RoleService) HasPermission(permission interfaces.PermissionInterface) bool {
	for _, existingPerm := range r.Role.Permissions {
		if existingPerm.ID == permission.GetID() {
			return true
		}
	}
	return false
}
