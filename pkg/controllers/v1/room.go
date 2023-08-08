package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	uuid5 "github.com/gofrs/uuid/v5"
	"github.com/google/uuid"
	"github.com/hackhack-Geek-vol6/backend/pkg/bootstrap"
	"github.com/hackhack-Geek-vol6/backend/pkg/domain"
	"github.com/hackhack-Geek-vol6/backend/pkg/gateways/infrastructure/httpserver/middleware"
	"github.com/hackhack-Geek-vol6/backend/pkg/usecase/inputport"
	"github.com/hackhack-Geek-vol6/backend/pkg/util/jwt"
)

type RoomController struct {
	RoomUsecase inputport.RoomUsecase
	Env         *bootstrap.Env
}

// ListRooms	godoc
// @Summary			List Account
// @Description		List Account
// @Tags			Rooms
// @Produce			json
// @Param			room_id path 		string					true	"Rooms API wildcard"
// @Success			200		{array}		[]db.ListRoomTxResult	"success response"
// @Failure 		400		{object}	ErrorResponse			"error response"
// @Failure 		500		{object}	ErrorResponse			"error response"
// @Router       	/rooms	[get]
func (rc *RoomController) ListRooms(ctx *gin.Context) {
	var (
		reqURI domain.ListRoomsRequest
	)
	if err := ctx.ShouldBindQuery(&reqURI); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	response, err := rc.RoomUsecase.ListRooms(ctx, reqURI)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, response)
}

// GetRoom	godoc
// @Summary			Get Room
// @Description		Get Room
// @Tags			Rooms
// @Produce			json
// @Param			room_id path 		string				true	"Rooms API wildcard"
// @Success			200		{object}	GetRoomResponse		"success response"
// @Failure 		400		{object}	ErrorResponse		"error response"
// @Failure 		500		{object}	ErrorResponse		"error response"
// @Router       	/rooms/:room_id		[get]
func (rc *RoomController) GetRoom(ctx *gin.Context) {
	var request domain.RoomsRequestWildCard
	if err := ctx.ShouldBindUri(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	roomID, err := uuid5.FromString(request.RoomID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	response, err := rc.RoomUsecase.GetRoom(ctx, uuid.UUID(roomID))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, response)
}

// CreateRoom		godoc
// @Summary			Create Rooms
// @Description		Create Rooms
// @Tags			Rooms
// @Produce			json
// @Param			CreateRoomRequestBody 	body 		CreateRoomRequestBody	true	"create Room Request Body"
// @Success			200						{object}	db.CreateRoomTxResult	"success response"
// @Failure 		400						{object}	ErrorResponse			"error response"
// @Failure 		500						{object}	ErrorResponse			"error response"
// @Router       	/rooms					[post]
func (rc *RoomController) CreateRoom(ctx *gin.Context) {
	var reqBody domain.CreateRoomRequestBody
	if err := ctx.ShouldBindJSON(&reqBody); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	response, err := rc.RoomUsecase.CreateRoom(ctx, domain.CreateRoomParam{
		Title:       reqBody.Title,
		Description: reqBody.Description,
		HackathonID: reqBody.HackathonID,
		MemberLimit: reqBody.MemberLimit,
		OwnerID:     reqBody.UserID,
	})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, response)
}

// UpdateRoom	godoc
// @Summary			update Room
// @Description		update Room
// @Tags			Rooms
// @Produce			json
// @Param			room_id 				path 		string					true	"Rooms API wildcard"
// @Param			UpdateRoomRequestBody 	body 		UpdateRoomRequestBody	true	"update Room Request body"
// @Success			200						{object}	GetRoomResponse			"success response"
// @Failure 		400						{object}	ErrorResponse			"error response"
// @Failure 		500						{object}	ErrorResponse			"error response"
// @Router       	/rooms/:room_id			[put]
func (rc *RoomController) UpdateRoom(ctx *gin.Context) {
	var (
		reqURI  domain.RoomsRequestWildCard
		reqBody domain.UpdateRoomRequestBody
	)

	if err := ctx.ShouldBindUri(&reqURI); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	if err := ctx.ShouldBindUri(&reqBody); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	roomID, err := uuid5.FromString(reqURI.RoomID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	payload := ctx.MustGet(middleware.AuthorizationClaimsKey).(*jwt.FireBaseCustomToken)

	response, err := rc.RoomUsecase.UpdateRoom(ctx, domain.UpdateRoomParam{
		RoomID:      uuid.UUID(roomID),
		Title:       reqBody.Title,
		Description: reqBody.Description,
		HackathonID: reqBody.HackathonID,
		MemberLimit: reqBody.MemberLimit,
		OwnerEmail:  payload.Email,
	})
	ctx.JSON(http.StatusOK, response)
}

// DeleteRoom	godoc
// @Summary			delete Room
// @Description		delete Room
// @Tags			Rooms
// @Produce			json
// @Param			room_id path 		string			true	"Rooms API wildcard"
// @Success			200		{object}	DeleteResponse	"success response"
// @Failure 		400		{object}	ErrorResponse	"error response"
// @Failure 		500		{object}	ErrorResponse	"error response"
// @Router       	/rooms/:room_id		[delete]
func (rc *RoomController) DeleteRoom(ctx *gin.Context) {
	var (
		reqURI domain.RoomsRequestWildCard
	)
	if err := ctx.ShouldBindUri(&reqURI); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	roomID, err := uuid5.FromString(reqURI.RoomID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	payload := ctx.MustGet(middleware.AuthorizationClaimsKey).(*jwt.FireBaseCustomToken)

	if err := rc.RoomUsecase.DeleteRoom(ctx, domain.DeleteRoomParam{
		OwnerEmail: payload.Email,
		RoomID:     uuid.UUID(roomID),
	}); err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, SuccessResponse{Result: "Delete Successful"})
}

// AddAccountInRoom	godoc
// @Summary			Add Account In Rooms
// @Description		Add Account In Rooms
// @Tags			Rooms
// @Produce			json
// @Param			room_id 	path 		string					true	"Rooms API wildcard"
// @Success			200			{object}	db.CreateRoomTxResult	"success response"
// @Failure 		400			{object}	ErrorResponse			"error response"
// @Failure 		500			{object}	ErrorResponse			"error response"
// @Router       	/rooms/:room_id/members	[post]
func (rc *RoomController) AddAccountInRoom(ctx *gin.Context) {
	var (
		reqURI  domain.RoomsRequestWildCard
		reqBody domain.AddAccountInRoomRequestBody
	)

	if err := ctx.ShouldBindUri(&reqURI); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	if err := ctx.ShouldBindJSON(&reqBody); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	roomID, err := uuid5.FromString(reqURI.RoomID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	if err := rc.RoomUsecase.AddAccountInRoom(ctx, domain.AddAccountInRoomParam{
		UserID: reqBody.UserID,
		RoomID: uuid.UUID(roomID),
	}); err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, SuccessResponse{Result: "Join Successful"})
}

// TODO:ルームからメンバーを削除するユースケース
// RemoveAccountInRoom	godoc
// @Summary			Remove Account In Rooms
// @Description		Remove Account In Rooms
// @Tags			Rooms
// @Produce			json
// @Param			room_id 	path 		string			true	"Rooms API wildcard"
// @Success			200			{object}	DeleteResponse	"success response"
// @Failure 		400			{object}	ErrorResponse	"error response"
// @Failure 		500			{object}	ErrorResponse	"error response"
// @Router       	/rooms/:room_id/members	[delete]
func (rc *RoomController) RemoveAccountInRoom(ctx *gin.Context) {
	var (
		reqURI domain.RoomsRequestWildCard
	)

	if err := ctx.ShouldBindUri(&reqURI); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	roomID, err := uuid5.FromString(reqURI.RoomID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	payload := ctx.MustGet(middleware.AuthorizationClaimsKey).(*jwt.FireBaseCustomToken)

	if err := rc.RoomUsecase.DeleteRoom(ctx, domain.DeleteRoomParam{
		OwnerEmail: payload.Email,
		RoomID:     uuid.UUID(roomID),
	}); err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, SuccessResponse{Result: "delete Successful"})

}

// AddChat	godoc
// @Summary			Add Chat Room
// @Description		Add Chat Room
// @Tags			Rooms
// @Produce			json
// @Param			room_id 			path 		string				true	"Rooms API wildcard"
// @Param			AddChatRequestBody 	body 		AddChatRequestBody	true	"add chat Room Request body"
// @Success			200					{object}	GetRoomResponse		"success response"
// @Failure 		400					{object}	ErrorResponse		"error response"
// @Failure 		500					{object}	ErrorResponse		"error response"
// @Router       	/rooms/:room_id/addchat			[post]
func (rc *RoomController) AddChat(ctx *gin.Context) {
	var reqtURI domain.RoomsRequestWildCard
	var reqBody domain.AddChatRequestBody
	if err := ctx.ShouldBindUri(&reqtURI); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	if err := ctx.ShouldBindJSON(&reqBody); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	roomID, err := uuid5.FromString(reqtURI.RoomID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	if err := rc.RoomUsecase.AddChat(ctx, domain.AddChatParams{
		RoomID:  uuid.UUID(roomID),
		UserID:  reqBody.UserID,
		Message: reqBody.Message,
	}); err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"success": "inserted successfully"})
}