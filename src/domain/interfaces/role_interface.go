package domain

type RoleInterface interface {
	GetID() string
	GetName() string
	GetPermissions() []PermissionInterface

	SetID(string)
	SetName(string)
	SetPermissions([]PermissionInterface)
}
