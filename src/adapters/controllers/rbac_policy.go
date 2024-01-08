package controllers

import (
	"github.com/Hack-Portal/backend/src/datastructure/request"
	_ "github.com/Hack-Portal/backend/src/datastructure/response"
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

// RBACPolicy		godoc
//
// @Summary			Create RBACPolicy
// @Description	Create RBACPolicy
// @Tags				RBACPolicy
// @Produce			json
// @Param				CreatePolicy 						body			request.CreateRbacPolicy	true		"request body"
// @Success			200											{object}	response.CreateRbacPolicy					"success response"
// @Failure			400											{object}	nil																"error response"
// @Failure			500											{object}	nil																"error response"
// @Router			/rbac										[POST]
func (r *RbacPolicyController) Create(ctx echo.Context) error {
	var req request.CreateRbacPolicy
	if err := ctx.Bind(&req); err != nil {
		return echo.ErrBadRequest
	}

	return ctx.JSON(r.inputPort.CreateRbacPolicy(ctx.Request().Context(), &req))
}

// RBACPolicy		godoc
//
// @Summary			List Policies
// @Description	List Policies
// @Tags				RBACPolicy
// @Produce			json
// @Param				ListPolicies						query			request.ListRbacPolicies		true	"request query"
// @Success			200											{object}	response.ListRbacPolicies				"success response"
// @Failure			400											{object}	nil																"error response"
// @Failure			500											{object}	nil																"error response"
// @Router			/rbac										[GET]
func (r *RbacPolicyController) ReadAll(ctx echo.Context) error {
	var req request.ListRbacPolicies
	if err := ctx.Bind(&req); err != nil {
		return echo.ErrBadRequest
	}

	return ctx.JSON(r.inputPort.ListRbacPolicies(ctx.Request().Context(), &req))
}

// RBACPolicy		godoc
//
// @Summary			Delete Policies
// @Description	Delete Policies
// @Tags				RBACPolicy
// @Produce			json
// @Param				policy_id								path			string											true	"request query"
// @Success			200											{object}	response.DeleteRbacPolicy					"success response"
// @Failure			400											{object}	nil																"error response"
// @Failure			500											{object}	nil																"error response"
// @Router			/rbac/{policy_id}				[DELETE]
func (r *RbacPolicyController) Delete(ctx echo.Context) error {
	var req request.DeleteRbacPolicy
	if err := ctx.Bind(&req); err != nil {
		return echo.ErrBadRequest
	}

	return ctx.JSON(r.inputPort.DeleteRbacPolicy(ctx.Request().Context(), &req))
}

// RBACPolicy		godoc
//
// @Summary			DeleteAll Policies
// @Description	DeleteAll Policies
// @Tags				RBACPolicy
// @Produce			json
// @Success			200											{object}	response.DeleteAllRbacPolicies		"success response"
// @Failure			400											{object}	nil																"error response"
// @Failure			500											{object}	nil																"error response"
// @Router			/rbac										[DELETE]
func (r *RbacPolicyController) DeleteAll(ctx echo.Context) error {
	return ctx.JSON(r.inputPort.DeleteAllRbacPolicies(ctx.Request().Context()))
}
