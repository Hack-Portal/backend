package server

import (
	"github.com/gin-gonic/gin"
	"github.com/hackhack-Geek-vol6/backend/pkg/logger"
	v1router "github.com/hackhack-Geek-vol6/backend/src/infrastructures/router/v1"
	"gorm.io/gorm"
)

type ginServer struct {
	engine *gin.Engine
}

func (g *ginServer) Run() {
	RunWithGracefulStop(g.engine)
}

func NewServer(db *gorm.DB, l logger.Logger) *ginServer {
	e := gin.New()

	// ここでミドルウェアを設定する
	e.Use(
		gin.Recovery(),
	)

	return &ginServer{
		engine: v1router.NewRouter(e, db, l).Router(),
	}
}
