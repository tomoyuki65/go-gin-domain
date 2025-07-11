package main

import (
	"go-gin-domain/internal/presentation/middleware"
	"go-gin-domain/internal/presentation/router"
	"go-gin-domain/internal/registry"
)

func main() {
	c := registry.NewController()
	m := middleware.NewMiddleware()
	r := router.SetupRouter(c, m)
	r.Run(":8080")
}
