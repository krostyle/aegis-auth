package usecase

import (
	"context"

	"github.com/krostyle/auth-systme-go/src/application/dto"
	"github.com/krostyle/auth-systme-go/src/application/interfaces"
	"github.com/krostyle/auth-systme-go/src/domain/domainerror"
	"github.com/krostyle/auth-systme-go/src/domain/entity"
	"github.com/krostyle/auth-systme-go/src/domain/repository"
	"github.com/krostyle/auth-systme-go/src/domain/service"
)

type UserUseCase struct {
	userRepository      repository.UserRepositoryInterface
	identifierGenerator service.IdentifierGeneratorInterface
	passwordHasher      service.PasswordHasherInterface
}

func NewUserUseCase(
	userRepository repository.UserRepositoryInterface,
	identifierGenerator service.IdentifierGeneratorInterface,
	passwordHasher service.PasswordHasherInterface) interfaces.UserUseCaseInterface {
	return &UserUseCase{
		userRepository:      userRepository,
		identifierGenerator: identifierGenerator,
		passwordHasher:      passwordHasher,
	}
}

func (u *UserUseCase) RegisterUser(ctx context.Context, user dto.UserCreateDTO) error {
	userID := u.identifierGenerator.GenerateUUID()
	hashedPassword, err := u.passwordHasher.HashPassword(user.Password)
	if err != nil {
		return err
	}

	u.userRepository.GetUserByEmail(ctx, user.Email)
	if err != nil {
		return domainerror.ErrUserExists
	}

	userEntity := entity.NewUser(userID, user.Name, user.Lastname, user.Email, hashedPassword)
	err = u.userRepository.CreateUser(ctx, userEntity)
	if err != nil {
		return err
	}
	return nil
}
