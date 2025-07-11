package user

import (
	"context"
)

type UserRepository interface {
	Create(ctx context.Context, user *User) (*User, error)
	FindAll(ctx context.Context) ([]*User, error)
	FindByUID(ctx context.Context, uid string) (*User, error)
	Save(ctx context.Context, user *User) (*User, error)
}
