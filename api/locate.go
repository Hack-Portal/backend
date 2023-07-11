package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// 都道府県マスタを全取得
func (server *Server) ListLocation(ctx *gin.Context) {
	locate, err := server.store.ListLocates(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, locate)
}
