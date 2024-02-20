package mapper

import (
	"github.com/krostyle/auth-systme-go/src/domain/entity"
	"github.com/krostyle/auth-systme-go/src/infrastructure/database/postgres_client/model"
)

func PermissionToDomain(permissionModel *model.Permission) *entity.Permission {
	return &entity.Permission{
		ID:          permissionModel.ID,
		Name:        permissionModel.Name,
		Description: permissionModel.Description,
		CreatedAt:   permissionModel.CreatedAt,
		UpdatedAt:   permissionModel.UpdatedAt,
	}
}

func PermissionToORM(permission *entity.Permission) *model.Permission {
	return &model.Permission{
		ID:          permission.ID,
		Name:        permission.Name,
		Description: permission.Description,
		CreatedAt:   permission.CreatedAt,
		UpdatedAt:   permission.UpdatedAt,
	}
}

func PermissionToDomainList(permissionModels []*model.Permission) []*entity.Permission {
	var permissions []*entity.Permission
	for _, permissionModel := range permissionModels {
		permissions = append(permissions, PermissionToDomain(permissionModel))
	}
	return permissions
}

func PermissionToORMList(permissions []*entity.Permission) []*model.Permission {
	var permissionModels []*model.Permission
	for _, permission := range permissions {
		permissionModels = append(permissionModels, PermissionToORM(permission))
	}
	return permissionModels
}
