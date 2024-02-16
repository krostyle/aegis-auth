package mapper

import (
	"github.com/krostyle/auth-systme-go/src/domain/entity"
	"github.com/krostyle/auth-systme-go/src/infrastructure/database/postgres_client/model"
)

func UserToDomain(userModel *model.User) *entity.User {
	return &entity.User{
		ID:           userModel.ID,
		Name:         userModel.Name,
		Lastname:     userModel.Lastname,
		PasswordHash: userModel.PasswordHash,
		Email:        userModel.Email,
		CreatedAt:    userModel.CreatedAt,
		UpdatedAt:    userModel.UpdatedAt,
	}
}

func UserToORM(user *entity.User) *model.User {
	return &model.User{
		ID:           user.ID,
		Name:         user.Name,
		Lastname:     user.Lastname,
		PasswordHash: user.PasswordHash,
		Email:        user.Email,
		CreatedAt:    user.CreatedAt,
		UpdatedAt:    user.UpdatedAt,
	}
}

func UserToDomainList(userModels []*model.User) []*entity.User {
	var users []*entity.User
	for _, userModel := range userModels {
		users = append(users, UserToDomain(userModel))
	}
	return users
}

func UserToORMList(users []*entity.User) []*model.User {
	var userModels []*model.User
	for _, user := range users {
		userModels = append(userModels, UserToORM(user))
	}
	return userModels
}
