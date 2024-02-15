package repository

import (
	"context"

	"github.com/krostyle/auth-systme-go/src/domain/entity"
)

type RoleRepositoryInterface interface {
	CreateRole(ctx context.Context, role *entity.Role) error
	GetRoleByID(ctx context.Context, id string) (*entity.Role, error)
	GetAllRoles(ctx context.Context) ([]*entity.Role, error)
	UpdateRole(ctx context.Context, id string, role *entity.Role) error
	DeleteRole(ctx context.Context, id string) error
}
