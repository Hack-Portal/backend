package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// ListTechTags		godoc
// @Summary			Get Framewroks
// @Description		Get Framewroks
// @Tags			TechTags
// @Produce			json
// @Success			200			{array}			db.TechTags	"succsss response"
// @Failure 		500			{object}		ErrorResponse		"error response"
// @Router       	/tech_tags	[get]
func (server *Server) ListTechTags(ctx *gin.Context) {
	techTags, err := server.store.ListTechTag(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, techTags)
}
