package contract

type UserInterface interface {
	GetID() string
	GetName() string
	GetLastname() string
	GetEmail() string
	GetRoles() []RoleInterface

	SetID(string)
	SetName(string)
	SetLastname(string)
	SetEmail(string)

	AddRole(RoleInterface) error
	RemoveRole(RoleInterface) error
	HasRole(RoleInterface) bool
}
