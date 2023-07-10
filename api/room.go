package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	db "github.com/hackhack-Geek-vol6/backend/db/sqlc"
	"github.com/hackhack-Geek-vol6/backend/util/token"
)

type CreateRoomRequest struct {
	HackathonID    int32   `json:"hackathon_id" binding:"required"`
	Title          string  `json:"title" binding:"required"`
	Description    string  `json:"description" binding:"required"`
	MemberLimit    int32   `json:"member_limit" binding:"required"`
	UserID         string  `json:"user_id" binding:"required"`
	RoomTechTags   []int32 `json:"room_tech_tags"`
	RoomFrameworks []int32 `json:"room_frameworks"`
}

type CreateRoomResponse struct {
	RoomID      uuid.UUID                `json:"room_id"`
	HackathonID int32                    `json:"hackathon_id"`
	Title       string                   `json:"title"`
	Description string                   `json:"description"`
	MemberLimit int32                    `json:"member_limit"`
	NowMember   []db.GetRoomsAccountsRow `json:"now_member"`
	TechTags    []db.TechTags            `json:"tech_tags"`
	Frameworks  []db.Frameworks          `json:"frameworks"`
}

// ルームを作るAPI　POST:
// 認証必須

func (server *Server) CreateRoom(ctx *gin.Context) {
	var request CreateRoomRequest
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
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	// ハッカソンがあるか？
	_, err = server.store.GetHackathon(ctx, request.HackathonID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	result, err := server.store.CreateRoomTx(ctx, db.CreateRoomTxParams{
		Rooms: db.Rooms{
			RoomID:      uuid.New(),
			HackathonID: request.HackathonID,
			Title:       request.Title,
			Description: request.Description,
			MemberLimit: request.MemberLimit,
		},
		UserID:          request.UserID,
		RoomsTechTags:   request.RoomTechTags,
		RoomsFrameworks: request.RoomFrameworks,
	})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	roomaccounts, err := server.store.GetRoomsAccounts(ctx, result.RoomID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	response := CreateRoomResponse{
		RoomID:      result.RoomID,
		HackathonID: result.HackathonID,
		Title:       result.Title,
		Description: result.Description,
		MemberLimit: result.MemberLimit,
		NowMember:   roomaccounts,
		TechTags:    result.RoomsTechTags,
		Frameworks:  result.RoomsFrameworks,
	}
	ctx.JSON(http.StatusOK, response)
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
