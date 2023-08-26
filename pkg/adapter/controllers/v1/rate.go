package controller

import (
	"database/sql"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	repository "github.com/hackhack-Geek-vol6/backend/pkg/adapter/gateways/repository/datasource"
	"github.com/hackhack-Geek-vol6/backend/pkg/bootstrap"
	"github.com/hackhack-Geek-vol6/backend/pkg/domain"
	"github.com/hackhack-Geek-vol6/backend/pkg/usecase/inputport"
	"github.com/newrelic/go-agent/v3/integrations/nrgin"
)

type RateController struct {
	RateUsecase inputport.RateUsecase
	Env         *bootstrap.Env
}

// CreateRate	godoc
//
//	@Summary		Create Rate
//	@Description	Create Rate for User
//	@Tags			Rate
//	@Produce		json
//	@Param			CreateRateRequestBody			body		domain.CreateRateRequestBody	true	"Create Rate Request Body"
//	@Success		200								{object}	SuccessResponse					"success response"
//	@Failure		400								{object}	ErrorResponse					"error response"
//	@Failure		500								{object}	ErrorResponse					"error response"
//	@Router			/accounts/{account_id}/rate 																																								[post]
func (rc *RateController) CreateRate(ctx *gin.Context) {
	txn := nrgin.Transaction(ctx)
	defer txn.End()
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

	if err := rc.RateUsecase.CreateRateEntry(ctx, repository.CreateRateEntitiesParams{
		AccountID: reqURI.AccountID,
		Rate:      reqBody.Rate,
	}); err != nil {
		switch err.Error() {
		case sql.ErrNoRows.Error():
			err := errors.New("そんなユーザおらんがな")
			ctx.JSON(http.StatusForbidden, errorResponse(err))
		default:
			err := errors.New("すまんサーバエラーや")
			ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		}
		return
	}

	ctx.JSON(http.StatusOK, SuccessResponse{Result: "Rate Update Successful"})
}

// ListRate	godoc
//
//	@Summary		List Rate
//	@Description	List Rate for User
//	@Tags			Rate
//	@Produce		json
//	@Param			account_id						path		string						true	"Account ID"
//	@Param			ListRequest						query		domain.ListRequest			true	"List Rate Params"
//	@Success		200								{array}		domain.AccountRateResponse	"success response"
//	@Failure		400								{object}	ErrorResponse				"error response"
//	@Failure		403								{object}	ErrorResponse				"error response"
//	@Failure		500								{object}	ErrorResponse				"error response"
//	@Router			/accounts/{account_id}/rate 																																				[get]
func (rc *RateController) ListRate(ctx *gin.Context) {
	txn := nrgin.Transaction(ctx)
	defer txn.End()
	var (
		reqURI   domain.AccountRequestWildCard
		reqQuery domain.ListRequest
	)

	if err := ctx.ShouldBindUri(&reqURI); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	if err := ctx.ShouldBindQuery(&reqQuery); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	response, err := rc.RateUsecase.ListRateEntry(ctx, reqURI.AccountID, reqQuery)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, response)
}

// ListAccountRate	godoc
//
//	@Summary		List Account Rate
//	@Description	List Account Rate
//	@Tags			Rate
//	@Produce		json
//	@Param			ListRequest	query		domain.ListRequest			true	"List Rate Params"
//	@Success		200			{array}		domain.AccountRateResponse	"success response"
//	@Failure		400			{object}	ErrorResponse				"error response"
//	@Failure		500			{object}	ErrorResponse				"error response"
//	@Router			/rate 																																									[get]
func (rc *RateController) ListAccountRate(ctx *gin.Context) {
	txn := nrgin.Transaction(ctx)
	defer txn.End()
	var reqQuery domain.ListRequest
	if err := ctx.ShouldBindQuery(&reqQuery); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	response, err := rc.RateUsecase.ListAccountRate(ctx, reqQuery)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, response)
}
