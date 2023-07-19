package api

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func (server *Server) setupRouter() {
	router := gin.Default()
	server.router = router
	server.setUpCors()
	server.publicRouter()
	server.authRouter()

}
func (server *Server) setUpCors() {
	// server.router.Use(cors.New(
	// 	cors.Config{
	// 		AllowOrigins: []string{"https://frontend-3muyo7jtb-qirenqiantian367-gmailcom-s-team.vercel.app/"},
	// 		AllowMethods: []string{"GET", "Fetch", "POST", "Delete", "PUT"},
	// 		AllowHeaders: []string{
	// 			"Access-Control-Allow-Credentials",
	// 			"Access-Control-Allow-Headers",
	// 			"Content-Type",
	// 			"Content-Length",
	// 			"Accept-Encoding",
	// 			"Authorization",
	// 		},
	// 		AllowCredentials: false,
	// 		MaxAge:           24 * time.Hour,
	// 	}))
	server.router.Use(cors.Default())
}

// 認証を必要としないルーティング
func (server *Server) publicRouter() {
	public := server.router.Group("/v1")

	public.GET("/ping", server.Ping)
	public.POST("/hackathons", server.CreateHackathon)
	public.GET("/hackathons", server.ListHackathons)
	public.GET("/hackathons/:hackathon_id", server.GetHackathon)
	public.GET("/locates", server.ListLocation)
	public.GET("/tech_tags", server.ListTechTags)
	public.GET("/frameworks", server.ListFrameworks)
}

// 認証ミドルウェアの必要なルーティング
func (server *Server) authRouter() {
	auth := server.router.Group("/v1")
	auth.Use(AuthMiddleware())
	// アカウント
	auth.POST("/accounts", server.CreateAccount)
	auth.GET("/accounts/:id", server.GetAccount)
	// ルーム
	auth.POST("/rooms", server.CreateRoom)
	auth.POST("/rooms/:room_id", server.AddAccountInRoom)
	auth.GET("/rooms", server.ListRooms)
	auth.GET("/rooms/:room_id", server.GetRoom)
	auth.POST("/rooms/:room_id/chatroom", server.AddChat)
	// TODO ブックマークURL 設計
	auth.POST("/bookmarks", server.CreateBookmark)
	auth.POST("/bookmarks/:hackathon_id", server.RemoveBookmark)
	auth.GET("/bookmarks/", server.ListBookmarkToHackathon)
}
