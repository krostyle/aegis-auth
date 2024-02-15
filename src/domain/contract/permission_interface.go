package contract

type PermissionInterface interface {
	GetID() string
	GetName() string

	SetID(string)
	SetName(string)
}
