package usecase

import (
	"context"
	"fmt"

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
	tokenGenerator      service.TokenGeneratorInterface
}

func NewUserUseCase(
	userRepository repository.UserRepositoryInterface,
	identifierGenerator service.IdentifierGeneratorInterface,
	passwordHasher service.PasswordHasherInterface,
	tokenGenerator service.TokenGeneratorInterface,
) interfaces.UserUseCaseInterface {
	return &UserUseCase{
		userRepository:      userRepository,
		identifierGenerator: identifierGenerator,
		passwordHasher:      passwordHasher,
		tokenGenerator:      tokenGenerator,
	}
}

func (u *UserUseCase) RegisterUser(ctx context.Context, user dto.UserCreateDTO) error {
	fmt.Println("Registering user...")
	userID := u.identifierGenerator.GenerateUUID()
	hashedPassword, err := u.passwordHasher.HashPassword(user.Password)
	if err != nil {
		fmt.Println("Error hashing password")
		return err
	}

	fmt.Println("Checking if user exists...")
	_, err = u.userRepository.GetUserByEmail(ctx, user.Email)

	if err == nil {
		fmt.Println("User already exists")
		return domainerror.ErrUserAlreadyExists
	}

	fmt.Println("No user with this email found. Creating user...")
	userEntity := entity.NewUser(userID, user.Name, user.Lastname, user.Email, hashedPassword)
	err = u.userRepository.CreateUser(ctx, userEntity)
	if err != nil {
		return err
	}
	fmt.Println("User created successfully")
	return nil
}

func (u *UserUseCase) GetAllUsers(ctx context.Context) (*dto.UserListDTO, error) {
	users, err := u.userRepository.GetAllUsers(ctx)
	if err != nil {
		return nil, err
	}
	userList := &dto.UserListDTO{}

	for _, user := range users {
		userDTO := dto.UserGetDTO{
			ID:        user.ID,
			Name:      user.Name,
			Lastname:  user.Lastname,
			Email:     user.Email,
			CreatedAt: user.CreatedAt,
			UpdatedAt: user.UpdatedAt,
		}
		userList.Users = append(userList.Users, &userDTO)
	}

	return userList, nil
}

func (u *UserUseCase) LoginUser(ctx context.Context, user dto.UserLoginDTO) (*dto.UserLoginResponseDTO, error) {
	userEntity, err := u.userRepository.GetUserByEmail(ctx, user.Email)
	if err != nil {
		return nil, domainerror.ErrUserNotFound
	}

	if !u.passwordHasher.VerifyPassword(userEntity.PasswordHash, user.Password) {
		return nil, domainerror.ErrInvalidPassword
	}

	token, err := u.tokenGenerator.GenerateToken(userEntity.ID)
	if err != nil {
		return nil, err
	}

	userLoginResponse := &dto.UserLoginResponseDTO{
		Token: token,
	}

	return userLoginResponse, nil
}
