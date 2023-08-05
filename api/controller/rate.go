package controller

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	db "github.com/hackhack-Geek-vol6/backend/db/sqlc"
)

// レートエントリーを作る時のリクエストパラメータ
type CreateRateRequestBody struct {
	Rate int32 `json:"rate"`
}

// レートエントリーに関するレスポンス
type RateResponses struct {
	UserID   string    `json:"user_id"`
	Rate     int32     `json:"rate"`
	CreateAt time.Time `json:"create_at"`
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
func (server *Server) CreateRate(ctx *gin.Context) {
	var (
		reqURI  AccountRequestWildCard
		reqBody CreateRateRequestBody
	)
	if err := ctx.ShouldBindUri(&reqURI); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	if err := ctx.ShouldBindJSON(&reqBody); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	// れーとエントリを追加
	rate, err := server.store.CreateRate(ctx, db.CreateRateParams{
		UserID: reqURI.ID,
		Rate:   reqBody.Rate,
	})

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	// アカウントのレートを更新
	_, err = server.store.UpdateRateByUserID(ctx, db.UpdateRateByUserIDParams{
		UserID: reqURI.ID,
		Rate:   reqBody.Rate,
	})

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, rate)
}

type ListRateParams struct {
	PageSize int32 `form:"page_size"`
	PageId   int32 `form:"page_id"`
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
func (server *Server) ListRate(ctx *gin.Context) {
	var (
		reqURI   AccountRequestWildCard
		reqQuery ListRateParams
	)

	if err := ctx.ShouldBindUri(&reqURI); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	if err := ctx.ShouldBindQuery(&reqQuery); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	rates, err := server.store.ListRate(ctx, db.ListRateParams{
		UserID: reqURI.ID,
		Limit:  reqQuery.PageSize,
		Offset: (reqQuery.PageId - 1) * reqQuery.PageSize,
	})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, rates)
}
