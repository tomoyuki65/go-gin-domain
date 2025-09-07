package user

import (
	"context"

	domain_user "go-gin-domain/internal/domain/user"
)

func (u *userUsecase) FindByUID(ctx context.Context, uid string) (*domain_user.User, error) {
	return u.userRepo.FindByUID(ctx, u.db, uid)
}
