package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// ListFrameworks	godoc
// @Summary			Remove follow
// @Description		Remove follow
// @Tags			AccountsFollow
// @Produce			json
// @Success			200			{array}			db.ListFrameworks	"succsss response"
// @Failure 		500			{object}		ErrorResponse		"error response"
// @Router       	/frameworks	[get]
func (server *Server) ListFrameworks(ctx *gin.Context) {
	frameworks, err := server.store.ListFrameworks(ctx, int32(10000))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, frameworks)
}
