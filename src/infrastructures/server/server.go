package server

import (
	"temp/pkg/logger"
	v1router "temp/src/infrastructures/router/v1"

	"github.com/gin-gonic/gin"
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
