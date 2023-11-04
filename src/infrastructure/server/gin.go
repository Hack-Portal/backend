package server

import (
	firebase "firebase.google.com/go"
	"github.com/hackhack-Geek-vol6/backend/src/infrastructure/http/router"
	"gorm.io/gorm"
)

type ginServer struct {
	db  *gorm.DB
	app *firebase.App
}

func NewGinServer(db *gorm.DB, app *firebase.App) *ginServer {
	return &ginServer{db, app}
}

func (g *ginServer) Run() {
	r := router.NewRouter(g.db, g.app)
	RunWithGracefulShutdown(r)
}
