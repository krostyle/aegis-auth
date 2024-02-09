package domain_entities_interfaces

type PermissionInterface interface {
	GetID() string
	GetName() string

	SetID(string)
	SetName(string)
}
