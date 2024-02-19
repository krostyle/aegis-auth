package controller

import (
	"github.com/krostyle/auth-systme-go/src/application/usecase"
)

type PermissionController struct {
	permissionUseCase usecase.PermissionUseCase
}

func NewPermissionController(permissionUseCase usecase.PermissionUseCase) *PermissionController {
	return &PermissionController{
		permissionUseCase: permissionUseCase,
	}
}
