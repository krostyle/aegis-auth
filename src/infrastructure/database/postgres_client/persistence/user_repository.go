package persistence

import (
	"context"

	"github.com/krostyle/auth-systme-go/src/domain/entity"
	"github.com/krostyle/auth-systme-go/src/infrastructure/database/postgres_client/mapper"
	"github.com/krostyle/auth-systme-go/src/infrastructure/database/postgres_client/model"
	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db}
}

func (ur *UserRepository) CreateUser(ctx context.Context, user *entity.User) error {
	userModel := mapper.UserToORM(user)
	return ur.db.WithContext(ctx).Create(userModel).Error
}

func (ur *UserRepository) GetAllUsers(ctx context.Context) ([]*entity.User, error) {
	var userModels []*model.User
	err := ur.db.WithContext(ctx).Find(&userModels).Error
	if err != nil {
		return nil, err
	}
	return mapper.UserToDomainList(userModels), nil
}

func (ur *UserRepository) GetUserByID(ctx context.Context, id string) (*entity.User, error) {
	var userModel model.User
	err := ur.db.WithContext(ctx).Where("id = ?", id).First(&userModel).Error
	if err != nil {
		return nil, err
	}
	return mapper.UserToDomain(&userModel), nil
}

func (ur *UserRepository) UpdateUser(ctx context.Context, id string, user *entity.User) error {
	userModel := mapper.UserToORM(user)
	return ur.db.WithContext(ctx).Model(&model.User{}).Where("id = ?", id).Updates(userModel).Error
}

func (ur *UserRepository) DeleteUser(ctx context.Context, id string) error {
	return ur.db.WithContext(ctx).Where("id = ?", id).Delete(&model.User{}).Error
}
