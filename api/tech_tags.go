package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (server *Server) ListTechTags(ctx *gin.Context) {
	techTags, err := server.store.ListTechTag(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, techTags)
}
