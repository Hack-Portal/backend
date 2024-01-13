package controllers

import (
	"mime/multipart"

	"github.com/Hack-Portal/backend/src/datastructure/request"

	// swaggerのために読み込んでる
	_ "github.com/Hack-Portal/backend/src/datastructure/response"
	"github.com/Hack-Portal/backend/src/usecases/ports"
	"github.com/labstack/echo/v4"
	"github.com/newrelic/go-agent/v3/newrelic"
)

type hackathonController struct {
	input ports.HackathonInputBoundary
}

// HackathonController はHackathonControllerのインターフェース
type HackathonController interface {
	CreateHackathon(ctx echo.Context) error
	ListHackathons(ctx echo.Context) error
	DeleteHackathon(ctx echo.Context) error
}

// NewHackathonController はHackathonControllerを返す
func NewHackathonController(input ports.HackathonInputBoundary) HackathonController {
	return &hackathonController{
		input: input,
	}
}

// Hackathon		godoc
//
// @Summary			Create Hackathon
// @Description	Create Hackathon
// @Tags				Hackathon
// @Produce			json
// @Param				CreateHackathonRequest	body			request.CreateHackathon	true			"request body"
// @Success			200											{object}	response.CreateHackathon					"success response"
// @Failure			400											{object}	nil																"error response"
// @Failure			500											{object}	nil																"error response"
// @Router			/hackathons							[POST]
func (hc *hackathonController) CreateHackathon(ctx echo.Context) error {
	defer newrelic.FromContext(ctx.Request().Context()).StartSegment("CreateHackathon").End()

	var input request.CreateHackathon
	if err := ctx.Bind(&input); err != nil {
		return echo.ErrBadRequest
	}

	form, err := ctx.MultipartForm()
	if err != nil {
		return echo.ErrBadRequest
	}
	var image *multipart.FileHeader

	imageFiles, ok := form.File["image"]
	if !ok {
		image = nil
	} else {
		image = imageFiles[0]
	}

	return ctx.JSON(hc.input.CreateHackathon(ctx.Request().Context(), &ports.InputCreatehackathonData{
		ImageFile: image,

		Name:      input.Name,
		Link:      input.Link,
		Expired:   input.Expired,
		StartDate: input.StartDate,
		Term:      input.Term,
		Statuses:  input.Statuses,
	}))
}

// Hackathon		godoc
//
// @Summary			List Hackathons
// @Description	List Hackathons
// @Tags				Hackathon
// @Produce			json
// @Param				ListHackathonRequest		query			request.ListHackathon		true			"request query"
// @Success			200											{array}	response.GetHackathon							"success response"
// @Failure			400											{object}	nil																"error response"
// @Failure			500											{object}	nil																"error response"
// @Router			/hackathons							[GET]
func (hc *hackathonController) ListHackathons(ctx echo.Context) error {
	defer newrelic.FromContext(ctx.Request().Context()).StartSegment("ListHackathons").End()

	var input request.ListHackathon = request.ListHackathon{
		PageSize: 10,
		PageID:   1,
		New:      true,
	}
	if ctx.Bind(&input) != nil {
		return echo.ErrBadRequest
	}

	return ctx.JSON(hc.input.ListHackathon(ctx.Request().Context(),
		input,
	))
}

// Hackathon		godoc
//
// @Summary			Delete Hackathons
// @Description	Delete Hackathons
// @Tags				Hackathon
// @Produce			json
// @Param				hackathon_id						path			string									true			"request body"
// @Success			200											{object}	response.DeleteHackathon					"success response"
// @Failure			400											{object}	nil																"error response"
// @Failure			500											{object}	nil																"error response"
// @Router			/hackathons/{hackathon_id}				[DELETE]
func (hc *hackathonController) DeleteHackathon(ctx echo.Context) error {
	defer newrelic.FromContext(ctx.Request().Context()).StartSegment("DeleteHackathon").End()

	var input request.DeleteHackathon
	if ctx.Bind(&input) != nil {
		return echo.ErrBadRequest
	}

	return ctx.JSON(hc.input.DeleteHackathon(ctx.Request().Context(), input.HackathonID))
}
