package post

import (
	"context"

	domain_post "go-gin-domain/internal/domain/post"
)

func (u *postUsecase) FindAll(ctx context.Context) ([]*domain_post.Post, error) {
	return u.postRepo.FindAll(ctx, u.db)
}
