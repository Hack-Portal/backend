package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hackhack-Geek-vol6/backend/util/token"
)

func (server *Server) Ping(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"message": "pong"})
}

func (server *Server) Pong(ctx *gin.Context) {
	authClaims := ctx.MustGet(AuthorizationClaimsKey).(*token.CustomClaims)
	ctx.JSON(http.StatusOK, gin.H{"message": "ping", "claims": authClaims})
}
