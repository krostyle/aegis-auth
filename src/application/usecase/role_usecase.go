package usecase

import (
	"context"

	"github.com/krostyle/auth-systme-go/src/application/dto"
	"github.com/krostyle/auth-systme-go/src/domain/contract/service"
	"github.com/krostyle/auth-systme-go/src/domain/entity"
	"github.com/krostyle/auth-systme-go/src/domain/repository"
)

type RoleUseCase struct {
	roleRepository      repository.RoleRepositoryInterface
	identifierGenerator service.IdentifierGeneratorInterface
}

func NewRoleUseCase(roleRepository repository.RoleRepositoryInterface, identifierGenerator service.IdentifierGeneratorInterface) *RoleUseCase {
	return &RoleUseCase{
		roleRepository:      roleRepository,
		identifierGenerator: identifierGenerator,
	}
}

func (r *RoleUseCase) CreateRole(ctx context.Context, role *dto.RoleCreateDTO) error {
	roleEntity := &entity.Role{
		ID:   r.identifierGenerator.GenerateUUID(),
		Name: role.Name,
	}
	return r.roleRepository.CreateRole(ctx, roleEntity)
}

func (r *RoleUseCase) GetAllRoles(ctx context.Context) (*dto.RoleListDTO, error) {
	roles, err := r.roleRepository.GetAllRoles(ctx)
	if err != nil {
		return nil, err
	}
	roleList := &dto.RoleListDTO{}
	for _, role := range roles {
		roleList.Roles = append(roleList.Roles, &dto.RoleGetDTO{
			ID:        role.ID,
			Name:      role.Name,
			CreatedAt: role.CreatedAt,
			UpdatedAt: role.UpdatedAt,
		})
	}
	return roleList, nil
}
