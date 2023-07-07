package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hackhack-Geek-vol6/backend/middleware"
	"github.com/hackhack-Geek-vol6/backend/util/token"
)

func (controller *Controller) Ping(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"message": "pong"})
}

func (controller *Controller) Pong(ctx *gin.Context) {
	authClaims := ctx.MustGet(middleware.AuthorizationClaimsKey).(*token.CustomClaims)
	ctx.JSON(http.StatusOK, gin.H{"message": "ping", "claims": authClaims})
}
