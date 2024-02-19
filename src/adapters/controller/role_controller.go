package controller

import (
	"github.com/krostyle/auth-systme-go/src/application/usecase"
)

type RoleController struct {
	roleUseCase usecase.RoleUseCase
}

func NewRoleController(roleUseCase usecase.RoleUseCase) *RoleController {
	return &RoleController{
		roleUseCase: roleUseCase,
	}
}
