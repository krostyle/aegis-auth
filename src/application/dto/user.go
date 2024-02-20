package dto

import "time"

type UserCreateDTO struct {
	Name     string `json:"name"`
	Lastname string `json:"lastname"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserUpdateDTO struct {
	Name     string `json:"name"`
	Lastname string `json:"lastname"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserGetDTO struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	Lastname  string    `json:"lastname"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type UserListDTO struct {
	Users []*UserGetDTO `json:"users"`
}

type UserLoginDTO struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserLoginResponseDTO struct {
	Token string `json:"token"`
}
