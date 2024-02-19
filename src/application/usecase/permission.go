package usecase

import (
	"github.com/krostyle/auth-systme-go/src/domain/repository"
	"github.com/krostyle/auth-systme-go/src/domain/service"
)

type PermissionUseCase struct {
	userRepository      repository.PermissionRepositoryInterface
	identifierGenerator service.IdentifierGeneratorInterface
}

func NewPermissionUseCase(userRepository repository.PermissionRepositoryInterface, identifierGenerator service.IdentifierGeneratorInterface) *PermissionUseCase {
	return &PermissionUseCase{
		userRepository:      userRepository,
		identifierGenerator: identifierGenerator,
	}
}
