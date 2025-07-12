package main

import (
	"fmt"
	"log/slog"
	"os"

	"go-gin-domain/internal/presentation/middleware"
	"go-gin-domain/internal/presentation/router"
	"go-gin-domain/internal/registry"

	"github.com/joho/godotenv"
)

func main() {
	// .env ファイルの読み込み
	err := godotenv.Load()
	if err != nil {
		slog.Error(".envファイルの読み込みに失敗しました。")
	}

	// ポート番号の設定
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	startPort := fmt.Sprintf(":%s", port)

	// サーバー起動ログ出力
	env := os.Getenv("ENV")
	slog.Info(fmt.Sprintf("[ENV=%s] Start Gin Server Port: %s", env, port))

	// サーバー起動
	c := registry.NewController()
	m := middleware.NewMiddleware()
	r := router.SetupRouter(c, m)
	r.Run(startPort)
}
