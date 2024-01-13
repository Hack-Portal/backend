package controllers

import (
	"github.com/Hack-Portal/backend/src/datastructure/request"
	"github.com/Hack-Portal/backend/src/usecases/ports"
	"github.com/labstack/echo/v4"
)

type userController struct {
	inputPort ports.UserInputBoundary
}

// UserController はUserControllerのインターフェース
type UserController interface {
	InitAdmin(ctx echo.Context) error
	Login(ctx echo.Context) error
}

// NewUserController はUserControllerを返す
func NewUserController(inputPort ports.UserInputBoundary) UserController {
	return &userController{
		inputPort: inputPort,
	}
}

// User		godoc
//
// @Summary			init admin
// @Description	init admin
// @Tags				User
// @Produce			json
// @Param				InitAdmin								body			request.InitAdmin				true			"request body"
// @Success			200											{object}	response.User											"success response"
// @Failure			400											{object}	nil																"error response"
// @Failure			500											{object}	nil																"error response"
// @Router			/init_admin							[POST]
func (uc *userController) InitAdmin(ctx echo.Context) error {
	var req request.InitAdmin
	if err := ctx.Bind(&req); err != nil {
		return echo.ErrBadRequest
	}

	return ctx.JSON(uc.inputPort.InitAdmin(ctx.Request().Context(), req))
}

func (uc *userController) Login(ctx echo.Context) error {
	var req request.Login
	if err := ctx.Bind(&req); err != nil {
		return echo.ErrBadRequest
	}

	return ctx.JSON(uc.inputPort.Login(ctx.Request().Context(), req))
}
