package user

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	domain_user "go-gin-domain/internal/domain/user"
	"go-gin-domain/internal/presentation/middleware"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// ユースケースのモック定義
type MockUserUsecase struct {
	mock.Mock
}

func (m *MockUserUsecase) Create(ctx context.Context, lastName, firstName, email string) (*domain_user.User, error) {
	args := m.Called(ctx, lastName, firstName, email)
	return args.Get(0).(*domain_user.User), args.Error(1)
}

func (m *MockUserUsecase) FindAll(ctx context.Context) ([]*domain_user.User, error) {
	args := m.Called(ctx)

	var users []*domain_user.User
	if arg := args.Get(0); arg != nil {
		users = arg.([]*domain_user.User)
	}

	return users, args.Error(1)
}

func (m *MockUserUsecase) FindByUID(ctx context.Context, uid string) (*domain_user.User, error) {
	args := m.Called(ctx, uid)
	return args.Get(0).(*domain_user.User), args.Error(1)
}

func (m *MockUserUsecase) Update(ctx context.Context, uid, lastName, firstName, email string) (*domain_user.User, error) {
	args := m.Called(ctx, uid, lastName, firstName, email)
	return args.Get(0).(*domain_user.User), args.Error(1)
}

func (m *MockUserUsecase) Delete(ctx context.Context, uid string) (*domain_user.User, error) {
	args := m.Called(ctx, uid)
	return args.Get(0).(*domain_user.User), args.Error(1)
}

// テスト用Ginの初期化処理
func initTestGin() (*gin.Engine, *gin.RouterGroup) {
	r := gin.New()

	// ミドルウェアの設定
	m := middleware.NewMiddleware()
	r.Use(m.Request())
	r.Use(m.CustomLogger())
	r.Use(gin.Recovery())

	apiV1 := r.Group("/api/v1")
	apiV1.Use(m.Auth())

	return r, apiV1
}

func TestUserHandler_FindAll(t *testing.T) {
	// Ginのテストモードに設定
	gin.SetMode(gin.TestMode)

	t.Run("should return 200 OK with users json", func(t *testing.T) {
		// モック化
		mockUserUsecase := new(MockUserUsecase)
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
		mockUserUsecase.On("FindAll", mock.Anything).Return(expectedUsers, nil)

		// ルーター設定
		r, apiV1 := initTestGin()
		h := NewUserHandler(mockUserUsecase)
		apiV1.GET("/users", h.FindAll)

		// リクエスト設定
		path := "/api/v1/users"
		req, _ := http.NewRequest(http.MethodGet, path, nil)
		req.Header.Set("Authorization", "Bearer xxxxxx")

		// テストの実行
		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)

		// 検証
		assert.Equal(t, http.StatusOK, w.Code)

		var responseUsers []*domain_user.User
		err := json.Unmarshal(w.Body.Bytes(), &responseUsers)
		assert.NoError(t, err)
		assert.Equal(t, len(expectedUsers), len(responseUsers))

		assert.Equal(t, expectedUsers[0].UID, responseUsers[0].UID)
		assert.Equal(t, expectedUsers[0].LastName, responseUsers[0].LastName)
		assert.Equal(t, expectedUsers[0].FirstName, responseUsers[0].FirstName)
		assert.Equal(t, expectedUsers[0].Email, responseUsers[0].Email)
		assert.Equal(t, expectedUsers[0].CreatedAt, responseUsers[0].CreatedAt)
		assert.Equal(t, expectedUsers[0].UpdatedAt, responseUsers[0].UpdatedAt)
		assert.Equal(t, expectedUsers[0].DeletedAt, responseUsers[0].DeletedAt)

		assert.Equal(t, expectedUsers[1].UID, responseUsers[1].UID)
		assert.Equal(t, expectedUsers[1].LastName, responseUsers[1].LastName)
		assert.Equal(t, expectedUsers[1].FirstName, responseUsers[1].FirstName)
		assert.Equal(t, expectedUsers[1].Email, responseUsers[1].Email)
		assert.Equal(t, expectedUsers[1].CreatedAt, responseUsers[1].CreatedAt)
		assert.Equal(t, expectedUsers[1].UpdatedAt, responseUsers[1].UpdatedAt)
		assert.Equal(t, expectedUsers[1].DeletedAt, responseUsers[1].DeletedAt)

		mockUserUsecase.AssertExpectations(t)
	})
}
