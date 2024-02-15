package service

import (
	"fmt"

	"github.com/krostyle/auth-systme-go/src/domain/contract"
	"github.com/krostyle/auth-systme-go/src/domain/domainerror"
	"github.com/krostyle/auth-systme-go/src/domain/entity"
)

type RoleService struct {
	Role *entity.Role
}

func NewRoleService() contract.RoleInterface {
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

func (r *RoleService) GetPermissions() []contract.PermissionInterface {
	permissionsInterfaces := make([]contract.PermissionInterface, len(r.Role.Permissions))
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

func validatePermission(permission contract.PermissionInterface) error {
	if permission.GetID() == "" || permission.GetName() == "" {
		return domainerror.ErrInvalidPermission

	}
	return nil
}

func (r *RoleService) AddPermission(permission contract.PermissionInterface) error {
	if err := validatePermission(permission); err != nil {
		return fmt.Errorf("invalid permission: %w", err)
	}

	if r.HasPermission(permission) {
		return domainerror.ErrPermissionExists
	}

	r.Role.Permissions = append(r.Role.Permissions, *permission.(*PermissionService).Permission) //Repasar Sintaxis (Type Assertion y Dereferenciación)
	return nil
}

func (r *RoleService) RemovePermission(permission contract.PermissionInterface) error {
	for i, existingPermission := range r.Role.Permissions {
		if existingPermission.ID == permission.GetID() {
			r.Role.Permissions = append(r.Role.Permissions[:i], r.Role.Permissions[i+1:]...)
			return nil
		}
	}

	return domainerror.ErrPermissionNotFound
}

func (r *RoleService) HasPermission(permission contract.PermissionInterface) bool {
	for _, existingPerm := range r.Role.Permissions {
		if existingPerm.ID == permission.GetID() {
			return true
		}
	}
	return false
}
