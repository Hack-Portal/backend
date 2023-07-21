package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// ListLocation	godoc
// @Summary			Get Framewroks
// @Description		Get Framewroks
// @Tags			Locates
// @Produce			json
// @Success			200			{array}			db.Locates			"succsss response"
// @Failure 		500			{object}		ErrorResponse		"error response"
// @Router       	/locates	[get]
func (server *Server) ListLocation(ctx *gin.Context) {
	locate, err := server.store.ListLocates(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, locate)
}
