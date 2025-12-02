package post

import (
	"context"
)

type PostRepository interface {
	// dbはトランザクションを使うことを考慮し、パラメータとして渡せるようにする。
	// 今回はdbはダミー設定を使うため、型はstringとしている。
	Create(ctx context.Context, db string, post *Post) (*Post, error)
	FindAll(ctx context.Context, db string) ([]*Post, error)
}
