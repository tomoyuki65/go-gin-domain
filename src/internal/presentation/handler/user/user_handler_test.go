package user

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	mockUser "go-gin-domain/internal/application/usecase/user/mock_user"
	domain_user "go-gin-domain/internal/domain/user"
	"go-gin-domain/internal/presentation/middleware"

	"github.com/gin-gonic/gin"
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

	t.Run("ステータス201で正常終了すること", func(t *testing.T) {
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

	t.Run("ユースケースでエラーが発生した場合にステータス500を返すこと", func(t *testing.T) {
		// モック化
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

	t.Run("バリデーションチェックでエラーの場合にステータス422を返すこと", func(t *testing.T) {
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

	t.Run("ステータス200で正常終了すること", func(t *testing.T) {
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

	t.Run("ユースケースでエラーが発生した場合にステータス500を返すこと", func(t *testing.T) {
		// モック化
		err := fmt.Errorf("Internal Server Error")
		mockUserUsecase.EXPECT().FindAll(gomock.Any()).Return(nil, err)

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
		assert.Equal(t, http.StatusInternalServerError, w.Code)
		assert.Contains(t, w.Body.String(), "Internal Server Error")
	})
}

func TestUserHandler_FindByUID(t *testing.T) {
	// Ginのテストモードに設定
	gin.SetMode(gin.TestMode)

	// ユースケースのモック
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockUserUsecase := mockUser.NewMockUserUsecase(ctrl)

	t.Run("ステータス200で正常終了すること", func(t *testing.T) {
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
		mockUserUsecase.EXPECT().FindByUID(gomock.Any(), gomock.Any()).Return(expectedUser, nil)

		// ルーター設定
		r, apiV1 := initTestGin()
		m := middleware.NewMiddleware()
		h := NewUserHandler(mockUserUsecase)
		apiV1.GET("/user/:uid", m.Auth(), h.FindByUID)

		// リクエスト設定
		path := "/api/v1/user/xxxx-xxxx-xxxx-0001"
		req := httptest.NewRequest(http.MethodGet, path, nil)
		req.Header.Set("Authorization", "Bearer xxxxxx")

		// テストの実行
		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)

		// 検証
		assert.Equal(t, http.StatusOK, w.Code)

		var data map[string]interface{}
		err := json.Unmarshal(w.Body.Bytes(), &data)
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

	t.Run("対象ユーザーが存在しない場合にステータス200で空のオブジェクトを返すこと", func(t *testing.T) {
		// モック化
		mockUserUsecase.EXPECT().FindByUID(gomock.Any(), gomock.Any()).Return(nil, nil)

		// ルーター設定
		r, apiV1 := initTestGin()
		m := middleware.NewMiddleware()
		h := NewUserHandler(mockUserUsecase)
		apiV1.GET("/user/:uid", m.Auth(), h.FindByUID)

		// リクエスト設定
		path := "/api/v1/user/xxxx-xxxx-xxxx-0002"
		req := httptest.NewRequest(http.MethodGet, path, nil)
		req.Header.Set("Authorization", "Bearer xxxxxx")

		// テストの実行
		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)

		// 検証
		assert.Equal(t, http.StatusOK, w.Code)

		var data map[string]interface{}
		err := json.Unmarshal(w.Body.Bytes(), &data)
		assert.NoError(t, err)
		assert.Empty(t, data)
	})

	t.Run("ユースケースでエラーが発生した場合にステータス500を返すこと", func(t *testing.T) {
		// モック化
		err := fmt.Errorf("Internal Server Error")
		mockUserUsecase.EXPECT().FindByUID(gomock.Any(), gomock.Any()).Return(nil, err)

		// ルーター設定
		r, apiV1 := initTestGin()
		m := middleware.NewMiddleware()
		h := NewUserHandler(mockUserUsecase)
		apiV1.GET("/user/:uid", m.Auth(), h.FindByUID)

		// リクエスト設定
		path := "/api/v1/user/xxxx-xxxx-xxxx-0002"
		req := httptest.NewRequest(http.MethodGet, path, nil)
		req.Header.Set("Authorization", "Bearer xxxxxx")

		// テストの実行
		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)

		// 検証
		assert.Equal(t, http.StatusInternalServerError, w.Code)
		assert.Contains(t, w.Body.String(), "Internal Server Error")
	})

	t.Run("バリデーションチェックでエラーの場合にステータス422を返すこと", func(t *testing.T) {
		// ルーター設定
		r, apiV1 := initTestGin()
		m := middleware.NewMiddleware()
		h := NewUserHandler(mockUserUsecase)
		apiV1.GET("/user/:uid", m.Auth(), h.FindByUID)

		// リクエスト設定
		path := "/api/v1/user/　"
		req := httptest.NewRequest(http.MethodGet, path, nil)
		req.Header.Set("Authorization", "Bearer xxxxxx")

		// テストの実行
		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)

		// 検証
		assert.Equal(t, http.StatusUnprocessableEntity, w.Code)
		assert.Contains(t, w.Body.String(), "バリデーションエラー")
	})
}

func TestUserHandler_Update(t *testing.T) {
	// Ginのテストモードに設定
	gin.SetMode(gin.TestMode)

	// ユースケースのモック
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockUserUsecase := mockUser.NewMockUserUsecase(ctrl)

	t.Run("ステータス200で正常終了すること", func(t *testing.T) {
		// モック化
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
		mockUserUsecase.EXPECT().Update(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).Return(expectedUser, nil)

		// ルーター設定
		r, apiV1 := initTestGin()
		h := NewUserHandler(mockUserUsecase)
		apiV1.PUT("/user/:uid", h.Update)

		// リクエスト設定
		path := "/api/v1/user/xxxx-xxxx-xxxx-0001"
		reqBody := UpdateUserRequestBody{
			LastName:  "佐藤",
			FirstName: "二郎",
			Email:     "z.satou@example.com",
		}
		jsonReqBody, err := json.Marshal(reqBody)
		if err != nil {
			t.Fatal(err)
		}
		req := httptest.NewRequest(http.MethodPut, path, bytes.NewBuffer(jsonReqBody))
		req.Header.Set("Authorization", "Bearer xxxxxx")

		// テストの実行
		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)

		// 検証
		assert.Equal(t, http.StatusOK, w.Code)

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
		assert.NotEqual(t, data["updated_at"], data["created_at"])
		assert.Nil(t, data["deleted_at"])
	})

	t.Run("ユースケースでエラーが発生した場合にステータス500を返すこと", func(t *testing.T) {
		// モック化
		err := fmt.Errorf("Internal Server Error")
		mockUserUsecase.EXPECT().Update(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).Return(nil, err)

		// ルーター設定
		r, apiV1 := initTestGin()
		h := NewUserHandler(mockUserUsecase)
		apiV1.PUT("/user/:uid", h.Update)

		// リクエスト設定
		path := "/api/v1/user/xxxx-xxxx-xxxx-0001"
		reqBody := UpdateUserRequestBody{
			LastName:  "佐藤",
			FirstName: "二郎",
			Email:     "z.satou@example.com",
		}
		jsonReqBody, err := json.Marshal(reqBody)
		if err != nil {
			t.Fatal(err)
		}
		req := httptest.NewRequest(http.MethodPut, path, bytes.NewBuffer(jsonReqBody))
		req.Header.Set("Authorization", "Bearer xxxxxx")

		// テストの実行
		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)

		// 検証
		assert.Equal(t, http.StatusInternalServerError, w.Code)
		assert.Contains(t, w.Body.String(), "Internal Server Error")
	})

	t.Run("UIDのバリデーションチェックでエラーの場合にステータス422を返すこと", func(t *testing.T) {
		// ルーター設定
		r, apiV1 := initTestGin()
		h := NewUserHandler(mockUserUsecase)
		apiV1.PUT("/user/:uid", h.Update)

		// リクエスト設定
		path := "/api/v1/user/　"
		reqBody := UpdateUserRequestBody{
			LastName:  "佐藤",
			FirstName: "二郎",
			Email:     "z.satou@example.com",
		}
		jsonReqBody, err := json.Marshal(reqBody)
		if err != nil {
			t.Fatal(err)
		}
		req := httptest.NewRequest(http.MethodPut, path, bytes.NewBuffer(jsonReqBody))
		req.Header.Set("Authorization", "Bearer xxxxxx")

		// テストの実行
		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)

		// 検証
		assert.Equal(t, http.StatusUnprocessableEntity, w.Code)
		assert.Contains(t, w.Body.String(), "バリデーションエラー")
	})

	t.Run("リクエストボディのバリデーションチェックでエラーの場合にステータス422を返すこと", func(t *testing.T) {
		// ルーター設定
		r, apiV1 := initTestGin()
		h := NewUserHandler(mockUserUsecase)
		apiV1.PUT("/user/:uid", h.Update)

		// リクエスト設定
		path := "/api/v1/user/xxxx-xxxx-xxxx-0001"
		reqBody := UpdateUserRequestBody{
			LastName:  "佐藤",
			FirstName: "",
			Email:     "z.satou@example.com",
		}
		jsonReqBody, err := json.Marshal(reqBody)
		if err != nil {
			t.Fatal(err)
		}
		req := httptest.NewRequest(http.MethodPut, path, bytes.NewBuffer(jsonReqBody))
		req.Header.Set("Authorization", "Bearer xxxxxx")

		// テストの実行
		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)

		// 検証
		assert.Equal(t, http.StatusUnprocessableEntity, w.Code)
		assert.Contains(t, w.Body.String(), "バリデーションエラー")
	})
}

func TestUserHandler_Delete(t *testing.T) {
	// Ginのテストモードに設定
	gin.SetMode(gin.TestMode)

	// ユースケースのモック
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockUserUsecase := mockUser.NewMockUserUsecase(ctrl)

	t.Run("ステータス200で正常終了すること", func(t *testing.T) {
		// モック化
		date := time.Now()
		dateString := date.Format("2006-01-02 15:04:05")
		updateEmail := "z.satou@example.com" + dateString
		expectedUser := &domain_user.User{
			ID:        1,
			UID:       "xxxx-xxxx-xxxx-0001",
			LastName:  "佐藤",
			FirstName: "二郎",
			Email:     updateEmail,
			CreatedAt: time.Time{},
			UpdatedAt: date,
			DeletedAt: &date,
		}
		mockUserUsecase.EXPECT().Delete(gomock.Any(), gomock.Any()).Return(expectedUser, nil)

		// ルーター設定
		r, apiV1 := initTestGin()
		h := NewUserHandler(mockUserUsecase)
		apiV1.DELETE("/user/:uid", h.Delete)

		// リクエスト設定
		path := "/api/v1/user/xxxx-xxxx-xxxx-0001"
		req := httptest.NewRequest(http.MethodDelete, path, nil)
		req.Header.Set("Authorization", "Bearer xxxxxx")

		// テストの実行
		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)

		// 検証
		assert.Equal(t, http.StatusOK, w.Code)

		var data map[string]interface{}
		err := json.Unmarshal(w.Body.Bytes(), &data)
		assert.NoError(t, err)

		assert.NotContains(t, data, "id")
		assert.NotNil(t, data["uid"])
		assert.Equal(t, expectedUser.LastName, data["last_name"])
		assert.Equal(t, expectedUser.FirstName, data["first_name"])
		assert.Equal(t, expectedUser.Email, data["email"])
		assert.NotNil(t, data["created_at"])
		assert.NotNil(t, data["updated_at"])
		assert.NotEqual(t, data["updated_at"], data["created_at"])
		assert.NotNil(t, data["deleted_at"])
		assert.Equal(t, data["deleted_at"], data["updated_at"])
	})

	t.Run("ユースケースでエラーが発生した場合にステータス500を返すこと", func(t *testing.T) {
		// モック化
		err := fmt.Errorf("Internal Server Error")
		mockUserUsecase.EXPECT().Delete(gomock.Any(), gomock.Any()).Return(nil, err)

		// ルーター設定
		r, apiV1 := initTestGin()
		h := NewUserHandler(mockUserUsecase)
		apiV1.DELETE("/user/:uid", h.Delete)

		// リクエスト設定
		path := "/api/v1/user/xxxx-xxxx-xxxx-0001"
		req := httptest.NewRequest(http.MethodDelete, path, nil)
		req.Header.Set("Authorization", "Bearer xxxxxx")

		// テストの実行
		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)

		// 検証
		assert.Equal(t, http.StatusInternalServerError, w.Code)
		assert.Contains(t, w.Body.String(), "Internal Server Error")
	})

	t.Run("バリデーションチェックでエラーの場合にステータス422を返すこと", func(t *testing.T) {
		// ルーター設定
		r, apiV1 := initTestGin()
		h := NewUserHandler(mockUserUsecase)
		apiV1.DELETE("/user/:uid", h.Delete)

		// リクエスト設定
		path := "/api/v1/user/　"
		req := httptest.NewRequest(http.MethodDelete, path, nil)
		req.Header.Set("Authorization", "Bearer xxxxxx")

		// テストの実行
		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)

		// 検証
		assert.Equal(t, http.StatusUnprocessableEntity, w.Code)
		assert.Contains(t, w.Body.String(), "バリデーションエラー")
	})
}
