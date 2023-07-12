package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	db "github.com/hackhack-Geek-vol6/backend/db/sqlc"
	"github.com/hackhack-Geek-vol6/backend/util/token"
)

type CreateFollow struct {
	ToUserID   string `json:"to_user_id" binding:"required"`
	FromUserID string `json:"from_user_id" binding:"required"`
}

// フォローするAPI
func (server *Server) CreateFollow(ctx *gin.Context) {
	var request CreateFollow
	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	// フォローする人がいるか
	// 本人確認
	payload := ctx.MustGet(AuthorizationClaimsKey).(*token.FireBaseCustomToken)
	account, err := server.store.GetAccountbyEmail(ctx, payload.Email)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	if account.UserID != request.ToUserID {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
	}

	// フォローされる人がいるか
	followedAccounts, err := server.store.ListFollowByToUserID(ctx, request.ToUserID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	if checkFollow(followedAccounts, request.ToUserID) {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	// フォローする
	result, err := server.store.CreateFollow(ctx, db.CreateFollowParams{
		ToUserID:   request.ToUserID,
		FromUserID: request.FromUserID,
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

// フォローを外すAPI
type RemoveFollow struct {
	ToUserID   string `json:"to_user_id" binding:"required"`
	FromUserID string `json:"from_user_id" binding:"required"`
}

func (server *Server) RemoveFollow(ctx *gin.Context) {
	var request RemoveFollow
	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	// フォローを外す
	// TODO: エラーハンドリング
	result, err := server.store.GetAccountByID(ctx, request.ToUserID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	err = server.store.RemoveFollow(ctx, db.RemoveFollowParams{
		ToUserID:   request.ToUserID,
		FromUserID: request.FromUserID,
	})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, result)
}
