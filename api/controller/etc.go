package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hackhack-Geek-vol6/backend/bootstrap"
	"github.com/hackhack-Geek-vol6/backend/domain"
)

type EtcController struct {
	Etc domain.EtcUsecase
	Env *bootstrap.Env
}

// ListFrameworks	godoc
// @Summary			Get Frameworks
// @Description		Get Frameworks
// @Tags			Frameworks
// @Produce			json
// @Success			200			{array}		db.Frameworks	"success response"
// @Failure 		500			{object}	ErrorResponse	"error response"
// @Router       	/frameworks	[get]
func (ec *EtcController) ListFrameworks(ctx *gin.Context) {
	response, err := ec.Etc.GetFramework(ctx, 1000)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, response)
}

// ListLocation	godoc
// @Summary			Get Frameworks
// @Description		Get Frameworks
// @Tags			Locates
// @Produce			json
// @Success			200			{array}		db.Locates		"success response"
// @Failure 		500			{object}	ErrorResponse	"error response"
// @Router       	/locates	[get]
func (ec *EtcController) ListLocation(ctx *gin.Context) {
	response, err := ec.Etc.GetLocat(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, response)
}

// ListTechTags		godoc
// @Summary			Get Frameworks
// @Description		Get Frameworks
// @Tags			TechTags
// @Produce			json
// @Success			200		{array}		db.TechTags		"success response"
// @Failure 		500		{object}	ErrorResponse	"error response"
// @Router       	/tech_tags			[get]
func (ec *EtcController) ListTechTags(ctx *gin.Context) {
	response, err := ec.Etc.GetTechTag(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, response)
}
