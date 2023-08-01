package user

import (
	"context"
	"devops-platform/internal/deploy-system/user/internal/domain"
	"devops-platform/pkg/types"
)

const (
	BeanRepository = domain.BeanRepository
	BeanService    = domain.BeanService
)

type CreateUserCommand = domain.CreateUserCommand

type User = domain.User

type Service interface {
	Create(ctx context.Context, command *CreateUserCommand) (id types.Long, err error)
	GetByID(ctx context.Context, ID types.Long) (user *domain.User, err error)
}

type Repository interface {
	GetByUsername(ctx context.Context, username string) (user *domain.User, err error)
}

type Query interface {
	GetByUsername(ctx context.Context, username string) (user *domain.User, err error)
	GetByID(ctx context.Context, ID types.Long) (user *domain.User, err error)
}
