package v1

import (
	firebase "firebase.google.com/go"
	"github.com/gin-gonic/gin"
	"github.com/hackhack-Geek-vol6/backend/pkg/logger"
	"gorm.io/gorm"
)

type ginRouter struct {
	e      *gin.Engine
	db     *gorm.DB
	app    *firebase.App
	logger logger.Logger
}

func NewRouter(e *gin.Engine, db *gorm.DB, l logger.Logger) *ginRouter {
	return &ginRouter{
		e:      e,
		db:     db,
		logger: l,
	}
}

func (g *ginRouter) Router() *gin.Engine {
	// ここにルーティングを集約する
	return g.e
}
