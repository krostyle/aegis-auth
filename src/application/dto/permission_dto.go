package dto

type PermissionCreateDTO struct {
	Name string `json:"name"`
}

type PermissionUpdateDTO struct {
	Name string `json:"name"`
}

type PermissionDTO struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}
