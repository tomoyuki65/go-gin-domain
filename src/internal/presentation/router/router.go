package router

import (
	"go-gin-domain/internal/presentation/middleware"
	"go-gin-domain/internal/registry"

	"github.com/gin-gonic/gin"
)

func SetupRouter(c *registry.Controller, m *middleware.Middleware) *gin.Engine {
	r := gin.New()

	// 共通ミドルウェアの適用
	r.Use(m.Request())
	r.Use(m.CustomLogger())
	r.Use(gin.Recovery())

	// ルーティングの設定
	apiV1 := r.Group("/api/v1")
	apiV1.POST("/user", c.User.Create)
	apiV1.GET("/users", m.Auth(), c.User.FindAll)
	apiV1.GET("/user/:uid", m.Auth(), c.User.FindByUID)
	apiV1.PUT("/user/:uid", m.Auth(), c.User.Update)
	apiV1.DELETE("/user/:uid", m.Auth(), c.User.Delete)

	return r
}
