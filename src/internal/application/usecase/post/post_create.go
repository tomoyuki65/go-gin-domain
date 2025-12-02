package post

import (
	"context"
	"fmt"

	domain_post "go-gin-domain/internal/domain/post"
)

func (u *postUsecase) Create(ctx context.Context, text string) (*domain_post.Post, error) {
	// Postエンティティを新規作成
	post, err := domain_post.NewPost(text)
	if err != nil {
		err := fmt.Errorf("バリデーションエラー: %w", err)
		u.logger.Warn(ctx, err.Error())
		return nil, err
	}

	return u.postRepo.Create(ctx, u.db, post)
}
