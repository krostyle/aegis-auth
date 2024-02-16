package interfaces

import (
	"context"

	"github.com/krostyle/auth-systme-go/src/application/dto"
)

type PermissionCrudInterface interface {
	CreatePermission(ctx context.Context, permission *dto.PermissionCreateDTO) error
	GetPermissionByID(ctx context.Context, id string) (*dto.PermissionGetDTO, error)
	UpdatePermission(ctx context.Context, id string, permission *dto.PermissionUpdateDTO) error
	DeletePermission(ctx context.Context, id string) error
	GetAllPermissions(ctx context.Context) (*dto.PermissionListDTO, error)
}
