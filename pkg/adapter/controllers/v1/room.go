package controller

import (
	"database/sql"
	"errors"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hackhack-Geek-vol6/backend/pkg/adapter/gateways/infrastructure/httpserver/middleware"
	"github.com/hackhack-Geek-vol6/backend/pkg/bootstrap"
	"github.com/hackhack-Geek-vol6/backend/pkg/domain/params"
	"github.com/hackhack-Geek-vol6/backend/pkg/domain/request"
	"github.com/hackhack-Geek-vol6/backend/pkg/usecase/inputport"
	"github.com/hackhack-Geek-vol6/backend/pkg/util/jwt"
	"github.com/newrelic/go-agent/v3/integrations/nrgin"
)

type RoomController struct {
	RoomUsecase inputport.RoomUsecase
	Env         *bootstrap.Env
}

// ListRooms	godoc
//
//	@Summary		List Account
//	@Description	List Account
//	@Tags			Rooms
//	@Produce		json
//	@Param			ListRequest	query		domain.ListRequest		true	"List Rooms Request"
//	@Success		200			{array}		domain.ListRoomResponse	"success response"
//	@Failure		400			{object}	ErrorResponse			"error response"
//	@Failure		500			{object}	ErrorResponse			"error response"
//	@Router			/rooms	[get]
func (rc *RoomController) ListRooms(ctx *gin.Context) {
	txn := nrgin.Transaction(ctx)
	defer txn.End()
	var (
		reqURI request.ListRequest
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
//
//	@Summary		Get Room
//	@Description	Get Room
//	@Tags			Rooms
//	@Produce		json
//	@Param			room_id				path		string					true	"Rooms API wildcard"
//	@Success		200					{object}	domain.GetRoomResponse	"success response"
//	@Failure		400					{object}	ErrorResponse			"error response"
//	@Failure		403	{object}	ErrorResponse				"error response"
//	@Failure		500					{object}	ErrorResponse			"error response"
//	@Router			/rooms/{room_id}																																[get]
func (rc *RoomController) GetRoom(ctx *gin.Context) {
	txn := nrgin.Transaction(ctx)
	defer txn.End()
	var request request.RoomsRequestWildCard
	if err := ctx.ShouldBindUri(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	response, err := rc.RoomUsecase.GetRoom(ctx, request.RoomID)
	if err != nil {
		switch err.Error() {
		case sql.ErrNoRows.Error():
			err := errors.New("そんなルームないで")
			ctx.JSON(http.StatusForbidden, errorResponse(err))
		default:
			err := errors.New("すまんサーバエラーや")
			ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		}
		return
	}
	ctx.JSON(http.StatusOK, response)
}

// CreateRoom		godoc
//
//	@Summary		Create Rooms
//	@Description	Create Rooms
//	@Tags			Rooms
//	@Produce		json
//	@Param			CreateRoomRequestBody	body		domain.CreateRoomRequestBody	true	"create Room Request Body"
//	@Success		200						{object}	domain.GetRoomResponse			"success response"
//	@Failure		400						{object}	ErrorResponse					"error response"
//	@Failure		403						{object}	ErrorResponse					"error response"
//	@Failure		500						{object}	ErrorResponse					"error response"
//	@Router			/rooms																																																																						[post]
func (rc *RoomController) CreateRoom(ctx *gin.Context) {
	txn := nrgin.Transaction(ctx)
	defer txn.End()
	var reqBody request.CreateRoomRequestBody
	if err := ctx.ShouldBindJSON(&reqBody); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	response, err := rc.RoomUsecase.CreateRoom(ctx, params.CreateRoomParams{
		Title:       reqBody.Title,
		Description: reqBody.Description,
		HackathonID: reqBody.HackathonID,
		MemberLimit: reqBody.MemberLimit,
		OwnerID:     reqBody.AccountID,
	})
	if err != nil {
		switch err.Error() {
		case sql.ErrNoRows.Error():
			err := errors.New("そんなユーザ/ルームないで")
			ctx.JSON(http.StatusForbidden, errorResponse(err))
		default:
			ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		}
		return
	}

	ctx.JSON(http.StatusOK, response)
}

// UpdateRoom	godoc
//
//	@Summary		update Room
//	@Description	update Room
//	@Tags			Rooms
//	@Produce		json
//	@Param			room_id					path		string							true	"Rooms API wildcard"
//	@Param			UpdateRoomRequestBody	body		domain.UpdateRoomRequestBody	true	"update Room Request body"
//	@Success		200						{object}	domain.GetRoomResponse			"success response"
//	@Failure		400						{object}	ErrorResponse					"error response"
//	@Failure		403						{object}	ErrorResponse					"error response"
//	@Failure		500						{object}	ErrorResponse					"error response"
//	@Router			/rooms/{room_id}	[put]
func (rc *RoomController) UpdateRoom(ctx *gin.Context) {
	txn := nrgin.Transaction(ctx)
	defer txn.End()
	var (
		reqURI  request.RoomsRequestWildCard
		reqBody request.UpdateRoomRequestBody
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

	response, err := rc.RoomUsecase.UpdateRoom(ctx, params.UpdateRoomParams{
		RoomID:      reqURI.RoomID,
		Title:       reqBody.Title,
		Description: reqBody.Description,
		HackathonID: reqBody.HackathonID,
		MemberLimit: reqBody.MemberLimit,
		OwnerEmail:  payload.Email,
		IsClosing:   reqBody.IsClosing,
	})
	if err != nil {
		switch err.Error() {
		case sql.ErrNoRows.Error():
			err := errors.New("そんなユーザ/ルームないで")
			ctx.JSON(http.StatusForbidden, errorResponse(err))
		default:
			err := errors.New("すまんサーバエラーや")
			ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		}
		return
	}
	ctx.JSON(http.StatusOK, response)
}

// DeleteRoom	godoc
//
//	@Summary		delete Room
//	@Description	delete Room
//	@Tags			Rooms
//	@Produce		json
//	@Param			room_id				path		string			true	"Rooms API wildcard"
//	@Success		200					{object}	SuccessResponse	"success response"
//	@Failure		400					{object}	ErrorResponse	"error response"
//	@Failure		403					{object}	ErrorResponse	"error response"
//	@Failure		500					{object}	ErrorResponse	"error response"
//	@Router			/rooms/{room_id}	[delete]
func (rc *RoomController) DeleteRoom(ctx *gin.Context) {
	txn := nrgin.Transaction(ctx)
	defer txn.End()
	var (
		reqURI request.RoomsRequestWildCard
	)
	if err := ctx.ShouldBindUri(&reqURI); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	payload := ctx.MustGet(middleware.AuthorizationClaimsKey).(*jwt.FireBaseCustomToken)

	if err := rc.RoomUsecase.DeleteRoom(ctx, params.DeleteRoomParams{
		OwnerEmail: payload.Email,
		RoomID:     reqURI.RoomID,
	}); err != nil {
		switch err.Error() {
		case sql.ErrNoRows.Error():
			err := errors.New("そんなユーザ/ルームないで")
			ctx.JSON(http.StatusForbidden, errorResponse(err))
		default:
			err := errors.New("すまんサーバエラーや")
			ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		}
		return
	}
	ctx.JSON(http.StatusOK, SuccessResponse{Result: "Delete Successful"})
}

// AddAccountInRoom	godoc
//
//	@Summary		Add Account In Rooms
//	@Description	Add Account In Rooms
//	@Tags			Rooms
//	@Produce		json
//	@Param			room_id						path		string								true	"Rooms API wildcard"
//	@Param			AddAccountInRoomRequestBody	body		domain.AddAccountInRoomRequestBody	true	"add account in room Request body"
//	@Success		200							{object}	SuccessResponse						"success response"
//	@Failure		400							{object}	ErrorResponse						"error response"
//	@Failure		403							{object}	ErrorResponse						"error response"
//	@Failure		500							{object}	ErrorResponse						"error response"
//	@Router			/rooms/{room_id}			[post]
func (rc *RoomController) AddAccountInRoom(ctx *gin.Context) {
	txn := nrgin.Transaction(ctx)
	defer txn.End()
	var (
		reqURI  request.RoomsRequestWildCard
		reqBody request.AddAccountInRoomRequestBody
	)

	if err := ctx.ShouldBindUri(&reqURI); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	if err := ctx.ShouldBindJSON(&reqBody); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	if err := rc.RoomUsecase.AddAccountInRoom(ctx, params.AddAccountInRoomParams{
		AccountID: reqBody.AccountID,
		RoomID:    reqURI.RoomID,
	}); err != nil {
		switch err.Error() {
		case sql.ErrNoRows.Error():
			err := errors.New("そんなユーザ/ルームないで")
			ctx.JSON(http.StatusForbidden, errorResponse(err))
		default:
			err := errors.New("すまんサーバエラーや")
			ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		}
		return
	}
	ctx.JSON(http.StatusOK, SuccessResponse{Result: "Join Successful"})
}

// TODO:ルームからメンバーを削除するユースケース
// RemoveAccountInRoom	godoc
//
//	@Summary		Remove Account In Rooms
//	@Description	Remove Account In Rooms
//	@Tags			Rooms
//	@Produce		json
//	@Param			room_id						path		string			true	"Rooms API wildcard"
//	@Param			RemoveAccountInRoom			query		domain.RemoveAccountInRoomRequest		true	"Remove Account In Room Request"
//	@Success		200							{object}	SuccessResponse	"success response"
//	@Failure		400							{object}	ErrorResponse	"error response"
//	@Failure		403							{object}	ErrorResponse	"error response"
//	@Failure		500							{object}	ErrorResponse	"error response"
//	@Router			/rooms/{room_id}/members	[delete]
func (rc *RoomController) RemoveAccountInRoom(ctx *gin.Context) {
	txn := nrgin.Transaction(ctx)
	defer txn.End()
	var (
		reqURI   request.RoomsRequestWildCard
		reqQuery request.RemoveAccountInRoomRequest
	)

	if err := ctx.ShouldBindUri(&reqURI); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	if err := ctx.ShouldBindQuery(&reqQuery); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	payload := ctx.MustGet(middleware.AuthorizationClaimsKey).(*jwt.FireBaseCustomToken)

	if err := rc.RoomUsecase.DeleteRoomAccount(ctx, params.DeleteRoomAccount{
		RoomID:    reqURI.RoomID,
		Email:     payload.Email,
		AccountID: reqQuery.AccountID,
	}); err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, SuccessResponse{Result: "delete Successful"})
}

// AddChat	godoc
//
//	@Summary		Add Chat Room
//	@Description	Add Chat Room
//	@Tags			Rooms
//	@Produce		json
//	@Param			room_id						path		string						true	"Rooms API wildcard"
//	@Param			AddChatRequestBody			body		domain.AddChatRequestBody	true	"add chat Room Request body"
//	@Success		200							{object}	SuccessResponse				"success response"
//	@Failure		400							{object}	ErrorResponse				"error response"
//	@Failure		403							{object}	ErrorResponse				"error response"
//	@Failure		500							{object}	ErrorResponse				"error response"
//	@Router			/rooms/{room_id}/addchat	[post]
func (rc *RoomController) AddChat(ctx *gin.Context) {
	txn := nrgin.Transaction(ctx)
	defer txn.End()
	var (
		reqtURI request.RoomsRequestWildCard
		reqBody request.AddChatRequestBody
	)

	if err := ctx.ShouldBindUri(&reqtURI); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	if err := ctx.ShouldBindJSON(&reqBody); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	if err := rc.RoomUsecase.AddChat(ctx, params.AddChatParams{
		RoomID:    reqtURI.RoomID,
		AccountID: reqBody.AccountID,
		Message:   reqBody.Message,
	}); err != nil {
		switch err.Error() {
		case sql.ErrNoRows.Error():
			err := errors.New("そんなユーザおらんがな")
			ctx.JSON(http.StatusForbidden, errorResponse(err))
		default:
			ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		}
		return
	}

	ctx.JSON(http.StatusOK, SuccessResponse{Result: "inserted successfully"})
}

// AddChat	godoc
//
//	@Summary		CloseRoom
//	@Description	CloseRoom
//	@Tags			Rooms
//	@Produce		json
//	@Param			room_id						path		string						true	"Rooms API wildcard"
//	@Param			CloseRoomRequest			body		domain.CloseRoomRequest		true	"Close Room Request body"
//	@Success		200							{object}	SuccessResponse				"success response"
//	@Failure		400							{object}	ErrorResponse				"error response"
//	@Failure		403							{object}	ErrorResponse				"error response"
//	@Failure		500							{object}	ErrorResponse				"error response"
//	@Router			/rooms/{room_id}/members	[post]
func (rc *RoomController) CloseRoom(ctx *gin.Context) {
	txn := nrgin.Transaction(ctx)
	defer txn.End()
	var (
		reqtURI request.RoomsRequestWildCard
		reqBody request.CloseRoomRequest
	)

	if err := ctx.ShouldBindUri(&reqtURI); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	if err := ctx.ShouldBindJSON(&reqBody); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	if err := rc.RoomUsecase.CloseRoom(ctx, params.CloseRoomParams{
		RoomID:    reqtURI.RoomID,
		AccountID: reqBody.AccountID,
	}); err != nil {
		switch err.Error() {
		case sql.ErrNoRows.Error():
			err := errors.New("そんなユーザおらんがな")
			ctx.JSON(http.StatusForbidden, errorResponse(err))
		default:
			ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		}
		return
	}

	ctx.JSON(http.StatusOK, SuccessResponse{Result: "successfully"})
}
