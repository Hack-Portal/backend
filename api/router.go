package api

import (
	"github.com/gin-gonic/gin"
)

// ルーティングをセットアップする
func (server *Server) setupRouter() {
	router := gin.New()
	server.router = router

	server.publicRouter()
	server.authRouter()

}

// 認証を必要としないルーティング
func (server *Server) publicRouter() {
	public := server.router.Group("/v1")

	public.GET("/ping", server.Ping)
	public.POST("/hackathons", server.CreateHackathon)
	public.GET("/hackathons", server.ListHackathons)
	public.GET("/hackathons/:hackathon_id", server.GetHackathon)
	public.GET("/accounts/:id", server.GetAccount)
}

// 認証ミドルウェアの必要なルーティング
func (server *Server) authRouter() {
	auth := server.router.Group("/v1")
	auth.Use(AuthMiddleware())
	// アカウント作成

	auth.POST("/accounts", server.CreateAccount)
	auth.POST("/rooms", server.CreateRoom)
	auth.POST("/rooms/:room_id", server.AddAccountInRoom)
	auth.GET("/rooms", server.ListRooms)
	auth.GET("/rooms/:room_id", server.GetRoom)
}
