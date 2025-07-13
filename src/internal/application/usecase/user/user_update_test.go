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

func TestUserUsecase_Update(t *testing.T) {
	// リポジトリのモック
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockRepo := mockUser.NewMockUserRepository(ctrl)

	// ロガーのモック
	mockLogger := mockLogger.NewMockLogger(ctrl)

	t.Run("正常終了すること", func(t *testing.T) {
		// モック化
		findUser := &domain_user.User{
			ID:        1,
			UID:       "xxxx-xxxx-xxxx-0001",
			LastName:  "田中",
			FirstName: "太郎",
			Email:     "t.tanaka@example.com",
			CreatedAt: time.Time{},
			UpdatedAt: time.Time{},
			DeletedAt: nil,
		}
		mockRepo.EXPECT().FindByUID(gomock.Any(), gomock.Any()).Return(findUser, nil)

		expectedUser := &domain_user.User{
			ID:        1,
			UID:       "xxxx-xxxx-xxxx-0001",
			LastName:  "佐藤",
			FirstName: "二郎",
			Email:     "z.satou@example.com",
			CreatedAt: time.Time{},
			UpdatedAt: time.Now(),
			DeletedAt: nil,
		}
		mockRepo.EXPECT().Save(gomock.Any(), gomock.Any()).Return(expectedUser, nil)

		// ユースケースのインスタンス化
		userUsecase := NewUserUsecase(mockRepo, mockLogger)

		// テストの実行
		ctx := context.Background()
		uid := "xxxx-xxxx-xxxx-0001"
		lastName := "佐藤"
		firstName := "二郎"
		email := "z.satou@example.com"
		user, err := userUsecase.Update(ctx, uid, lastName, firstName, email)

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
		assert.NotEqual(t, user.UpdatedAt, user.CreatedAt)
		assert.Nil(t, user.DeletedAt)
	})

	t.Run("対象ユーザー取得でエラーの場合にエラーを返すこと", func(t *testing.T) {
		// モック化
		err := fmt.Errorf("Internal Server Error")
		mockRepo.EXPECT().FindByUID(gomock.Any(), gomock.Any()).Return(nil, err)

		// ユースケースのインスタンス化
		userUsecase := NewUserUsecase(mockRepo, mockLogger)

		// テストの実行
		ctx := context.Background()
		uid := "xxxx-xxxx-xxxx-0001"
		lastName := "佐藤"
		firstName := "二郎"
		email := "z.satou@example.com"
		user, err := userUsecase.Update(ctx, uid, lastName, firstName, email)

		// 検証
		assert.Error(t, err)
		assert.Nil(t, user)
	})

	t.Run("対象ユーザーが存在しない場合にエラーを返すこと", func(t *testing.T) {
		// モック化
		mockRepo.EXPECT().FindByUID(gomock.Any(), gomock.Any()).Return(nil, nil)
		mockLogger.EXPECT().Error(gomock.Any(), gomock.Any()).Return()

		// ユースケースのインスタンス化
		userUsecase := NewUserUsecase(mockRepo, mockLogger)

		// テストの実行
		ctx := context.Background()
		uid := "xxxx-xxxx-xxxx-0001"
		lastName := "佐藤"
		firstName := "二郎"
		email := "z.satou@example.com"
		user, err := userUsecase.Update(ctx, uid, lastName, firstName, email)

		// 検証
		assert.Error(t, err)
		assert.Nil(t, user)
	})

	t.Run("プロフィール更新でエラーの場合にエラーを返すこと", func(t *testing.T) {
		// モック化
		findUser := &domain_user.User{
			ID:        1,
			UID:       "xxxx-xxxx-xxxx-0001",
			LastName:  "田中",
			FirstName: "太郎",
			Email:     "t.tanaka@example.com",
			CreatedAt: time.Time{},
			UpdatedAt: time.Time{},
			DeletedAt: nil,
		}
		mockRepo.EXPECT().FindByUID(gomock.Any(), gomock.Any()).Return(findUser, nil)

		// ユースケースのインスタンス化
		userUsecase := NewUserUsecase(mockRepo, mockLogger)

		// テストの実行
		ctx := context.Background()
		uid := "xxxx-xxxx-xxxx-0001"
		lastName := ""
		firstName := "二郎"
		email := "z.satou@example.com"
		user, err := userUsecase.Update(ctx, uid, lastName, firstName, email)

		// 検証
		assert.Error(t, err)
		assert.Nil(t, user)
	})
}
