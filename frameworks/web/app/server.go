package api

import (
	"github.com/gin-gonic/gin"
	db "github.com/hackhack-Geek-vol6/backend/db/sqlc"
	"github.com/hackhack-Geek-vol6/backend/util"
)

type Server struct {
	config util.EnvConfig
	store  db.Store
	router *gin.Engine
}

// 新しいサーバを定義する
func NewServer(config util.EnvConfig, store db.Store) (*Server, error) {
	server := &Server{
		config: config,
		store:  store,
	}
	server.setupRouter()
	return server, nil
}

// サーバを開始する
func (server *Server) Start(address string) error {
	return server.router.Run(address)
}
