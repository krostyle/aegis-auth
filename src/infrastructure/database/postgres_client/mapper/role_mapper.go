package mapper

import (
	"github.com/krostyle/auth-systme-go/src/domain/entity"
	"github.com/krostyle/auth-systme-go/src/infrastructure/database/postgres_client/model"
)

func RoleToDomain(roleModel *model.Role) *entity.Role {
	return &entity.Role{
		ID:        roleModel.ID,
		Name:      roleModel.Name,
		CreatedAt: roleModel.CreatedAt,
		UpdatedAt: roleModel.UpdatedAt,
	}
}

func RoleToORM(role *entity.Role) *model.Role {
	return &model.Role{
		ID:        role.ID,
		Name:      role.Name,
		CreatedAt: role.CreatedAt,
		UpdatedAt: role.UpdatedAt,
	}
}

func RoleToDomainList(roleModels []*model.Role) []*entity.Role {
	var roles []*entity.Role
	for _, roleModel := range roleModels {
		roles = append(roles, RoleToDomain(roleModel))
	}
	return roles
}

func RoleToORMList(roles []*entity.Role) []*model.Role {
	var roleModels []*model.Role
	for _, role := range roles {
		roleModels = append(roleModels, RoleToORM(role))
	}
	return roleModels
}
