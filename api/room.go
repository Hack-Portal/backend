package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	uuid5 "github.com/gofrs/uuid/v5"
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
type AddAccountInRoomRequest struct {
	RoomID uuid.UUID `uri:"room_id"`
}

func (server *Server) AddAccountInRoom(ctx *gin.Context) {
	var request AddAccountInRoomRequest
	if err := ctx.ShouldBindUri(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	payload := ctx.MustGet(AuthorizationClaimsKey).(*token.FireBaseCustomToken)
	account, err := server.store.GetAccountbyEmail(ctx, payload.Email)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	roomAccounts, err := server.store.CreateRoomsAccounts(ctx, db.CreateRoomsAccountsParams{
		UserID:  account.UserID,
		RoomID:  request.RoomID,
		IsOwner: false,
	})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, roomAccounts)
}

// ルームからアカウントを削除するAPI
// 認証必須
// 時間あれば追加
func (server *Server) RemoveAccountInRoom(ctx *gin.Context) {

}

type ListRoomsRequest struct {
	PageSize int32 `form:"page_size"`
}

// ルームを取得するAPI
// 認証必須
func (server *Server) ListRooms(ctx *gin.Context) {
	var request ListRoomsRequest
	if err := ctx.ShouldBindQuery(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	rooms, err := server.store.ListRoom(ctx, request.PageSize)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, rooms)
}

type GetRoomRequest struct {
	RoomID string `uri:"room_id"`
}

// ルームの詳細を取得する
// 認証必須
func (server *Server) GetRoom(ctx *gin.Context) {
	var request GetRoomRequest
	if err := ctx.ShouldBindUri(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	roomID, err := uuid5.FromString(request.RoomID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	room, err := server.store.GetRoom(ctx, uuid.UUID(roomID))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, room)
}