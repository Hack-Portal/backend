package api

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (server *Server) Ping(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"message": pong})
}
