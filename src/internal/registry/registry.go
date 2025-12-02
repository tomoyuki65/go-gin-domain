package registry

import (
	"context"
	"fmt"

	usecase_post "go-gin-domain/internal/application/usecase/post"
	usecase_user "go-gin-domain/internal/application/usecase/user"
	"go-gin-domain/internal/infrastructure/database"
	"go-gin-domain/internal/infrastructure/logger"
	persistence_post "go-gin-domain/internal/infrastructure/persistence/post"
	persistence_user "go-gin-domain/internal/infrastructure/persistence/user"
	handler_post "go-gin-domain/internal/presentation/handler/post"
	handler_user "go-gin-domain/internal/presentation/handler/user"
)

// ハンドラーをまとめるコントローラー構造体
type Controller struct {
	User handler_user.UserHandler
	Post handler_post.PostHandler
}

func NewController() *Controller {
	// コンテキスト
	ctx := context.Background()

	// ロガー設定
	logger := logger.NewSlogLogger()

	// DB設定（今回はダミー設定とする）
	cfg := database.DummyConfig{
		Dummy: "dummy",
	}
	db_dummy, err := database.NewDummyConnection(cfg, logger)
	if err != nil {
		msg := fmt.Sprintf("エラー: %s", err.Error())
		logger.Error(ctx, msg)
	}

	// userドメインのハンドラー設定
	userRepo := persistence_user.NewUserRepository(logger)
	userUsecase := usecase_user.NewUserUsecase(db_dummy, userRepo, logger)
	userHandler := handler_user.NewUserHandler(userUsecase)

	// postドメインのハンドラー設定
	postRepo := persistence_post.NewPostRepository(logger)
	postUsecase := usecase_post.NewPostUsecase(db_dummy, postRepo, logger)
	postHandler := handler_post.NewPostHandler(postUsecase)

	return &Controller{
		User: userHandler,
		Post: postHandler,
	}
}
