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
}
