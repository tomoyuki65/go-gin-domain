package user

import (
	"context"
	"fmt"
	"testing"
	"time"

	mockLogger "go-gin-domain/internal/application/usecase/logger/mock_logger"
	domain_user "go-gin-domain/internal/domain/user"
	mockUser "go-gin-domain/internal/domain/user/mock_user_repository"

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

func TestUserUsecase_Create(t *testing.T) {
	// リポジトリのモック
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockRepo := mockUser.NewMockUserRepository(ctrl)

	// ロガーのモック
	mockLogger := mockLogger.NewMockLogger(ctrl)

	t.Run("正常終了すること", func(t *testing.T) {
		// モック化
		expectedUser := &domain_user.User{
			ID:        1,
			UID:       "xxxx-xxxx-xxxx-0001",
			LastName:  "田中",
			FirstName: "太郎",
			Email:     "t.tanaka@example.com",
			CreatedAt: time.Time{},
			UpdatedAt: time.Time{},
			DeletedAt: nil,
		}
		mockRepo.EXPECT().Create(gomock.Any(), gomock.Any()).Return(expectedUser, nil)

		// ユースケースのインスタンス化
		userUsecase := NewUserUsecase(mockRepo, mockLogger)

		// テストの実行
		ctx := context.Background()
		lastName := "田中"
		firstName := "太郎"
		email := "t.tanaka@example.com"
		user, err := userUsecase.Create(ctx, lastName, firstName, email)

		// 検証
		assert.NoError(t, err)
		assert.NotNil(t, user)

		assert.Equal(t, expectedUser.ID, user.ID)
		assert.Equal(t, expectedUser.UID, user.UID)
		assert.Equal(t, expectedUser.LastName, user.LastName)
		assert.Equal(t, expectedUser.FirstName, user.FirstName)
		assert.Equal(t, expectedUser.Email, user.Email)
		assert.NotNil(t, user.CreatedAt)
		assert.NotNil(t, user.UpdatedAt)
		assert.Nil(t, user.DeletedAt)
	})
}
