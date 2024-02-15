package interfaces

import (
	"context"

	"github.com/krostyle/auth-systme-go/src/application/dto"
)

type RoleUseCaseInterface interface {
	CreateRole(ctx context.Context, role *dto.RoleCreateDTO) error
	// GetRoleByID(ctx context.Context, id string) (*dto.RoleGetDTO, error)
	// UpdateRole(ctx context.Context, id string, role *dto.RoleUpdateDTO) error
	// DeleteRole(ctx context.Context, id string) error
	GetAllRoles(ctx context.Context) (*dto.RoleListDTO, error)
}
