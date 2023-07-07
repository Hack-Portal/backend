package server

import (
	"github.com/gin-gonic/gin"
	"github.com/hackhack-Geek-vol6/backend/middleware"
)

// ルーティングをセットアップする
func (server *Server) setupRouter() {
	router := gin.New()

	router.GET("/v1/ping", server.controller.Ping)

	server.router = router
}

// 認証ミドルウェアの必要なルーティング
func (server *Server) authRouter() error {
	auth := server.router.Group("/auth")
	auth.Use(middleware.AuthMiddleware())
	auth.GET("/v1/pong", server.controller.Pong)

	return nil
}
