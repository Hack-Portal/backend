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
	var request CreateRateRequestBody
	if err := ctx.ShouldBindJSON(&request); err != nil {
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
	ctx.JSON(http.StatusOK, response)
}
