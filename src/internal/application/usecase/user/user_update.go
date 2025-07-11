package user

import (
	"context"
	"fmt"

	domain_user "go-gin-domain/internal/domain/user"
)

func (u *userUsecase) Update(ctx context.Context, uid, lastName, firstName, email string) (*domain_user.User, error) {
	user, err := u.userRepo.FindByUID(ctx, uid)
	if err != nil {
		return nil, err
	}

	// 対象ユーザーが存在しない場合はエラー
	if user == nil {
		msg := fmt.Sprintf("対象ユーザーが存在しません。: UID=%s", uid)
		u.logger.Error(ctx, msg)
		return nil, fmt.Errorf("%s", msg)
	}

	// プロフィール更新
	err = user.UpdateProfile(lastName, firstName, email)
	if err != nil {
		return nil, err
	}

	return u.userRepo.Save(ctx, user)
}
