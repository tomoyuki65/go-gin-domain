package user

import (
	"context"

	domain_user "go-gin-domain/internal/domain/user"

	"github.com/google/uuid"
)

func (u *userUsecase) Create(ctx context.Context, lastName, firstName, email string) (*domain_user.User, error) {
	// UIDの設定（仮）
	uid := uuid.New().String()

	// 新規ユーザー作成
	user := domain_user.NewUser(uid, lastName, firstName, email)

	return u.userRepo.Create(ctx, u.db, user)
}
