package middleware

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type contextKey string

const (
	RequestId      contextKey = "Request-Id"
	XRequestSource contextKey = "X-Request-Source"
	UID            contextKey = "UID"
)

type Middleware struct{}

func NewMiddleware() *Middleware {
	return &Middleware{}
}

// リクエスト用
func (m *Middleware) Request() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 一意のIDを取得
		uuid := uuid.New().String()

		// 共通コンテキストにX-Request-Idを設定
		ctx := c.Request.Context()
		ctx = context.WithValue(ctx, RequestId, uuid)

		// リクエストヘッダーからX-Request-Sourceを取得
		xRequestSource := c.GetHeader(string(XRequestSource))
		if xRequestSource == "" {
			xRequestSource = "-"
		}

		// 共通コンテキストにX-Request-Sourceを設定
		ctx = context.WithValue(ctx, XRequestSource, xRequestSource)

		// 共通コンテキストの設定
		c.Request = c.Request.WithContext(ctx)

		c.Next()
	}
}

// カスタムロガー
func (m *Middleware) CustomLogger() gin.HandlerFunc {
	return gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
		// Request-Idの取得
		requestID, ok := param.Request.Context().Value(RequestId).(string)
		if !ok {
			requestID = "-"
		}

		// カラーの取得
		var statusColor, methodColor, resetColor string
		if param.IsOutputColor() {
			statusColor = param.StatusCodeColor()
			methodColor = param.MethodColor()
			resetColor = param.ResetColor()
		}

		// ログのフォーマットを定義
		// [GIN] 2006/01/02 - 15:04:05 | {Request-Id} | 200 | 1.2345ms | 127.0.0.1 | GET /path
		return fmt.Sprintf("[GIN] %v| %s |%s %3d %s| %13v | %15s |%s %-7s %s %#v\n%s",
			param.TimeStamp.Format("2006/01/02 - 15:04:05"),
			requestID,
			statusColor,
			param.StatusCode,
			resetColor,
			param.Latency,
			param.ClientIP,
			methodColor,
			param.Method,
			resetColor,
			param.Path,
			param.ErrorMessage,
		)
	})
}

// 認証用
func (m *Middleware) Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Bearerトークン取得
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"message": "認証用トークンが設定されていません。",
			})
			return
		}
		token := strings.TrimPrefix(authHeader, "Bearer ")
		if token == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"message": "認証用トークンが設定されていません。",
			})
			return
		}

		// TODO: 認証チェックを追加する
		// TODO: 認証済みならuidを取得
		uid := "-"

		// 共通コンテキストにuidを設定
		ctx := c.Request.Context()
		ctx = context.WithValue(ctx, UID, uid)

		// 共通コンテキストの設定
		c.Request = c.Request.WithContext(ctx)

		c.Next()
	}
}
