package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hackhack-Geek-vol6/backend/util/token"
)

type CreateRoomResponse struct {
	HackathonID    int32   `json:"hackathon_id" binding:"required"`
	Title          string  `json:"title" binding:"required"`
	Description    string  `json:"description" binding:"required"`
	MemberLimit    int32   `json:"member_limit" binding:"required"`
	UserID         string  `json:"user_id" binding:"required"`
	RoomTechTags   []int32 `json:"room_tech_tags"`
	RoomFrameworks []int32 `json:"room_frameworks"`
}

// ルームを作るAPI　POST:
// 認証必須

func (server *Server) CreateRoom(ctx *gin.Context) {
	var request CreateAccountRequestParam
	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	payload := ctx.MustGet(AuthorizationClaimsKey).(*token.FireBaseCustomToken)
	account, err := server.store.GetAccount(ctx, request.UserID)
	if err != nil {
		// Userがいない時の処理も必要
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	// AuthとUIDが違う時
	if payload.Email != account.Email {
		ctx.JSON(http.StatusBadGateway, errorResponse(err))
		return
	}

}

// ルームにアカウントを追加するAPI
// 認証必須
func (server *Server) AddAccountInRoom(ctx *gin.Context) {

}

// ルームからアカウントを削除するAPI
// 認証必須
func (server *Server) RemoveAccountInRoom(ctx *gin.Context) {

}

// ルームを取得するAPI
// 認証必須
func (server *Server) ListRooms(ctx *gin.Context) {

}

// ルームの詳細を取得する
// 認証必須
func (server *Server) GetRoom(ctx *gin.Context) {

}
