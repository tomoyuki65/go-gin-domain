package user

import (
	"context"
	"fmt"

	domain_user "go-gin-domain/internal/domain/user"
)

func (u *userUsecase) Delete(ctx context.Context, uid string) (*domain_user.User, error) {
	user, err := u.userRepo.FindByUID(ctx, u.db, uid)
	if err != nil {
		return nil, err
	}

	// 対象ユーザーが存在しない場合はエラー
	if user == nil {
		msg := fmt.Sprintf("対象ユーザーが存在しません。: UID=%s", uid)
		u.logger.Error(ctx, msg)
		return nil, fmt.Errorf("%s", msg)
	}

	// 論理削除設定
	user.SetDelete()

	return u.userRepo.Save(ctx, u.db, user)
}
