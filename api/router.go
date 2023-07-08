package api

import (
	"github.com/gin-gonic/gin"
)

// ルーティングをセットアップする
func (server *Server) setupRouter() {
	router := gin.New()

	router.GET("/v1/ping", server.Ping)

	server.router = router
}

// 認証ミドルウェアの必要なルーティング
func (server *Server) authRouter() error {
	auth := server.router.Group("/auth")
	auth.Use(AuthMiddleware())
	auth.GET("/v1/pong", server.Pong)

	return nil
}
