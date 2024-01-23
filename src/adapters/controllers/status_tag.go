package controllers

import (
	"github.com/Hack-Portal/backend/src/datastructure/request"
	_ "github.com/Hack-Portal/backend/src/datastructure/response"
	"github.com/Hack-Portal/backend/src/usecases/ports"
	"github.com/labstack/echo/v4"
	"github.com/newrelic/go-agent/v3/integrations/nrecho-v4"
)

type StatusTagController struct {
	inputPort ports.StatusTagInputBoundary
}

func NewStatusTagController(inputPort ports.StatusTagInputBoundary) *StatusTagController {
	return &StatusTagController{
		inputPort: inputPort,
	}
}

// StatusTag		godoc
//
// @Summary			Create a new StatusTag
// @Description	Create a new StatusTag
// @Tags				StatusTag
// @Produce			json
// @Param				CreateStatusTagRequest	body			request.CreateStatusTag	true			"request body"
// @Success			200											{object}	response.StatusTag								"success response"
// @Failure			400											{object}	nil																"error response"
// @Failure			500											{object}	nil																"error response"
// @Router			/status_tags						[POST]
func (stc *StatusTagController) CreateStatusTag(ctx echo.Context) error {
	defer nrecho.FromContext(ctx).StartSegment("CreateStatusTag").End()
	var req request.CreateStatusTag
	if err := ctx.Bind(&req); err != nil {
		return echo.ErrBadRequest
	}

	return ctx.JSON(stc.inputPort.CreateStatusTag(ctx.Request().Context(), &req))
}

// StatusTag		godoc
//
// @Summary			Get all StatusTag
// @Description	Get all StatusTag
// @Tags				StatusTag
// @Produce			json
// @Success			200											{array}		response.StatusTag								"success response"
// @Failure			400											{object}	nil																"error response"
// @Failure			500											{object}	nil																"error response"
// @Router			/status_tags						[GET]
func (stc *StatusTagController) FindAllStatusTag(ctx echo.Context) error {
	defer nrecho.FromContext(ctx).StartSegment("FindAllStatusTag").End()
	return ctx.JSON(stc.inputPort.FindAllStatusTag(ctx.Request().Context()))
}

// StatusTag		godoc
//
// @Summary			Update StatusTag by id
// @Description	Update StatusTag by id
// @Tags				StatusTag
// @Produce			json
// @Param				id											path			int												true		"status tag id"
// @Param				CreateStatusTagRequest	body			request.UpdateStatusTag		true		"request body"
// @Success			200											{object}	response.StatusTag								"success response"
// @Failure			400											{object}	nil																"error response"
// @Failure			500											{object}	nil																"error response"
// @Router			/status_tags/{id}				[PUT]
func (stc *StatusTagController) UpdateStatusTag(ctx echo.Context) error {
	defer nrecho.FromContext(ctx).StartSegment("UpdateStatusTag").End()
	var req request.UpdateStatusTag
	if err := ctx.Bind(&req); err != nil {
		return echo.ErrBadRequest
	}

	return ctx.JSON(stc.inputPort.UpdateStatusTag(ctx.Request().Context(), &req))
}
