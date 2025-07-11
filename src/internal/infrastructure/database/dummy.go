package database

import (
	"context"
	"fmt"

	"go-gin-domain/internal/application/usecase/logger"
)

type DummyConfig struct {
	Dummy string
}

func NewDummyConnection(cfg DummyConfig, logger logger.Logger) (string, error) {
	dsn := fmt.Sprintf("dummy=%s", cfg.Dummy)

	// ログ出力
	ctx := context.Background()
	msg := fmt.Sprintf("Successfully connected to %s", "Dummy")
	logger.Info(ctx, msg)

	return dsn, nil
}
