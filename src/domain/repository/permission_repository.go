package repository

import (
	"context"

	"github.com/krostyle/auth-systme-go/src/domain/entity"
)

type PermissionRepositoryInterface interface {
	CreatePermission(ctx context.Context, permission *entity.Permission) error
	GetPermissionByID(ctx context.Context, id string) (*entity.Permission, error)
	UpdatePermission(ctx context.Context, id string, permission *entity.Permission) error
	DeletePermission(ctx context.Context, id string) error
	GetAllPermissions(ctx context.Context) ([]*entity.Permission, error)
	GetPermissionByName(ctx context.Context, name string) (*entity.Permission, error)
}
