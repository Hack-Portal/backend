package server

import (
	"github.com/hackhack-Geek-vol6/backend/src/infrastructure/http/router"
	"gorm.io/gorm"
)

type ginServer struct {
	db *gorm.DB
}

func NewGinServer(db *gorm.DB) *ginServer {
	return &ginServer{db}
}

func (g *ginServer) Run() {
	r := router.NewRouter(g.db)
	RunWithGracefulShutdown(r)
}
