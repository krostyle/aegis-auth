package contract

type RoleInterface interface {
	GetID() string
	GetName() string
	GetPermissions() []PermissionInterface

	SetID(string)
	SetName(string)

	AddPermission(PermissionInterface) error
	RemovePermission(PermissionInterface) error
	HasPermission(PermissionInterface) bool
}
