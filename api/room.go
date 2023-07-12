package api

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	uuid5 "github.com/gofrs/uuid/v5"
	"github.com/google/uuid"
	db "github.com/hackhack-Geek-vol6/backend/db/sqlc"
	"github.com/hackhack-Geek-vol6/backend/util/token"
)

type CreateRoomRequest struct {
	HackathonID int32  `json:"hackathon_id" binding:"required"`
	Title       string `json:"title" binding:"required"`
	Description string `json:"description" binding:"required"`
	MemberLimit int32  `json:"member_limit" binding:"required"`
	UserID      string `json:"user_id" binding:"required"`
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
	account, err := server.store.GetAccountByID(ctx, request.UserID)
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
	_, err = server.store.GetHackathonByID(ctx, request.HackathonID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	roomID, err := uuid.NewRandom()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	log.Println(roomID)
	result, err := server.store.CreateRoomTx(ctx, db.CreateRoomTxParams{
		Rooms: db.Rooms{
			RoomID:      roomID,
			HackathonID: request.HackathonID,
			Title:       request.Title,
			Description: request.Description,
			MemberLimit: request.MemberLimit,
		},
		UserID: request.UserID,
	})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, result)
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

type ListRoomsResponse struct {
	Rooms []GetRoomResponse `json:"get_rooms_response"`
}

// ルームを取得するAPI
// 認証必須
func (server *Server) ListRooms(ctx *gin.Context) {
	var request ListRoomsRequest
	if err := ctx.ShouldBindQuery(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	response, err := server.store.ListRoomTx(ctx, db.ListRoomTxParam{})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, response)
}

type GetRoomRequest struct {
	RoomID string `uri:"room_id"`
}

// ルームの詳細を取得する
// 認証必須

type GetRoomResponse struct {
	Room     db.Rooms
	Accounts []db.GetRoomsAccountsByRoomIDRow
}

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
	room, err := server.store.GetRoomsByID(ctx, uuid.UUID(roomID))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	roomAccounts, err := server.store.GetRoomsAccountsByRoomID(ctx, room.RoomID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	response := GetRoomResponse{
		Room:     room,
		Accounts: roomAccounts,
	}
	ctx.JSON(http.StatusOK, response)
}
