package api

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	db "github.com/hackhack-Geek-vol6/backend/db/sqlc"
)

// レートエントリーを作る時のリクエストパラメータ
type CreateRateRequestBody struct {
	UserID string `json:"user_id"`
	Rate   int32  `json:"rate"`
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
		request     CreateRateRequestBody
		requestBody UpdateAccountRequestBody
	)
	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	if err := ctx.ShouldBindJSON(&requestBody); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	rate, err := server.store.CreateRate(ctx, db.CreateRateParams{
		UserID: request.UserID,
		Rate:   request.Rate,
	})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	response := RateResponses{
		UserID:   rate.UserID,
		Rate:     rate.Rate,
		CreateAt: rate.CreateAt,
	}
	// アカウントのレートを更新
	account, err := server.store.GetAccountByID(ctx, request.UserID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	upRate, err := server.store.UpdateAccount(ctx, parseUpdateAccountRate(account, requestBody))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, response)
}

func parseUpdateAccountRate(account db.GetAccountByIDRow, body UpdateAccountRequestBody) (result db.UpdateAccountParams) {
	result.UserID = account.UserID
	result.Username = account.Username
	result.Icon = account.Icon
	result.ExplanatoryText = account.ExplanatoryText
	result.LocateID = account.LocateID
	result.Rate = account.Rate + body.Rate
	result.HashedPassword = account.HashedPassword
	result.Email = account.Email
	result.ShowLocate = account.ShowLocate
	result.ShowRate = account.ShowRate
	return
}

type ListRateParams struct {
	UserID   string `form:"user_id"`
	PageSize int32  `form:"page_size"`
	PageId   int32  `form:"page_id"`
}
type ListRateResponses struct {
	UserID   string    `json:"user_id"`
	Rate     int32     `json:"rate"`
	CreateAt time.Time `json:"create_at"`
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
	var request ListRateParams
	if err := ctx.ShouldBindQuery(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	rate, err := server.store.ListRate(ctx, db.ListRateParams{
		UserID: request.UserID,
		Limit:  request.PageSize,
		Offset: (request.PageId - 1) * request.PageSize,
	})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	var response []ListRateResponses
	for _, rate := range rate {
		response = append(response, ListRateResponses{
			UserID:   rate.UserID,
			Rate:     rate.Rate,
			CreateAt: rate.CreateAt,
		})
	}
	ctx.JSON(http.StatusOK, response)
}
