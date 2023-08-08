package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	repository "github.com/hackhack-Geek-vol6/backend/pkg/adapter/gateways/repository/datasource"
	"github.com/hackhack-Geek-vol6/backend/pkg/bootstrap"
	"github.com/hackhack-Geek-vol6/backend/pkg/domain"
	"github.com/hackhack-Geek-vol6/backend/pkg/usecase/inputport"
)

type RateController struct {
	RateUsecase inputport.RateUsecase
	Env         *bootstrap.Env
}

// CreateRate	godoc
// @Summary			Create Rate
// @Description		Create Rate for User
// @Tags			Rate
// @Produce			json
// @Param  CreateRateRequestBody body CreateRateRequestBody true "Create Rate Request Body"
// @Success			200				{object}		RateResponses		"success response"
// @Failure 		400				{object}		ErrorResponse		"error response"
// @Failure 		500				{object}		ErrorResponse		"error response"
// @Router       	/accounts/:id/rate 		[post]
func (rc *RateController) CreateRate(ctx *gin.Context) {
	var (
		reqURI  domain.AccountRequestWildCard
		reqBody domain.CreateRateRequestBody
	)
	if err := ctx.ShouldBindUri(&reqURI); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	if err := ctx.ShouldBindJSON(&reqBody); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	response, err := rc.RateUsecase.CreateRateEntry(ctx, repository.CreateRateParams{
		UserID: reqURI.UserID,
		Rate:   reqBody.Rate,
	})

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, response)
}

// ListRate	godoc
// @Summary			List Rate
// @Description		List Rate for User
// @Tags			Rate
// @Produce			json
// @Param  ListRateParams uri ListRateParams true "List Rate Params"
// @Param  ListRateParams query ListRateParams true "List Rate Params"
// @Success			200				{array}			ListRateResponses	"success response"
// @Failure 		400				{object}		ErrorResponse		"error response"
// @Failure 		500				{object}		ErrorResponse		"error response"
// @Router       	/accounts/:id/rate 		[get]
func (rc *RateController) ListRate(ctx *gin.Context) {
	var (
		reqURI   domain.AccountRequestWildCard
		reqQuery domain.ListRateParams
	)

	if err := ctx.ShouldBindUri(&reqURI); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	if err := ctx.ShouldBindQuery(&reqQuery); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	response, err := rc.RateUsecase.ListRateEntry(ctx, reqURI.UserID, reqQuery)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, response)
}
