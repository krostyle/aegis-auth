package crud

import (
	"context"

	"github.com/krostyle/auth-systme-go/src/application/dto"
	"github.com/krostyle/auth-systme-go/src/application/interfaces"
	"github.com/krostyle/auth-systme-go/src/domain/contract/service"
	"github.com/krostyle/auth-systme-go/src/domain/entity"
	"github.com/krostyle/auth-systme-go/src/domain/repository"
)

type PermissionCrud struct {
	permissionRepository repository.PermissionRepositoryInterface
	identifierGenerator  service.IdentifierGeneratorInterface
}

func NewPermissionCrud(permissionRepository repository.PermissionRepositoryInterface, identifierGenerator service.IdentifierGeneratorInterface) interfaces.PermissionCrudInterface {
	return &PermissionCrud{
		permissionRepository: permissionRepository,
		identifierGenerator:  identifierGenerator,
	}
}

func (p *PermissionCrud) CreatePermission(ctx context.Context, permission *dto.PermissionCreateDTO) error {
	permissionEntity := &entity.Permission{
		ID:   p.identifierGenerator.GenerateUUID(),
		Name: permission.Name,
	}
	return p.permissionRepository.CreatePermission(ctx, permissionEntity)
}

func (p *PermissionCrud) GetPermissionByID(ctx context.Context, id string) (*dto.PermissionGetDTO, error) {
	permission, err := p.permissionRepository.GetPermissionByID(ctx, id)
	if err != nil {
		return nil, err
	}
	return &dto.PermissionGetDTO{
		ID:        permission.ID,
		Name:      permission.Name,
		CreatedAt: permission.CreatedAt,
		UpdatedAt: permission.UpdatedAt,
	}, nil

}

func (p *PermissionCrud) UpdatePermission(ctx context.Context, id string, permission *dto.PermissionUpdateDTO) error {
	permissionEntity := &entity.Permission{
		Name: permission.Name,
	}
	return p.permissionRepository.UpdatePermission(ctx, id, permissionEntity)
}

func (p *PermissionCrud) DeletePermission(ctx context.Context, id string) error {
	return p.permissionRepository.DeletePermission(ctx, id)
}

func (p *PermissionCrud) GetAllPermissions(ctx context.Context) (*dto.PermissionListDTO, error) {
	permissions, err := p.permissionRepository.GetAllPermissions(ctx)
	if err != nil {
		return nil, err
	}
	var permissionListDTO dto.PermissionListDTO
	for _, permission := range permissions {
		permissionListDTO.Permissions = append(permissionListDTO.Permissions, &dto.PermissionGetDTO{
			ID:        permission.ID,
			Name:      permission.Name,
			CreatedAt: permission.CreatedAt,
			UpdatedAt: permission.UpdatedAt,
		})
	}
	return &permissionListDTO, nil
}
