package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (server *Server) ListFrameworks(ctx *gin.Context) {
	frameworks, err := server.store.ListFrameworks(ctx, int32(10000))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, frameworks)
}
