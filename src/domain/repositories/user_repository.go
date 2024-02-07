package domain

import (
	"context"

	domain "github.com/krostyle/auth-systme-go/domain/entities"
)

// User represents the user entity's repository interface.
type UserRepositoryInterface interface {
	CreateUser(ctx context.Context, user *domain.User) error
	GetUserByID(ctx context.Context, id string) (*domain.User, error)
	GetUserByEmail(ctx context.Context, email string) (*domain.User, error)
	UpdateUser(ctx context.Context, user *domain.User) error
	DeleteUser(ctx context.Context, id string) error
}
