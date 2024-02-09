package domain_entities_interfaces

type RoleInterface interface {
	GetID() string
	GetName() string
	GetPermissions() []PermissionInterface

	SetID(string)
	SetName(string)
	SetPermissions([]PermissionInterface)
}
