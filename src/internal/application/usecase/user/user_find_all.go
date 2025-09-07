package user

import (
	"context"

	domain_user "go-gin-domain/internal/domain/user"
)

func (u *userUsecase) FindAll(ctx context.Context) ([]*domain_user.User, error) {
	return u.userRepo.FindAll(ctx, u.db)
}
