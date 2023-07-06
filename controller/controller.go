package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hackhack-Geek-vol6/backend/server"
)

type Controller struct {
	server *server.Server
}

func NewController(server *server.Server) *Controller {
	return &Controller{
		server: server,
	}
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}

func (controller *Controller) ping(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"message": "pong"})
}
