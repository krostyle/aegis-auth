package crud

import (
	"context"

	"github.com/krostyle/auth-systme-go/src/application/dto"
	"github.com/krostyle/auth-systme-go/src/domain/contract/service"
	"github.com/krostyle/auth-systme-go/src/domain/entity"
	"github.com/krostyle/auth-systme-go/src/domain/repository"
)

type RoleCrud struct {
	roleRepository      repository.RoleRepositoryInterface
	identifierGenerator service.IdentifierGeneratorInterface
}

func NewRoleCrud(roleRepository repository.RoleRepositoryInterface, identifierGenerator service.IdentifierGeneratorInterface) *RoleCrud {
	return &RoleCrud{
		roleRepository:      roleRepository,
		identifierGenerator: identifierGenerator,
	}
}

func (r *RoleCrud) CreateRole(ctx context.Context, role *dto.RoleCreateDTO) error {
	roleEntity := &entity.Role{
		ID:   r.identifierGenerator.GenerateUUID(),
		Name: role.Name,
	}
	return r.roleRepository.CreateRole(ctx, roleEntity)
}

func (r *RoleCrud) GetAllRoles(ctx context.Context) (*dto.RoleListDTO, error) {
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

func (r *RoleCrud) GetRoleByID(ctx context.Context, id string) (*dto.RoleGetDTO, error) {
	role, err := r.roleRepository.GetRoleByID(ctx, id)
	if err != nil {
		return nil, err
	}
	return &dto.RoleGetDTO{
		ID:        role.ID,
		Name:      role.Name,
		CreatedAt: role.CreatedAt,
		UpdatedAt: role.UpdatedAt,
	}, nil
}

func (r *RoleCrud) UpdateRole(ctx context.Context, id string, role *dto.RoleUpdateDTO) error {
	roleEntity := &entity.Role{
		Name: role.Name,
	}
	return r.roleRepository.UpdateRole(ctx, id, roleEntity)
}

func (r *RoleCrud) DeleteRole(ctx context.Context, id string) error {
	return r.roleRepository.DeleteRole(ctx, id)
}
