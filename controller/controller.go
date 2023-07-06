package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Controller struct{}

func NewController() *Controller {
	return &Controller{}
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}

func (controller *Controller) Ping(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"message": "pong"})
}
