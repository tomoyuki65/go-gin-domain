package user

import (
	"context"
	"testing"
	"time"

	domain_user "go-gin-domain/internal/domain/user"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// リポジトリのモック定義
type MockUserRepository struct {
	mock.Mock
}

func (m *MockUserRepository) Create(ctx context.Context, user *domain_user.User) (*domain_user.User, error) {
	args := m.Called(ctx, user)
	return args.Get(0).(*domain_user.User), args.Error(1)
}
func (m *MockUserRepository) FindAll(ctx context.Context) ([]*domain_user.User, error) {
	args := m.Called(ctx)

	var users []*domain_user.User
	if arg := args.Get(0); arg != nil {
		users = arg.([]*domain_user.User)
	}

	return users, args.Error(1)
}

func (m *MockUserRepository) FindByUID(ctx context.Context, uid string) (*domain_user.User, error) {
	args := m.Called(ctx, uid)
	return args.Get(0).(*domain_user.User), args.Error(1)
}

func (m *MockUserRepository) Save(ctx context.Context, user *domain_user.User) (*domain_user.User, error) {
	args := m.Called(ctx, user)
	return args.Get(0).(*domain_user.User), args.Error(1)
}

// ロガーのモック定義
type MockLogger struct {
	mock.Mock
}

func (m *MockLogger) Info(ctx context.Context, msg string) {
	m.Called(ctx, msg)
}
func (m *MockLogger) Warn(ctx context.Context, msg string) {
	m.Called(ctx, msg)
}
func (m *MockLogger) Error(ctx context.Context, msg string) {
	m.Called(ctx, msg)
}

func setupMockLogger(t *testing.T) *MockLogger {
	t.Helper() // これを呼び出すとテスト失敗時にこの関数がコールスタックに表示されなくなる

	mockLogger := new(MockLogger)
	logLevels := []string{"Info", "Warn", "Error"}

	for _, level := range logLevels {
		mockLogger.On(level, mock.Anything).Return()
	}

	return mockLogger
}

func TestUserUsecase_FindAll(t *testing.T) {
	t.Run("should return all users successfully", func(t *testing.T) {
		// モック化
		mockRepo := new(MockUserRepository)
		ctx := context.Background()
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
		mockRepo.On("FindAll", ctx).Return(expectedUsers, nil)
		mockLogger := setupMockLogger(t)

		// ユースケースのインスタンス化
		userUsecase := NewUserUsecase(mockRepo, mockLogger)

		// テストの実行
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

		mockRepo.AssertExpectations(t)
	})
}
