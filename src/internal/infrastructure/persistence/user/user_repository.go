package user

import (
	"context"
	"fmt"
	"time"

	logger_usecase "go-gin-domain/internal/application/usecase/logger"
	domain "go-gin-domain/internal/domain/user"
)

type userRepository struct {
	logger logger_usecase.Logger
}

func NewUserRepository(logger logger_usecase.Logger) domain.UserRepository {
	return &userRepository{
		logger: logger,
	}
}

func (r *userRepository) Create(ctx context.Context, db string, user *domain.User) (*domain.User, error) {
	// 戻り値の例
	createUser := &domain.User{
		ID:        1,
		UID:       user.UID,
		LastName:  user.LastName,
		FirstName: user.FirstName,
		Email:     user.Email,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		DeletedAt: nil,
	}

	dummy := &domain.User{}
	if createUser == dummy {
		msg := fmt.Sprintf("[%s] user is nil", db)
		r.logger.Error(ctx, msg)
	}

	return createUser, nil
}

func (r *userRepository) FindAll(ctx context.Context, db string) ([]*domain.User, error) {
	// 戻り値の例
	users := []*domain.User{
		{
			ID:        1,
			UID:       "xxxx-xxxx-xxxx-0001",
			LastName:  "田中",
			FirstName: "太郎",
			Email:     "t.tanaka@example.com",
			CreatedAt: time.Time{},
			UpdatedAt: time.Time{},
			DeletedAt: nil,
		},
		{
			ID:        2,
			UID:       "xxxx-xxxx-xxxx-0002",
			LastName:  "佐藤",
			FirstName: "一郎",
			Email:     "i.satou@example.com",
			CreatedAt: time.Time{},
			UpdatedAt: time.Time{},
			DeletedAt: nil,
		},
	}

	if len(users) == 0 {
		msg := fmt.Sprintf("[%s] users is nil", db)
		r.logger.Error(ctx, msg)
	}

	return users, nil
}

func (r *userRepository) FindByUID(ctx context.Context, db string, uid string) (*domain.User, error) {
	// 戻り値の例
	var user *domain.User
	if uid == "xxxx-xxxx-xxxx-0001" {
		user = &domain.User{
			ID:        1,
			UID:       "xxxx-xxxx-xxxx-0001",
			LastName:  "田中",
			FirstName: "太郎",
			Email:     "t.tanaka@example.com",
			CreatedAt: time.Time{},
			UpdatedAt: time.Time{},
			DeletedAt: nil,
		}
	} else {
		user = nil
	}

	return user, nil
}

func (r *userRepository) Save(ctx context.Context, db string, user *domain.User) (*domain.User, error) {
	// 戻り値の例
	return user, nil
}
