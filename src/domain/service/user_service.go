package service

import (
	"github.com/krostyle/auth-systme-go/src/domain/domain_errors"
	"github.com/krostyle/auth-systme-go/src/domain/entity"
	"github.com/krostyle/auth-systme-go/src/domain/interfaces"
)

type UserService struct {
	User           *entity.User
	PasswordHasher interfaces.PasswordHasherInterface
}

func NewUserService() interfaces.UserInterface {
	userService := &UserService{
		User: &entity.User{},
	}
	return userService
}

func (u *UserService) GetID() string {
	return u.User.ID
}

func (u *UserService) GetName() string {
	return u.User.Name
}

func (u *UserService) GetLastname() string {
	return u.User.Lastname
}

func (u *UserService) GetEmail() string {
	return u.User.Email
}

func (u *UserService) GetRoles() []interfaces.RoleInterface {
	rolesInterfaces := make([]interfaces.RoleInterface, len(u.User.Roles))
	for i := range u.User.Roles {
		rolesInterfaces[i] = &RoleService{Role: &u.User.Roles[i]}
	}
	return rolesInterfaces
}

func (u *UserService) SetID(id string) {
	u.User.ID = id
}

func (u *UserService) SetName(name string) {
	u.User.Name = name
}

func (u *UserService) SetLastname(lastname string) {
	u.User.Lastname = lastname
}

func (u *UserService) SetEmail(email string) {
	u.User.Email = email
}

func (u *UserService) HasRole(role interfaces.RoleInterface) bool {
	for _, r := range u.User.Roles {
		if r.ID == role.GetID() {
			return true
		}
	}
	return false
}

func (u *UserService) AddRole(role interfaces.RoleInterface) error {
	if u.HasRole(role) {
		return domain_errors.ErrRoleExists
	}
	u.User.Roles = append(u.User.Roles, *role.(*RoleService).Role)
	return nil
}

func (u *UserService) RemoveRole(role interfaces.RoleInterface) error {
	for i, r := range u.User.Roles {
		if r.ID == role.GetID() {
			u.User.Roles = append(u.User.Roles[:i], u.User.Roles[i+1:]...)
			return nil
		}
	}
	return domain_errors.ErrRoleNotFound
}
