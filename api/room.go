package api

import (
	"errors"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	uuid5 "github.com/gofrs/uuid/v5"
	"github.com/google/uuid"
	db "github.com/hackhack-Geek-vol6/backend/db/sqlc"
	"github.com/hackhack-Geek-vol6/backend/util"
	"github.com/hackhack-Geek-vol6/backend/util/token"
)

type RoomsRequestURI struct {
	RoomID string `uri:"room_id"`
}

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

	// TODO:チャットルームの初期化
	// _, err = server.store.InitChatRoom(ctx, result.Rooms.RoomID.String())
	// if err != nil {
	// 	ctx.JSON(http.StatusInternalServerError, errorResponse(err))
	// 	return
	// }

	ctx.JSON(http.StatusOK, result)
}

// ルームにアカウントを追加するAPI
// 認証必須

func (server *Server) AddAccountInRoom(ctx *gin.Context) {
	var reqURI RoomsRequestURI
	if err := ctx.ShouldBindUri(&reqURI); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	payload := ctx.MustGet(AuthorizationClaimsKey).(*token.FireBaseCustomToken)
	account, err := server.store.GetAccountByEmail(ctx, payload.Email)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	roomId, err := uuid5.FromString(reqURI.RoomID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	roomAccounts, err := server.store.CreateRoomsAccounts(ctx, db.CreateRoomsAccountsParams{
		UserID:  account.UserID,
		RoomID:  uuid.UUID(roomId),
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
type DeleteAccountInRoom struct {
	userId string `form:"user_id"`
}

func (server *Server) RemoveAccountInRoom(ctx *gin.Context) {
	var (
		reqURI   RoomsRequestURI
		reqQuery DeleteAccountInRoom
	)
	if err := ctx.ShouldBindUri(&reqURI); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	if err := ctx.ShouldBindQuery(&reqQuery); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	payload := ctx.MustGet(AuthorizationClaimsKey).(*token.FireBaseCustomToken)
	account, err := server.store.GetAccountByEmail(ctx, payload.Email)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	roomId, err := uuid5.FromString(reqURI.RoomID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	roomAccounts, err := server.store.GetRoomsAccountsByRoomID(ctx, uuid.UUID(roomId))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	if !isOwner(roomAccounts, account) {
		err := errors.New("your not owner")
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	err = server.store.RemoveAccountInRoom(ctx, db.RemoveAccountInRoomParams{
		RoomID: uuid.UUID(roomId),
		UserID: reqQuery.userId,
	})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"success": "delete successfully"})

}

type ListRoomsRequest struct {
	PageSize int32 `form:"page_size"`
}

type ListRoomsResponse struct {
	Rooms []GetRoomResponse `json:"get_rooms_response"`
}

// ルームを取得するAPI
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

type hackathonInfo struct {
	HackathonID int32           `json:"hackathon_id"`
	Name        string          `json:"name"`
	Icon        string          `json:"icon"`
	Description string          `json:"description"`
	Link        string          `json:"link"`
	Expired     time.Time       `json:"expired"`
	StartDate   time.Time       `json:"start_date"`
	Term        int32           `json:"term"`
	Tags        []db.StatusTags `json:"tags"`
}

// ルームの詳細を取得する
type GetRoomResponse struct {
	RoomID            uuid.UUID                        `json:"room_id"`
	Title             string                           `json:"title"`
	Description       string                           `json:"description"`
	MemberLimit       int32                            `json:"member_limit"`
	IsStatus          bool                             `json:"is_status"`
	CreateAt          time.Time                        `json:"create_at"`
	Hackathon         hackathonInfo                    `json:"hackathon"`
	NowMember         []db.GetRoomsAccountsByRoomIDRow `json:"now_member"`
	MembersTechTags   []db.RoomTechTags                `json:"members_tech_tags"`
	MembersFrameworks []db.RoomFramework               `json:"members_frameworks"`
}

func (server *Server) GetRoom(ctx *gin.Context) {
	var request RoomsRequestURI
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
	hackathon, err := server.store.GetHackathonByID(ctx, room.HackathonID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	hacathonTags, err := server.store.GetHackathonStatusTagsByHackathonID(ctx, hackathon.HackathonID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	var (
		statusTags        []db.StatusTags
		membersTechTags   []db.RoomTechTags
		MembersFrameworks []db.RoomFramework
	)

	for _, hackathonTag := range hacathonTags {
		tag, err := server.store.GetStatusTagByStatusID(ctx, hackathonTag.StatusID)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, errorResponse(err))
			return
		}
		statusTags = append(statusTags, tag)
	}

	for _, account := range roomAccounts {
		// タグの追加
		techTags, err := server.store.ListAccountTagsByUserID(ctx, account.UserID.String)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, errorResponse(err))
			return
		}
		for _, techTag := range techTags {
			membersTechTags = db.MargeTechTagArray(membersTechTags, db.TechTags{
				TechTagID: techTag.TechTagID.Int32,
				Language:  techTag.Language.String,
			})
		}
		// FWの追加
		frameworks, err := server.store.ListAccountFrameworksByUserID(ctx, account.UserID.String)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, errorResponse(err))
			return
		}
		for _, framework := range frameworks {
			MembersFrameworks = db.MargeFrameworkArray(MembersFrameworks, db.Frameworks{
				FrameworkID: framework.FrameworkID.Int32,
				TechTagID:   framework.TechTagID.Int32,
				Framework:   framework.Framework.String,
			})
		}
	}

	response := GetRoomResponse{
		RoomID:      room.RoomID,
		Title:       room.Title,
		Description: room.Description,
		MemberLimit: room.MemberLimit,
		IsStatus:    room.IsDelete,
		CreateAt:    room.CreateAt,
		Hackathon: hackathonInfo{
			HackathonID: hackathon.HackathonID,
			Name:        hackathon.Name,
			Icon:        hackathon.Icon.String,
			Description: hackathon.Description,
			Link:        hackathon.Link,
			Expired:     hackathon.Expired,
			StartDate:   hackathon.StartDate,
			Term:        hackathon.Term,
			Tags:        statusTags,
		},
		NowMember:         roomAccounts,
		MembersTechTags:   membersTechTags,
		MembersFrameworks: MembersFrameworks,
	}
	ctx.JSON(http.StatusOK, response)
}

// チャットを追加する
// 認証必須
type AddChatRequestBody struct {
	UserID  string `json:"user_id" binding:"required"`
	Message string `json:"message" binding:"required"`
}

func (server *Server) AddChat(ctx *gin.Context) {
	var requestURI RoomsRequestURI
	var requestBody AddChatRequestBody
	if err := ctx.ShouldBindUri(&requestURI); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	if err := ctx.ShouldBindJSON(&requestBody); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	payload := ctx.MustGet(AuthorizationClaimsKey).(*token.FireBaseCustomToken)
	account, err := server.store.GetAccountByEmail(ctx, payload.Email)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	// 本人確認
	if requestBody.UserID != account.UserID {
		err := errors.New("アカウントが一致しません")
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	// ルームメンバか確認する
	roomId, err := uuid5.FromString(requestURI.RoomID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	roomAccounts, err := server.store.GetRoomsAccountsByRoomID(ctx, uuid.UUID(roomId))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	if !checkAccount(roomAccounts, requestBody.UserID) {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	data, err := server.store.ReadDocsByRoomID(ctx, requestURI.RoomID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	_, err = server.store.WriteFireStore(ctx, db.WriteFireStoreParam{
		RoomID:  requestURI.RoomID,
		Index:   len(data) + 1,
		UID:     account.UserID,
		Message: requestBody.Message,
	})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"success": "inserted successfully"})
}

// ユーザが含まれているかの確認
func checkAccount(accounts []db.GetRoomsAccountsByRoomIDRow, roomId string) bool {
	for _, account := range accounts {
		if account.UserID.String == roomId {
			return true
		}
	}
	return false
}

type UpdateRoomRequestBody struct {
	Title       string `json:"title" binding:"required"`
	Description string `json:"description" binding:"required"`
	MemberLimit int32  `json:"member_limit" binding:"required"`
}

// ルームの更新
func (server *Server) UpdateRoom(ctx *gin.Context) {
	var (
		reqURI  RoomsRequestURI
		reqBody UpdateRoomRequestBody
	)
	if err := ctx.ShouldBindUri(&reqURI); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	if err := ctx.ShouldBindUri(&reqBody); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	payload := ctx.MustGet(AuthorizationClaimsKey).(*token.FireBaseCustomToken)
	account, err := server.store.GetAccountByEmail(ctx, payload.Email)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	roomId, err := uuid5.FromString(reqURI.RoomID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	roomAccounts, err := server.store.GetRoomsAccountsByRoomID(ctx, uuid.UUID(roomId))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	if !isOwner(roomAccounts, account) {
		err := errors.New("your not owner")
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	room, err := server.store.GetRoomsByID(ctx, uuid.UUID(roomId))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	result, err := server.store.UpdateRoomByID(ctx, parseUpdateRoomParam(room, reqBody))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, result)
}

func isOwner(roomAccounts []db.GetRoomsAccountsByRoomIDRow, account db.GetAccountByEmailRow) bool {
	for i := range roomAccounts {
		if roomAccounts[i].UserID.String == account.UserID {
			return roomAccounts[i].IsOwner
		}
	}
	return false
}
func parseUpdateRoomParam(room db.Rooms, reqBody UpdateRoomRequestBody) (result db.UpdateRoomByIDParams) {
	result.RoomID = room.RoomID

	if util.StringLength(reqBody.Title) != 0 {
		if !util.EqualString(room.Title, reqBody.Title) {
			result.Title = reqBody.Title
		} else {
			result.Title = room.Title
		}
	}

	if util.StringLength(reqBody.Description) != 0 {
		if !util.EqualString(room.Description, reqBody.Description) {
			result.Description = reqBody.Description
		} else {
			result.Description = room.Description
		}
	}

	if reqBody.MemberLimit != 0 {
		if !util.Equalint(int(room.MemberLimit), int(reqBody.MemberLimit)) {
			result.MemberLimit = reqBody.MemberLimit
		} else {
			result.MemberLimit = room.MemberLimit
		}
	}
	return
}
func (server *Server) DeleteRoom(ctx *gin.Context) {
	var (
		reqURI RoomsRequestURI
	)
	if err := ctx.ShouldBindUri(&reqURI); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	payload := ctx.MustGet(AuthorizationClaimsKey).(*token.FireBaseCustomToken)
	account, err := server.store.GetAccountByEmail(ctx, payload.Email)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	roomId, err := uuid5.FromString(reqURI.RoomID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	roomAccounts, err := server.store.GetRoomsAccountsByRoomID(ctx, uuid.UUID(roomId))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	if !isOwner(roomAccounts, account) {
		err := errors.New("your not owner")
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	result, err := server.store.SoftDeleteRoomByID(ctx, uuid.UUID(roomId))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, result)
}
