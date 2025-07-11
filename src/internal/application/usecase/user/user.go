package user

import (
	"context"

	"go-gin-domain/internal/application/usecase/logger"
	domain_user "go-gin-domain/internal/domain/user"
)

type UserUsecase interface {
	Create(ctx context.Context, lastName, firstName, email string) (*domain_user.User, error)
	FindAll(ctx context.Context) ([]*domain_user.User, error)
	FindByUID(ctx context.Context, uid string) (*domain_user.User, error)
	Update(ctx context.Context, uid, lastName, firstName, email string) (*domain_user.User, error)
	Delete(ctx context.Context, uid string) (*domain_user.User, error)
}

type userUsecase struct {
	userRepo domain_user.UserRepository
	logger   logger.Logger
}

func NewUserUsecase(userRepo domain_user.UserRepository, logger logger.Logger) UserUsecase {
	return &userUsecase{
		userRepo: userRepo,
		logger:   logger,
	}
}
