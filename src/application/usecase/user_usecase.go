package usecase

import (
	"context"

	"github.com/krostyle/auth-systme-go/src/domain/entity"
	"github.com/krostyle/auth-systme-go/src/domain/repository"
)

type UserUseCase struct {
	UserRepository repository.UserRepositoryInterface
}

func NewUserUseCase(userRepository repository.UserRepositoryInterface) *UserUseCase {
	return &UserUseCase{
		UserRepository: userRepository,
	}
}

func (u *UserUseCase) CreateUser(ctx context.Context, user *entity.User) error {
	return u.UserRepository.CreateUser(ctx, user)
}
