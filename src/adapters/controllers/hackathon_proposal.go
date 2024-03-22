package controllers

import (
	"github.com/Hack-Portal/backend/src/datastructure/request"
	"github.com/Hack-Portal/backend/src/usecases/ports"
	"github.com/labstack/echo/v4"
	"github.com/newrelic/go-agent/v3/integrations/nrecho-v4"
)

type HackathonProposalController struct {
	input ports.HackathonProposalInputBoundary
}

func NewHackathonProposalController(input ports.HackathonProposalInputBoundary) *HackathonProposalController {
	return &HackathonProposalController{
		input: input,
	}
}

func (hc *HackathonProposalController) CreateHackathonProposal(ctx echo.Context) error {
	defer nrecho.FromContext(ctx).StartSegment("CreateHackathonProposal").End()

	var input request.CreateHackathonProposal
	if err := ctx.Bind(&input); err != nil {
		return echo.ErrBadRequest
	}

	return ctx.JSON(hc.input.CreateHackathonProposal(ctx.Request().Context(), request.CreateHackathonProposal{
		URL: input.URL,
	}))
}

// func (hc *HackathonProposalController) ListHackathonProposal(ctx echo.Context) error {
// 	defer nrecho.FromContext(ctx).StartSegment("ListHackathonProposal").End()

// 	input := request.ListHackathonProposal{
// 		PageSize: 10,
// 		PageID:   1,
// 	}
// 	if ctx.Bind(&input) != nil {
// 		return echo.ErrBadRequest
// 	}

// 	return ctx.JSON(hc.input.ListHackathonProposal(ctx.Request().Context(),
// 		input,
// 	))
// }

// func (hc HackathonProposalController) DeleteHackathonProposal(ctx echo.Context) error {
// 	defer nrecho.FromContext(ctx).StartSegment("DeleteHackathonProposal").End()

// 	var input request.DeleteHackathonProposal
// 	if err := ctx.Bind(&input); err != nil {
// 		return echo.ErrBadRequest
// 	}

// 	return ctx.JSON(hc.input.DeleteHackathonProposal(ctx.Request().Context(),
// 		input.HackathonProposalID,
// 	))
// }
