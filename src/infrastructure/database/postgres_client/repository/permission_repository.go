package repository

import (
	"context"

	"github.com/krostyle/auth-systme-go/src/domain/entity"
	"github.com/krostyle/auth-systme-go/src/infrastructure/database/postgres_client/mapper"
	"github.com/krostyle/auth-systme-go/src/infrastructure/database/postgres_client/model"
	"gorm.io/gorm"
)

type PermissionRepository struct {
	db *gorm.DB
}

func NewPermissionRepository(db *gorm.DB) *PermissionRepository {
	return &PermissionRepository{db}
}

func (pr *PermissionRepository) CreatePermission(ctx context.Context, permission *entity.Permission) error {
	permissionModel := mapper.ToORM(permission)
	return pr.db.WithContext(ctx).Create(permissionModel).Error
}

func (pr *PermissionRepository) GetPermissionByID(ctx context.Context, id string) (*entity.Permission, error) {
	var permissionModel model.Permission
	err := pr.db.WithContext(ctx).Where("id = ?", id).First(&permissionModel).Error
	if err != nil {
		return nil, err
	}
	return mapper.ToDomain(&permissionModel), nil
}

func (pr *PermissionRepository) GetPermissionByName(ctx context.Context, name string) (*entity.Permission, error) {
	var permissionModel model.Permission
	err := pr.db.WithContext(ctx).Where("name = ?", name).First(&permissionModel).Error
	if err != nil {
		return nil, err
	}
	return mapper.ToDomain(&permissionModel), nil
}

func (pr *PermissionRepository) UpdatePermission(ctx context.Context, permission *entity.Permission) error {
	permissionModel := mapper.ToORM(permission)
	return pr.db.WithContext(ctx).Model(&model.Permission{}).Where("id = ?", permission.ID).Updates(permissionModel).Error
}

func (pr *PermissionRepository) DeletePermission(ctx context.Context, id string) error {
	return pr.db.WithContext(ctx).Where("id = ?", id).Delete(&model.Permission{}).Error
}
