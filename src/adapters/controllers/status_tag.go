package controllers

import (
	"github.com/hackhack-Geek-vol6/backend/src/datastructure/request"
	"github.com/hackhack-Geek-vol6/backend/src/usecases/ports"
	"github.com/labstack/echo/v4"
)

type StatusTagController struct {
	inputPort ports.StatusTagInputBoundary
}

func NewStatusTagController(inputPort ports.StatusTagInputBoundary) *StatusTagController {
	return &StatusTagController{
		inputPort: inputPort,
	}
}

func (stc *StatusTagController) CreateStatusTag(ctx echo.Context) error {
	var req request.CreateStatusTag
	if err := ctx.Bind(&req); err != nil {
		return echo.ErrBadRequest
	}

	return ctx.JSON(stc.inputPort.CreateStatusTag(ctx.Request().Context(), &req))
}

func (stc *StatusTagController) FindAllStatusTag(ctx echo.Context) error {
	return ctx.JSON(stc.inputPort.FindAllStatusTag(ctx.Request().Context()))
}

func (stc *StatusTagController) FindByIdStatusTag(ctx echo.Context) error {
	var req request.GetStatusTagByID
	if err := ctx.Bind(&req); err != nil {
		return echo.ErrBadRequest
	}

	return ctx.JSON(stc.inputPort.FindByIdStatusTag(ctx.Request().Context(), &req))
}

func (stc *StatusTagController) UpdateStatusTag(ctx echo.Context) error {
	var req request.UpdateStatusTag
	if err := ctx.Bind(&req); err != nil {
		return echo.ErrBadRequest
	}

	return ctx.JSON(stc.inputPort.UpdateStatusTag(ctx.Request().Context(), &req))
}
