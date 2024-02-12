package domain_services

import (
	domain_interfaces "github.com/krostyle/auth-systme-go/src/domain/domain_interfaces"
	domain_entities "github.com/krostyle/auth-systme-go/src/domain/entities"
	domain_errors "github.com/krostyle/auth-systme-go/src/domain/errors"
)

type UserService struct {
	User           *domain_entities.User
	PasswordHasher domain_interfaces.PasswordHasherInterface
}

func NewUserService() domain_interfaces.UserInterface {
	userService := &UserService{
		User: &domain_entities.User{},
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

func (u *UserService) GetRoles() []domain_interfaces.RoleInterface {
	rolesInterfaces := make([]domain_interfaces.RoleInterface, len(u.User.Roles))
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

func (u *UserService) HasRole(role domain_interfaces.RoleInterface) bool {
	for _, r := range u.User.Roles {
		if r.ID == role.GetID() {
			return true
		}
	}
	return false
}

func (u *UserService) AddRole(role domain_interfaces.RoleInterface) error {
	if u.HasRole(role) {
		return domain_errors.ErrRoleExists
	}
	u.User.Roles = append(u.User.Roles, *role.(*RoleService).Role)
	return nil
}

func (u *UserService) RemoveRole(role domain_interfaces.RoleInterface) error {
	for i, r := range u.User.Roles {
		if r.ID == role.GetID() {
			u.User.Roles = append(u.User.Roles[:i], u.User.Roles[i+1:]...)
			return nil
		}
	}
	return domain_errors.ErrRoleNotFound
}

// func (u *UserService) VerifyPassword(password string) bool {
// 	return u.User.VerifyPassword(password)
// }
