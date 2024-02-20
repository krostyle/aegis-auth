package interfaces

import (
	"context"

	"github.com/krostyle/auth-systme-go/src/application/dto"
)

type PermissionUseCaseInterface interface {
	CreatePermission(ctx context.Context, permission dto.PermissionCreateDTO) error
	GetAllPermissions(ctx context.Context) (*dto.PermissionListDTO, error)
}
