package usecase

import (
	"context"

	"github.com/krostyle/auth-systme-go/src/application/dto"
	"github.com/krostyle/auth-systme-go/src/application/interfaces"
	"github.com/krostyle/auth-systme-go/src/domain/domainerror"
	"github.com/krostyle/auth-systme-go/src/domain/entity"
	"github.com/krostyle/auth-systme-go/src/domain/repository"
	"github.com/krostyle/auth-systme-go/src/domain/service"
)

type PermissionUseCase struct {
	userRepository      repository.PermissionRepositoryInterface
	identifierGenerator service.IdentifierGeneratorInterface
}

func NewPermissionUseCase(
	userRepository repository.PermissionRepositoryInterface,
	identifierGenerator service.IdentifierGeneratorInterface,
) interfaces.PermissionUseCaseInterface {
	return &PermissionUseCase{
		userRepository:      userRepository,
		identifierGenerator: identifierGenerator,
	}
}

func (u *PermissionUseCase) CreatePermission(ctx context.Context, permission dto.PermissionCreateDTO) error {
	permissionID := u.identifierGenerator.GenerateUUID()
	// Check if permission already exists
	_, err := u.userRepository.GetPermissionByName(ctx, permission.Name)
	if err == nil {
		return domainerror.ErrPermissionAlreadyExists
	}
	permissionEntity := entity.NewPermission(permissionID, permission.Name, permission.Description)
	err = u.userRepository.CreatePermission(ctx, permissionEntity)
	if err != nil {
		return err
	}
	return nil

}

func (u *PermissionUseCase) GetAllPermissions(ctx context.Context) (*dto.PermissionListDTO, error) {
	permissions, err := u.userRepository.GetAllPermissions(ctx)
	if err != nil {
		return nil, err
	}

	permissionsList := &dto.PermissionListDTO{}
	for _, permission := range permissions {
		permissionDTO := dto.PermissionGetDTO{
			ID:          permission.ID,
			Name:        permission.Name,
			Description: permission.Description,
			CreatedAt:   permission.CreatedAt,
			UpdatedAt:   permission.UpdatedAt,
		}
		permissionsList.Permissions = append(permissionsList.Permissions, &permissionDTO)
	}
	return permissionsList, nil
}
