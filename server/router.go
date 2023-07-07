package server

import "github.com/gin-gonic/gin"

// ルーティングをセットアップする
func (server *Server) setupRouter() {
	router := gin.New()

	router.GET("/v1/ping", server.controller.Ping)

	server.router = router
}

// 認証ミドルウェアの必要なルーティング
func (server *Server) authRouter() error {
	middleware, err := middleware.New("./credentials.json", nil)
	if err != nil {
		return err
	}

	auth := server.router.Group("/auth")
	auth.Use(middleware.AuthMiddleware(server))
	auth.GET("/v1/pong", server.controller.Pong)

	return nil
}
