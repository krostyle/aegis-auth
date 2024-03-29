package domainerror

import "errors"

var (
	ErrInvalidPermission       = errors.New("invalid permission: ID and Name are required")
	ErrPermissionAlreadyExists = errors.New("permission already exists")
	ErrPermissionNotFound      = errors.New("permission not found in the role")
	ErrInvalidRole             = errors.New("invalid role: ID and Name are required")
	ErrRoleExists              = errors.New("role already exists in the user")
	ErrRoleNotFound            = errors.New("role not found in the user")
	ErrInvalidUser             = errors.New("invalid user: ID, Name, Lastname and Email are required")
	ErrUserAlreadyExists       = errors.New("user already exists")
	ErrUserNotFound            = errors.New("user not found")
	ErrInvalidPassword         = errors.New("invalid password")
	ErrInvalidEmail            = errors.New("invalid email")
)
