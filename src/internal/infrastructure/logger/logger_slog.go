package logger

import (
	"context"
	"log/slog"
	"os"

	logger_usecase "go-gin-domain/internal/application/usecase/logger"
	"go-gin-domain/internal/presentation/middleware"
)

// slogの設定
type SlogHandler struct {
	slog.Handler
}

func (h *SlogHandler) Handle(ctx context.Context, r slog.Record) error {
	requestId, ok := ctx.Value(middleware.RequestId).(string)
	if ok {
		r.AddAttrs(slog.Attr{Key: "requestId", Value: slog.String("requestId", requestId).Value})
	}

	xRequestSource, ok := ctx.Value(middleware.XRequestSource).(string)
	if ok {
		r.AddAttrs(slog.Attr{Key: "xRequestSource", Value: slog.String("xRequestSource", xRequestSource).Value})
	}

	uid, ok := ctx.Value(middleware.UID).(string)
	if ok {
		r.AddAttrs(slog.Attr{Key: "UID", Value: slog.String("UID", uid).Value})
	}

	return h.Handler.Handle(ctx, r)
}

var slogHandler = &SlogHandler{
	slog.NewTextHandler(os.Stdout, nil),
}
var logger = slog.New(slogHandler)

// ロガーの設定
type slogLogger struct{}

func NewSlogLogger() logger_usecase.Logger {
	return &slogLogger{}
}

func (l *slogLogger) Info(ctx context.Context, message string) {
	logger.InfoContext(ctx, message)
}

func (l *slogLogger) Warn(ctx context.Context, message string) {
	logger.WarnContext(ctx, message)
}

func (l *slogLogger) Error(ctx context.Context, message string) {
	logger.ErrorContext(ctx, message)
}
