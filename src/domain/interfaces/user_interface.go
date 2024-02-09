package domain_entities_interfaces

type UserInterface interface {
	GetID() string
	GetName() string
	GetLastName() string
	GetEmail() string
	GetPassword() string
	GetRoles() []RoleInterface

	SetID(string)
	SetName(string)
	SetLastName(string)
	SetEmail(string)
	SetPassword(string)
	SetRoles([]RoleInterface)
}
