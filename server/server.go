package server

import (
	"github.com/gin-gonic/gin"
	"github.com/hackhack-Geek-vol6/backend/controller"
	db "github.com/hackhack-Geek-vol6/backend/db/sqlc"
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
		config: config,
		store:  store,
	}
	server.controller = controller.NewController()
	server.setupRouter()

	return server, nil
}

func (server *Server) setupRouter() {
	router := gin.Default()

	router.GET("/v1/ping", server.controller.Ping)
	server.router = router
}

// サーバを開始する
func (server *Server) Start(address string) error {
	return server.router.Run(address)
}
