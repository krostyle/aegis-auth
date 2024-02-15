package repository

import (
	"context"

	"github.com/krostyle/auth-systme-go/src/domain/entity"
	"github.com/krostyle/auth-systme-go/src/infrastructure/database/postgres_client/mapper"
	"github.com/krostyle/auth-systme-go/src/infrastructure/database/postgres_client/model"
	"gorm.io/gorm"
)

type RoleRepository struct {
	db *gorm.DB
}

func NewRoleRepository(db *gorm.DB) *RoleRepository {
	return &RoleRepository{db}
}

func (rr *RoleRepository) CreateRole(ctx context.Context, role *entity.Role) error {
	roleModel := mapper.RoleToORM(role)
	return rr.db.WithContext(ctx).Create(roleModel).Error
}

func (rr *RoleRepository) GetAllRoles(ctx context.Context) ([]*entity.Role, error) {
	var roleModels []*model.Role
	err := rr.db.WithContext(ctx).Find(&roleModels).Error
	if err != nil {
		return nil, err
	}
	return mapper.RoleToDomainList(roleModels), nil
}
