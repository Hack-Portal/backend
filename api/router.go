package api

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/hackhack-Geek-vol6/backend/docs"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func (server *Server) setupRouter() {
	router := gin.Default()
	server.router = router
	server.setUpCors()
	server.publicRouter()
	server.authRouter()
	server.setupSwagger()

}

func (server *Server) setupSwagger() {
	server.router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
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

	public.POST("/ping", server.Ping)
	public.GET("/locates", server.ListLocation)
	public.GET("/tech_tags", server.ListTechTags)
	public.GET("/frameworks", server.ListFrameworks)

	public.POST("/hackathons", server.CreateHackathon)
	public.GET("/hackathons", server.ListHackathons)
	public.GET("/hackathons/:hackathon_id", server.GetHackathon)
}

// 認証ミドルウェアの必要なルーティング
func (server *Server) authRouter() {
	auth := server.router.Group("/v1")
	auth.Use(AuthMiddleware())
	// アカウント
	auth.POST("/accounts", server.CreateAccount)

	auth.GET("/accounts/:id", server.GetAccount)
	auth.PUT("/accounts/:id", server.UpdateAccount)
	auth.DELETE("/acccounts/:id", server.DeleteAccount)

	// auth.GET("/acccounts/:id/follow")
	auth.POST("/acccounts/:id/follow", server.CreateFollow)
	auth.DELETE("/acccounts/:id/follow", server.RemoveFollow)

	// レート
	auth.POST("/accounts/:id/rate", server.CreateRate)
	auth.GET("/accounts/:id/rate", server.ListRate)

	// ルーム
	auth.GET("/rooms", server.ListRooms)
	auth.POST("/rooms", server.CreateRoom)

	auth.GET("/rooms/:room_id", server.GetRoom)
	auth.POST("/rooms/:room_id", server.AddAccountInRoom)
	auth.PUT("/rooms/:room_id", server.UpdateRoom)
	auth.DELETE("/rooms/:room_id", server.DeleteRoom)

	auth.POST("/rooms/:room_id/members", server.AddAccountInRoom)
	auth.DELETE("/rooms/:room_id/members/user_id", server.RemoveAccountInRoom)
	auth.POST("/rooms/:room_id/addchat", server.AddChat)

	// TODO ブックマークURL 設計
	auth.POST("/bookmarks", server.CreateBookmark)
	auth.POST("/bookmarks/:hackathon_id", server.RemoveBookmark)
	auth.GET("/bookmarks/", server.ListBookmarkToHackathon)
}
