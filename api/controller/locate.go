package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// ListLocation	godoc
// @Summary			Get Frameworks
// @Description		Get Frameworks
// @Tags			Locates
// @Produce			json
// @Success			200			{array}		db.Locates		"success response"
// @Failure 		500			{object}	ErrorResponse	"error response"
// @Router       	/locates	[get]
func (server *Server) ListLocation(ctx *gin.Context) {
	locate, err := server.store.ListLocates(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, locate)
}