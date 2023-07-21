package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	db "github.com/hackhack-Geek-vol6/backend/db/sqlc"
	"github.com/hackhack-Geek-vol6/backend/util/token"
)

type FollowRequestURI struct {
	FromUserID string `uri:"from_user_id"`
}
type CreateFollowRequestBody struct {
	ToUserID string `json:"to_user_id" binding:"required"`
}

// CreateFollow	godoc
// @Summary			Create Follow
// @Description		Create Follow
// @Tags			AccountsFollow
// @Produce			json
// @Param			from_user_id 				path 		string						true	"create Follow Request path"
// @Param			CreateFollowRequestBody 	body 		CreateFollowRequestBody		true	"create Follow Request Body"
// @Success			200			{array}			db.Follows		"succsss response"
// @Failure 		400			{object}		ErrorResponse	"error response"
// @Failure 		500			{object}		ErrorResponse	"error response"
// @Router       	/acccounts/{from_user_id}/follow	[post]
func (server *Server) CreateFollow(ctx *gin.Context) {
	var (
		reqURI  FollowRequestURI
		reqBody CreateFollowRequestBody
	)
	if err := ctx.ShouldBindUri(&reqURI); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	if err := ctx.ShouldBindJSON(&reqBody); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	// フォローする人がいるか
	// 本人確認
	payload := ctx.MustGet(AuthorizationClaimsKey).(*token.FireBaseCustomToken)
	account, err := server.store.GetAccountByEmail(ctx, payload.Email)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	if account.UserID != reqBody.ToUserID {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
	}

	// フォローされる人がいるか
	followedAccounts, err := server.store.ListFollowByToUserID(ctx, reqBody.ToUserID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	if checkFollow(followedAccounts, reqBody.ToUserID) {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	// フォローする
	result, err := server.store.CreateFollow(ctx, db.CreateFollowParams{
		ToUserID:   reqBody.ToUserID,
		FromUserID: reqURI.FromUserID,
	})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, result)
}

// フォローしていないか
func checkFollow(accounts []db.Follows, userID string) bool {
	for _, account := range accounts {
		if account.ToUserID == userID {
			return true
		}
	}
	return false
}

type RemoveFollowRequestQueries struct {
	ToUserID   string `json:"to_user_id" binding:"required"`
	FromUserID string `json:"from_user_id" binding:"required"`
}

// RemoveFollow	godoc
// @Summary			Remove follow
// @Description		Remove follow
// @Tags			AccountsFollow
// @Produce			json
// @Param			from_user_id 				path 		string						true	"remove Follow Request path"
// @Param			RemoveFollowRequestQueries 	body 		CreateFollowRequestBody		true	"remove Follow Request Body"
// @Success			200			{array}			db.Follows		"succsss response"
// @Failure 		400			{object}		ErrorResponse	"error response"
// @Failure 		500			{object}		ErrorResponse	"error response"
// @Router       	/acccounts/{from_user_id}/follow	[delete]
func (server *Server) RemoveFollow(ctx *gin.Context) {
	var (
		reqURI   FollowRequestURI
		reqQuery RemoveFollowRequestQueries
	)
	if err := ctx.ShouldBindUri(&reqURI); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	if err := ctx.ShouldBindQuery(&reqQuery); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	// フォローを外す
	// TODO: エラーハンドリング
	result, err := server.store.GetAccountByID(ctx, reqQuery.ToUserID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	err = server.store.RemoveFollow(ctx, db.RemoveFollowParams{
		ToUserID:   reqQuery.ToUserID,
		FromUserID: reqURI.FromUserID,
	})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, result)
}
