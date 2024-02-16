package crud

import (
	"context"

	"github.com/krostyle/auth-systme-go/src/application/dto"
	"github.com/krostyle/auth-systme-go/src/domain/contract/service"
	"github.com/krostyle/auth-systme-go/src/domain/entity"
	"github.com/krostyle/auth-systme-go/src/domain/repository"
)

type UserCrud struct {
	UserRepository      repository.UserRepositoryInterface
	identifierGenerator service.IdentifierGeneratorInterface
	passwordHasher      service.PasswordHasherInterface
}

func NewUserCrud(userRepository repository.UserRepositoryInterface, identifierGenerator service.IdentifierGeneratorInterface, passwordHasher service.PasswordHasherInterface) *UserCrud {
	return &UserCrud{
		UserRepository:      userRepository,
		identifierGenerator: identifierGenerator,
		passwordHasher:      passwordHasher,
	}
}

func (u *UserCrud) CreateUser(ctx context.Context, user *dto.UserCreateDTO) error {
	passwordHash, err := u.passwordHasher.HashPassword(user.Password)
	if err != nil {
		return err
	}
	userEntity := &entity.User{
		ID:           u.identifierGenerator.GenerateUUID(),
		Name:         user.Name,
		Lastname:     user.Lastname,
		Email:        user.Email,
		PasswordHash: passwordHash,
	}
	return u.UserRepository.CreateUser(ctx, userEntity)

}

func (u *UserCrud) GetAllUsers(ctx context.Context) (*dto.UserListDTO, error) {
	users, err := u.UserRepository.GetAllUsers(ctx)
	if err != nil {
		return nil, err
	}
	userList := &dto.UserListDTO{}
	for _, user := range users {
		userList.Users = append(userList.Users, &dto.UserGetDTO{
			ID:        user.ID,
			Name:      user.Name,
			Lastname:  user.Lastname,
			Email:     user.Email,
			CreatedAt: user.CreatedAt,
			UpdatedAt: user.UpdatedAt,
		})
	}
	return userList, nil
}

func (u *UserCrud) GetUserByID(ctx context.Context, id string) (*dto.UserGetDTO, error) {
	user, err := u.UserRepository.GetUserByID(ctx, id)
	if err != nil {
		return nil, err
	}
	return &dto.UserGetDTO{
		ID:        user.ID,
		Name:      user.Name,
		Email:     user.Email,
		Lastname:  user.Lastname,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}, nil
}

func (u *UserCrud) UpdateUser(ctx context.Context, id string, user *dto.UserUpdateDTO) error {
	userEntity := &entity.User{
		Name:         user.Name,
		Lastname:     user.Lastname,
		Email:        user.Email,
		PasswordHash: user.Password,
	}
	return u.UserRepository.UpdateUser(ctx, id, userEntity)
}

func (u *UserCrud) DeleteUser(ctx context.Context, id string) error {
	return u.UserRepository.DeleteUser(ctx, id)
}
