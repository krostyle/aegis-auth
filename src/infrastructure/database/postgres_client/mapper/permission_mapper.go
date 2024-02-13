package mapper

import (
	"github.com/krostyle/auth-systme-go/src/domain/entity"
	"github.com/krostyle/auth-systme-go/src/infrastructure/database/postgres_client/model"
)

func ToDomain(permissionModel *model.Permission) *entity.Permission {
	return &entity.Permission{
		ID:   permissionModel.ID,
		Name: permissionModel.Name,
	}
}

func ToORM(permission *entity.Permission) *model.Permission {
	return &model.Permission{
		ID:   permission.ID,
		Name: permission.Name,
	}
}
