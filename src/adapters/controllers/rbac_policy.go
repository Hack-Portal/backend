package controllers

import (
	"github.com/Hack-Portal/backend/src/datastructure/request"
	"github.com/Hack-Portal/backend/src/usecases/ports"
	"github.com/labstack/echo/v4"
)

type RbacPolicyController struct {
	inputPort ports.RbacPolicyInputBoundary
}

func NewRbacPolicyController(inputPort ports.RbacPolicyInputBoundary) *RbacPolicyController {
	return &RbacPolicyController{
		inputPort: inputPort,
	}
}

func (r *RbacPolicyController) Create(ctx echo.Context) error {
	var req request.CreateRbacPolicy
	if err := ctx.Bind(&req); err != nil {
		return echo.ErrBadRequest
	}

	return ctx.JSON(r.inputPort.CreateRbacPolicy(ctx.Request().Context(), &req))
}

func (r *RbacPolicyController) ReadAll(ctx echo.Context) error {
	var req request.ListRbacPolicies
	if err := ctx.Bind(&req); err != nil {
		return echo.ErrBadRequest
	}

	return ctx.JSON(r.inputPort.ListRbacPolicies(ctx.Request().Context(), &req))
}

func (r *RbacPolicyController) Delete(ctx echo.Context) error {
	var req request.DeleteRbacPolicy
	if err := ctx.Bind(&req); err != nil {
		return echo.ErrBadRequest
	}

	return ctx.JSON(r.inputPort.DeleteRbacPolicy(ctx.Request().Context(), &req))
}

func (r *RbacPolicyController) DeleteAll(ctx echo.Context) error {
	return ctx.JSON(r.inputPort.DeleteAllRbacPolicies(ctx.Request().Context()))
}
