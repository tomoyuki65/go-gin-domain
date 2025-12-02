package post

import (
	"context"

	"go-gin-domain/internal/application/usecase/logger"
	domain_post "go-gin-domain/internal/domain/post"
)

type PostUsecase interface {
	Create(ctx context.Context, text string) (*domain_post.Post, error)
	FindAll(ctx context.Context) ([]*domain_post.Post, error)
}

type postUsecase struct {
	db       string
	postRepo domain_post.PostRepository
	logger   logger.Logger
}

func NewPostUsecase(db string, postRepo domain_post.PostRepository, logger logger.Logger) PostUsecase {
	return &postUsecase{
		db:       db,
		postRepo: postRepo,
		logger:   logger,
	}
}
