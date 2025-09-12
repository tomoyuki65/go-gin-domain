//go:build integration

package user

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	usecase_user "go-gin-domain/internal/application/usecase/user"
	"go-gin-domain/internal/infrastructure/database"
	"go-gin-domain/internal/infrastructure/logger"
	persistence_user "go-gin-domain/internal/infrastructure/persistence/user"
	"go-gin-domain/internal/presentation/middleware"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

// テスト用Ginの初期化処理
func initTestGin() *gin.Engine {
	// Ginのテストモードに設定
	gin.SetMode(gin.TestMode)

	// ハンドラーのインスタンス化
	ctx := context.Background()
	logger := logger.NewSlogLogger()
	cfg := database.DummyConfig{
		Dummy: "dummy",
	}
	db_dummy, err := database.NewDummyConnection(cfg, logger)
	if err != nil {
		msg := fmt.Sprintf("エラー: %s", err.Error())
		logger.Error(ctx, msg)
	}
	userRepo := persistence_user.NewUserRepository(logger)
	userUsecase := usecase_user.NewUserUsecase(db_dummy, userRepo, logger)
	h := NewUserHandler(userUsecase)

	// ルーターの初期化
	r := gin.New()

	// ミドルウェアの設定
	m := middleware.NewMiddleware()
	r.Use(m.Request())
	r.Use(m.CustomLogger())
	r.Use(gin.Recovery())

	// ルーティング設定
	apiV1 := r.Group("/api/v1")
	apiV1.POST("/user", h.Create)
	apiV1.GET("/users", m.Auth(), h.FindAll)
	apiV1.GET("/user/:uid", m.Auth(), h.FindByUID)
	apiV1.PUT("/user/:uid", m.Auth(), h.Update)
	apiV1.DELETE("/user/:uid", m.Auth(), h.Delete)

	return r
}

func TestUserHandler_Create_Integration(t *testing.T) {
	// ルーター設定
	r := initTestGin()

	t.Run("Create_Integrationが正常終了すること", func(t *testing.T) {
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
		assert.Equal(t, reqBody.LastName, data["last_name"])
		assert.Equal(t, reqBody.FirstName, data["first_name"])
		assert.Equal(t, reqBody.Email, data["email"])
		assert.NotNil(t, data["created_at"])
		assert.NotNil(t, data["updated_at"])
		assert.Nil(t, data["deleted_at"])
	})
}

/******************************
 * ベンチマーク関数を追加
 ******************************/
func BenchmarkUserHandler_Create_Integration(b *testing.B) {
	// ルーター設定
	r := initTestGin()

	// ログ出力を無効化する
	gin.DefaultWriter = io.Discard

	// リクエストパス設定
	path := "/api/v1/user"

	// タイマーリセット（セットアップの時間を計測から除外）
	b.ResetTimer()

	// --- ベンチマーク実行 (b.N回ループ) ---
	// b.N は go test が自動的に調整するループ回数
	for i := 0; i < b.N; i++ {
		// ループごとに新しいリクエストとレスポンスレコーダーを作成
		// これらは１回のHTTPリクエストに相当するため、ループ内で都度生成する
		lastName := fmt.Sprintf("田中%d", i)
		firstName := fmt.Sprintf("太郎%d", i)
		email := fmt.Sprintf("t.tanaka%d@example.com", i)
		reqBody := CreateUserRequestBody{
			LastName:  lastName,
			FirstName: firstName,
			Email:     email,
		}
		jsonReqBody, err := json.Marshal(reqBody)
		if err != nil {
			b.Fatal(err)
		}
		w := httptest.NewRecorder()
		req := httptest.NewRequest(http.MethodPost, path, bytes.NewBuffer(jsonReqBody))

		// リクエスト実行
		r.ServeHTTP(w, req)
	}
}
