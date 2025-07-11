package user

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	domain_user "go-gin-domain/internal/domain/user"
	"go-gin-domain/internal/presentation/middleware"
	mockUser "go-gin-domain/internal/application/usecase/user/mock_user"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

// テスト用Ginの初期化処理
func initTestGin() (*gin.Engine, *gin.RouterGroup) {
	r := gin.New()

	// ミドルウェアの設定
	m := middleware.NewMiddleware()
	r.Use(m.Request())
	r.Use(m.CustomLogger())
	r.Use(gin.Recovery())

	apiV1 := r.Group("/api/v1")

	return r, apiV1
}

func TestUserHandler_Create(t *testing.T) {
	// Ginのテストモードに設定
	gin.SetMode(gin.TestMode)

	// ユースケースのモック
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockUserUsecase := mockUser.NewMockUserUsecase(ctrl)

	t.Run("should return 201 Created with user json", func(t *testing.T) {
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
		mockUserUsecase.EXPECT().Create(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).Return(expectedUser, nil)

		// ルーター設定
		r, apiV1 := initTestGin()
		h := NewUserHandler(mockUserUsecase)
		apiV1.POST("/user", h.Create)

		// リクエスト設定
		path := "/api/v1/user"
		reqBody := CreateUserRequestBody{
			LastName:  "田中",
			FirstName: "太郎",
			Email:     "t.tanaka@example.com",
		}
		jsonReqBody, err := json.Marshal(reqBody)
		if err != nil {
			t.Fatal(err)
		}
		req := httptest.NewRequest(http.MethodPost, path, bytes.NewBuffer(jsonReqBody))

		// テストの実行
		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)

		// 検証
		assert.Equal(t, http.StatusCreated, w.Code)

		var data map[string]interface{}
		err = json.Unmarshal(w.Body.Bytes(), &data)
		assert.NoError(t, err)

		assert.NotContains(t, data, "id")
		assert.NotNil(t, data["uid"])
		assert.Equal(t, expectedUser.LastName, data["last_name"])
		assert.Equal(t, expectedUser.FirstName, data["first_name"])
		assert.Equal(t, expectedUser.Email, data["email"])
		assert.NotNil(t, data["created_at"])
		assert.NotNil(t, data["updated_at"])
		assert.Nil(t, data["deleted_at"])
	})

	t.Run("should return 500 Internal Server Error with error message", func(t *testing.T) {
		err := fmt.Errorf("Internal Server Error")
		mockUserUsecase.EXPECT().Create(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).Return(nil, err)

		// ルーター設定
		r, apiV1 := initTestGin()
		h := NewUserHandler(mockUserUsecase)
		apiV1.POST("/user", h.Create)

		// リクエスト設定
		path := "/api/v1/user"
		reqBody := CreateUserRequestBody{
			LastName:  "田中",
			FirstName: "太郎",
			Email:     "t.tanaka@example.com",
		}
		jsonReqBody, err := json.Marshal(reqBody)
		if err != nil {
			t.Fatal(err)
		}
		req := httptest.NewRequest(http.MethodPost, path, bytes.NewBuffer(jsonReqBody))

		// テストの実行
		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)

		// 検証
		assert.Equal(t, http.StatusInternalServerError, w.Code)
		assert.Contains(t, w.Body.String(), "Internal Server Error")
	})

	t.Run("should return 422 Unprocessable Entity with error message", func(t *testing.T) {
		// ルーター設定
		r, apiV1 := initTestGin()
		h := NewUserHandler(mockUserUsecase)
		apiV1.POST("/user", h.Create)

		// リクエスト設定
		path := "/api/v1/user"
		reqBody := CreateUserRequestBody{
			LastName:  "",
			FirstName: "太郎",
			Email:     "t.tanaka@example.com",
		}
		jsonReqBody, err := json.Marshal(reqBody)
		if err != nil {
			t.Fatal(err)
		}
		req := httptest.NewRequest(http.MethodPost, path, bytes.NewBuffer(jsonReqBody))

		// テストの実行
		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)

		// 検証
		assert.Equal(t, http.StatusUnprocessableEntity, w.Code)
		assert.Contains(t, w.Body.String(), "バリデーションエラー")
	})
}

func TestUserHandler_FindAll(t *testing.T) {
	// Ginのテストモードに設定
	gin.SetMode(gin.TestMode)

	// ユースケースのモック
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockUserUsecase := mockUser.NewMockUserUsecase(ctrl)

	t.Run("should return 200 OK with users json", func(t *testing.T) {
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
		mockUserUsecase.EXPECT().FindAll(gomock.Any()).Return(expectedUsers, nil)

		// ルーター設定
		r, apiV1 := initTestGin()
		m := middleware.NewMiddleware()
		h := NewUserHandler(mockUserUsecase)
		apiV1.GET("/users", m.Auth(), h.FindAll)

		// リクエスト設定
		path := "/api/v1/users"
		req := httptest.NewRequest(http.MethodGet, path, nil)
		req.Header.Set("Authorization", "Bearer xxxxxx")

		// テストの実行
		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)

		// 検証
		assert.Equal(t, http.StatusOK, w.Code)

		var list []map[string]interface{}
		err := json.Unmarshal(w.Body.Bytes(), &list)
		assert.NoError(t, err)

		assert.NotContains(t, list[0], "id")
		assert.Equal(t, expectedUsers[0].UID, list[0]["uid"])
		assert.Equal(t, expectedUsers[0].LastName, list[0]["last_name"])
		assert.Equal(t, expectedUsers[0].FirstName, list[0]["first_name"])
		assert.Equal(t, expectedUsers[0].Email, list[0]["email"])
		assert.NotNil(t, list[0]["created_at"])
		assert.NotNil(t, list[0]["updated_at"])
		assert.Nil(t, list[0]["deleted_at"])

		assert.NotContains(t, list[1], "id")
		assert.Equal(t, expectedUsers[1].UID, list[1]["uid"])
		assert.Equal(t, expectedUsers[1].LastName, list[1]["last_name"])
		assert.Equal(t, expectedUsers[1].FirstName, list[1]["first_name"])
		assert.Equal(t, expectedUsers[1].Email, list[1]["email"])
		assert.NotNil(t, list[1]["created_at"])
		assert.NotNil(t, list[1]["updated_at"])
		assert.Nil(t, list[1]["deleted_at"])
	})
}
