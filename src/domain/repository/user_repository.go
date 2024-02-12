package repository

import (
	"context"

	"github.com/krostyle/auth-systme-go/src/domain/entity"
)

type UserRepositoryInterface interface {
	CreateUser(ctx context.Context, user *entity.User) error
	GetUserByID(ctx context.Context, id string) (*entity.User, error)
	GetUserByEmail(ctx context.Context, email string) (*entity.User, error)
	UpdateUser(ctx context.Context, user *entity.User) error
	DeleteUser(ctx context.Context, id string) error
}
