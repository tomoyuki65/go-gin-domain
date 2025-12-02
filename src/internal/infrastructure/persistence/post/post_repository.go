package post

import (
	"context"

	logger_usecase "go-gin-domain/internal/application/usecase/logger"
	domain "go-gin-domain/internal/domain/post"
)

type postRepository struct {
	logger logger_usecase.Logger
}

func NewPostRepository(logger logger_usecase.Logger) domain.PostRepository {
	return &postRepository{
		logger: logger,
	}
}

func (r *postRepository) Create(ctx context.Context, db string, post *domain.Post) (*domain.Post, error) {
	// DBへの登録処理をした後に戻り値を返す想定
	return post, nil
}

func (r *postRepository) FindAll(ctx context.Context, db string) ([]*domain.Post, error) {
	// DBからPostデータを取得したことを想定として固定値を定義
	dbPosts := []struct {
		Text string
	}{
		{
			Text: "テキスト１",
		},
		{
			Text: "テキスト２",
		},
	}

	// Postエンティティを利用してスライスを定義
	posts := make([]*domain.Post, 0, len(dbPosts))

	// ループ処理でPostエンティティのスライスへ変換
	for _, dbPost := range dbPosts {
		// 値のチェックは不要とし、DBから復元するためのコンストラクタを利用
		posts = append(posts, domain.ReconstitutePost(dbPost.Text))
	}

	return posts, nil
}
