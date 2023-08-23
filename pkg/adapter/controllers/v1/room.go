package controller

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hackhack-Geek-vol6/backend/pkg/adapter/gateways/infrastructure/httpserver/middleware"
	"github.com/hackhack-Geek-vol6/backend/pkg/bootstrap"
	"github.com/hackhack-Geek-vol6/backend/pkg/domain"
	"github.com/hackhack-Geek-vol6/backend/pkg/usecase/inputport"
	"github.com/hackhack-Geek-vol6/backend/pkg/util/jwt"
)

type RoomController struct {
	RoomUsecase inputport.RoomUsecase
	Env         *bootstrap.Env
}

// ListRooms	godoc
//	@Summary		List Account
//	@Description	List Account
//	@Tags			Rooms
//	@Produce		json
//	@Param			ListRequest	query		domain.ListRequest		true	"List Rooms Request"
//	@Success		200			{array}		domain.ListRoomResponse	"success response"
//	@Failure		400			{object}	ErrorResponse			"error response"
//	@Failure		500			{object}	ErrorResponse			"error response"
//	@Router			/rooms													[get]
func (rc *RoomController) ListRooms(ctx *gin.Context) {
	var (
		reqURI domain.ListRequest
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
//	@Summary		Get Room
//	@Description	Get Room
//	@Tags			Rooms
//	@Produce		json
//	@Param			room_id				path		string					true	"Rooms API wildcard"
//	@Success		200					{object}	domain.GetRoomResponse	"success response"
//	@Failure		400					{object}	ErrorResponse			"error response"
//	@Failure		500					{object}	ErrorResponse			"error response"
//	@Router			/rooms/{room_id}																		[get]
func (rc *RoomController) GetRoom(ctx *gin.Context) {
	var request domain.RoomsRequestWildCard
	if err := ctx.ShouldBindUri(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	response, err := rc.RoomUsecase.GetRoom(ctx, request.RoomID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, response)
}

// CreateRoom		godoc
//	@Summary		Create Rooms
//	@Description	Create Rooms
//	@Tags			Rooms
//	@Produce		json
//	@Param			CreateRoomRequestBody	body		domain.CreateRoomRequestBody	true	"create Room Request Body"
//	@Success		200						{object}	domain.GetRoomResponse			"success response"
//	@Failure		400						{object}	ErrorResponse					"error response"
//	@Failure		500						{object}	ErrorResponse					"error response"
//	@Router			/rooms																																												[post]
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
		OwnerID:     reqBody.AccountID,
	})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, response)
}

// UpdateRoom	godoc
//	@Summary		update Room
//	@Description	update Room
//	@Tags			Rooms
//	@Produce		json
//	@Param			room_id					path		string							true	"Rooms API wildcard"
//	@Param			UpdateRoomRequestBody	body		domain.UpdateRoomRequestBody	true	"update Room Request body"
//	@Success		200						{object}	domain.GetRoomResponse			"success response"
//	@Failure		400						{object}	ErrorResponse					"error response"
//	@Failure		500						{object}	ErrorResponse					"error response"
//	@Router			/rooms/{room_id}																																	[put]
func (rc *RoomController) UpdateRoom(ctx *gin.Context) {
	var (
		reqURI  domain.RoomsRequestWildCard
		reqBody domain.UpdateRoomRequestBody
	)

	if err := ctx.ShouldBindUri(&reqURI); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	if err := ctx.ShouldBindJSON(&reqBody); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	payload := ctx.MustGet(middleware.AuthorizationClaimsKey).(*jwt.FireBaseCustomToken)
	fmt.Println(reqBody)

	response, err := rc.RoomUsecase.UpdateRoom(ctx, domain.UpdateRoomParam{
		RoomID:      reqURI.RoomID,
		Title:       reqBody.Title,
		Description: reqBody.Description,
		HackathonID: reqBody.HackathonID,
		MemberLimit: reqBody.MemberLimit,
		OwnerEmail:  payload.Email,
	})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, response)
}

// DeleteRoom	godoc
//	@Summary		delete Room
//	@Description	delete Room
//	@Tags			Rooms
//	@Produce		json
//	@Param			room_id				path		string			true	"Rooms API wildcard"
//	@Success		200					{object}	SuccessResponse	"success response"
//	@Failure		400					{object}	ErrorResponse	"error response"
//	@Failure		500					{object}	ErrorResponse	"error response"
//	@Router			/rooms/{room_id}														[delete]
func (rc *RoomController) DeleteRoom(ctx *gin.Context) {
	var (
		reqURI domain.RoomsRequestWildCard
	)
	if err := ctx.ShouldBindUri(&reqURI); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	payload := ctx.MustGet(middleware.AuthorizationClaimsKey).(*jwt.FireBaseCustomToken)

	if err := rc.RoomUsecase.DeleteRoom(ctx, domain.DeleteRoomParam{
		OwnerEmail: payload.Email,
		RoomID:     reqURI.RoomID,
	}); err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, SuccessResponse{Result: "Delete Successful"})
}

// AddAccountInRoom	godoc
//	@Summary		Add Account In Rooms
//	@Description	Add Account In Rooms
//	@Tags			Rooms
//	@Produce		json
//	@Param			room_id						path		string			true	"Rooms API wildcard"
//	@Success		200							{object}	SuccessResponse	"success response"
//	@Failure		400							{object}	ErrorResponse	"error response"
//	@Failure		500							{object}	ErrorResponse	"error response"
//	@Router			/rooms/{room_id}/members	[post]
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

	if err := rc.RoomUsecase.AddAccountInRoom(ctx, domain.AddAccountInRoomParam{
		AccountID: reqBody.AccountID,
		RoomID:    reqURI.RoomID,
	}); err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, SuccessResponse{Result: "Join Successful"})
}

// TODO:ルームからメンバーを削除するユースケース
// RemoveAccountInRoom	godoc
//	@Summary		Remove Account In Rooms
//	@Description	Remove Account In Rooms
//	@Tags			Rooms
//	@Produce		json
//	@Param			room_id						path		string			true	"Rooms API wildcard"
//	@Success		200							{object}	SuccessResponse	"success response"
//	@Failure		400							{object}	ErrorResponse	"error response"
//	@Failure		500							{object}	ErrorResponse	"error response"
//	@Router			/rooms/{room_id}/members	[delete]
func (rc *RoomController) RemoveAccountInRoom(ctx *gin.Context) {
	var (
		reqURI domain.RoomsRequestWildCard
	)

	if err := ctx.ShouldBindUri(&reqURI); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	payload := ctx.MustGet(middleware.AuthorizationClaimsKey).(*jwt.FireBaseCustomToken)

	if err := rc.RoomUsecase.DeleteRoomAccount(ctx, domain.DeleteRoomAccount{
		RoomID: reqURI.RoomID,
		Email:  payload.Email,
	}); err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, SuccessResponse{Result: "delete Successful"})
}

// AddChat	godoc
//	@Summary		Add Chat Room
//	@Description	Add Chat Room
//	@Tags			Rooms
//	@Produce		json
//	@Param			room_id						path		string						true	"Rooms API wildcard"
//	@Param			AddChatRequestBody			body		domain.AddChatRequestBody	true	"add chat Room Request body"
//	@Success		200							{object}	domain.GetRoomResponse		"success response"
//	@Failure		400							{object}	ErrorResponse				"error response"
//	@Failure		500							{object}	ErrorResponse				"error response"
//	@Router			/rooms/{room_id}/addchat																											[post]
func (rc *RoomController) AddChat(ctx *gin.Context) {
	var (
		reqtURI domain.RoomsRequestWildCard
		reqBody domain.AddChatRequestBody
	)

	if err := ctx.ShouldBindUri(&reqtURI); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	if err := ctx.ShouldBindJSON(&reqBody); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	if err := rc.RoomUsecase.AddChat(ctx, domain.AddChatParams{
		RoomID:    reqtURI.RoomID,
		AccountID: reqBody.AccountID,
		Message:   reqBody.Message,
	}); err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"success": "inserted successfully"})
}
