package interfaces

import (
	"context"

	"github.com/krostyle/auth-systme-go/src/application/dto"
)

type UserCrudInterface interface {
	CreateUser(ctx context.Context, user *dto.UserCreateDTO) error
	GetAllUsers(ctx context.Context) (*dto.UserListDTO, error)
	GetUserByID(ctx context.Context, id string) (*dto.UserGetDTO, error)
	UpdateUser(ctx context.Context, id string, user *dto.UserUpdateDTO) error
	DeleteUser(ctx context.Context, id string) error
}
