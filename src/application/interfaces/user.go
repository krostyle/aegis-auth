package interfaces

import (
	"context"

	"github.com/krostyle/auth-systme-go/src/application/dto"
)

type UserUseCaseInterface interface {
	RegisterUser(ctx context.Context, user dto.UserCreateDTO) error
	GetAllUsers(ctx context.Context) (*dto.UserListDTO, error)
	LoginUser(ctx context.Context, user dto.UserLoginDTO) (*dto.UserLoginResponseDTO, error)
}
