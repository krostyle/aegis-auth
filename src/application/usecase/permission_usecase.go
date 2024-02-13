package usecase

import (
	"context"

	"github.com/krostyle/auth-systme-go/src/domain/entity"
	"github.com/krostyle/auth-systme-go/src/domain/repository"
)

type PermissionUseCase struct {
	permissionRepository repository.PermissionRepositoryInterface
}

func NewPermissionUseCase(permissionRepository repository.PermissionRepositoryInterface) *PermissionUseCase {
	return &PermissionUseCase{
		permissionRepository: permissionRepository,
	}
}

func (p *PermissionUseCase) CreatePermission(ctx context.Context, permission *entity.Permission) error {
	return p.permissionRepository.CreatePermission(ctx, permission)
}

func (p *PermissionUseCase) GetPermissionByID(ctx context.Context, id string) (*entity.Permission, error) {
	return p.permissionRepository.GetPermissionByID(ctx, id)
}

func (p *PermissionUseCase) GetPermissionByName(ctx context.Context, name string) (*entity.Permission, error) {
	return p.permissionRepository.GetPermissionByName(ctx, name)
}

func (p *PermissionUseCase) UpdatePermission(ctx context.Context, permission *entity.Permission) error {
	return p.permissionRepository.UpdatePermission(ctx, permission)
}

func (p *PermissionUseCase) DeletePermission(ctx context.Context, id string) error {
	return p.permissionRepository.DeletePermission(ctx, id)
}
