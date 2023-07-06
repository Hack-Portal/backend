package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hackhack-Geek-vol6/backend/middleware"
)

func (controller *Controller) Ping(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"message": "pong"})
}

func (controller *Controller) Pong(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"message": "ping", "claims": middleware.ExtractClaims(ctx)})
}
