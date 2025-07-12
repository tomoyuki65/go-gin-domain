package user

import (
	"context"
	"fmt"
	"testing"
	"time"

	mockLogger "go-gin-domain/internal/application/usecase/logger/mock_logger"
	domain_user "go-gin-domain/internal/domain/user"
	mockUser "go-gin-domain/internal/infrastructure/persistence/user/mock_user_repository"

	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

// 初期処理
func init() {
	// テスト用の環境変数ファイル「.env.testing」を読み込んで使用する。
	if err := godotenv.Load("../../../../.env.testing"); err != nil {
		fmt.Println(".env.testingの読み込みに失敗しました。")
	}
}

func TestUserUsecase_FindAll(t *testing.T) {
	// リポジトリのモック
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockRepo := mockUser.NewMockUserRepository(ctrl)

	// ロガーのモック
	mockLogger := mockLogger.NewMockLogger(ctrl)

	t.Run("正常終了すること", func(t *testing.T) {
		// モック化
		expectedUsers := []*domain_user.User{
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
		mockRepo.EXPECT().FindAll(gomock.Any()).Return(expectedUsers, nil)

		// ユースケースのインスタンス化
		userUsecase := NewUserUsecase(mockRepo, mockLogger)

		// テストの実行
		ctx := context.Background()
		users, err := userUsecase.FindAll(ctx)

		// 検証
		assert.NoError(t, err)
		assert.NotNil(t, users)
		assert.Equal(t, len(expectedUsers), len(users))

		assert.Equal(t, expectedUsers[0].ID, users[0].ID)
		assert.Equal(t, expectedUsers[0].UID, users[0].UID)
		assert.Equal(t, expectedUsers[0].LastName, users[0].LastName)
		assert.Equal(t, expectedUsers[0].FirstName, users[0].FirstName)
		assert.Equal(t, expectedUsers[0].Email, users[0].Email)
		assert.Equal(t, expectedUsers[0].CreatedAt, users[0].CreatedAt)
		assert.Equal(t, expectedUsers[0].UpdatedAt, users[0].UpdatedAt)
		assert.Equal(t, expectedUsers[0].DeletedAt, users[0].DeletedAt)

		assert.Equal(t, expectedUsers[1].ID, users[1].ID)
		assert.Equal(t, expectedUsers[1].UID, users[1].UID)
		assert.Equal(t, expectedUsers[1].LastName, users[1].LastName)
		assert.Equal(t, expectedUsers[1].FirstName, users[1].FirstName)
		assert.Equal(t, expectedUsers[1].Email, users[1].Email)
		assert.Equal(t, expectedUsers[1].CreatedAt, users[1].CreatedAt)
		assert.Equal(t, expectedUsers[1].UpdatedAt, users[1].UpdatedAt)
		assert.Equal(t, expectedUsers[1].DeletedAt, users[1].DeletedAt)
	})
}
