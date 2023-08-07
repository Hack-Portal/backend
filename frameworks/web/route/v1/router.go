package route

import (
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/hackhack-Geek-vol6/backend/bootstrap"
	db "github.com/hackhack-Geek-vol6/backend/db/sqlc"
	_ "github.com/hackhack-Geek-vol6/backend/docs"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func Setup(env *bootstrap.Env, timeout time.Duration, store db.Store, gin *gin.Engine) {
	setupCors(gin)

	publicRouter := gin.Group("/v1")
	// All Public APIs
	NewEtcRouter(env, timeout, store, publicRouter)
	NewHackathonRouter(env, timeout, store, publicRouter)
	//TODO:middlewareの追加
	protectRouter := gin.Group("/v1").Use()
	// All Protect APIs
	NewAccountRouter(env, timeout, store, protectRouter)
	NewRoomRouter(env, timeout, store, protectRouter)
}

func (server *Server) setupSwagger() {
	server.router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}

func setupCors(router *gin.Engine) {
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
	router.Use(cors.Default())
}

// 認証を必要としないルーティング
func (server *Server) publicRouter() {
	public := server.router.Group("/v1")

	public.GET("/ping", server.Ping)
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
