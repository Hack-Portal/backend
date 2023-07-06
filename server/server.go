package server

import (
	"github.com/gin-gonic/gin"
	"github.com/hackhack-Geek-vol6/backend/controller"
	db "github.com/hackhack-Geek-vol6/backend/db/sqlc"
	"github.com/hackhack-Geek-vol6/backend/middleware"
	"github.com/hackhack-Geek-vol6/backend/util"
)

type Server struct {
	config     util.EnvConfig
	store      db.Store
	router     *gin.Engine
	controller *controller.Controller
}

// 新しいサーバを定義する
func NewServer(config util.EnvConfig, store db.Store) (*Server, error) {
	server := &Server{
		config:     config,
		store:      store,
		controller: controller.NewController(),
	}
	server.setupRouter()
	err := server.authRouter()
	if err != nil {
		return nil, err
	}

	return server, nil
}

// ルーティングをセットアップする
func (server *Server) setupRouter() {
	router := gin.New()

	router.GET("/v1/ping", controller.Ping)

	server.router = router
}

// 認証ミドルウェアの必要なルーティング
func (server *Server) authRouter() error {
	middleware, err := middleware.New("../credentials.json", nil)
	if err != nil {
		return err
	}

	auth := server.router.Group("/auth")
	auth.Use(middleware.MiddlewareFunc())
	auth.GET("/v1/pong", controller.Pong)

	return nil
}

// サーバを開始する
func (server *Server) Start(address string) error {
	return server.router.Run(address)
}
