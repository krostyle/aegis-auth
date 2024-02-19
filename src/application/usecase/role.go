package usecase

import (
	"github.com/krostyle/auth-systme-go/src/domain/repository"
	"github.com/krostyle/auth-systme-go/src/domain/service"
)

type RoleUseCase struct {
	roleRepository      repository.PermissionRepositoryInterface
	identifierGenerator service.IdentifierGeneratorInterface
}
