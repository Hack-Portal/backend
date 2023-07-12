package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	db "github.com/hackhack-Geek-vol6/backend/db/sqlc"
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
	_, err := server.store.GetAccountByID(ctx, request.ToUserID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	// フォローされる人がいるか
	_, err = server.store.GetAccountByID(ctx, request.FromUserID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	// フォローしていないか
	_, err = server.store.ListFollowByToUserID(ctx, request.ToUserID)
	if err == nil {
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

	// フォローしているか
	_, err := server.store.ListFollowByToUserID(ctx, request.ToUserID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	result, err := server.store.GetAccountByID(ctx, request.ToUserID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	// フォローを外す
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
