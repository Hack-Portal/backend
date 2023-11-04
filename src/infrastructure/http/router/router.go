package router

import (
	firebase "firebase.google.com/go"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type ginRouter struct {
	e   *gin.Engine
	db  *gorm.DB
	app *firebase.App
}

func NewRouter(db *gorm.DB) *gin.Engine {
	r := &ginRouter{
		e:  gin.Default(),
		db: db,
	}
	r.hackathonRouter()

	return r.e
}
