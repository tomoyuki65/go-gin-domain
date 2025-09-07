package user

import (
	"context"
)

type UserRepository interface {
	// dbはトランザクションを使うことを考慮し、パラメータとして渡せるようにする。
	// 今回はdbはダミー設定を使うため、型はstringとしている。
	Create(ctx context.Context, db string, user *User) (*User, error)
	FindAll(ctx context.Context, db string) ([]*User, error)
	FindByUID(ctx context.Context, db string, uid string) (*User, error)
	Save(ctx context.Context, db string, user *User) (*User, error)
}
