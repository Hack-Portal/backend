package controllers

import (
	"log"
	"mime/multipart"

	"github.com/Hack-Portal/backend/src/datastructure/request"
	_ "github.com/Hack-Portal/backend/src/datastructure/response"
	"github.com/Hack-Portal/backend/src/usecases/ports"
	"github.com/labstack/echo/v4"
)

type HackathonController struct {
	input ports.HackathonInputBoundary
}

func NewHackathonController(input ports.HackathonInputBoundary) *HackathonController {
	return &HackathonController{
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
func (hc *HackathonController) CreateHackathon(ctx echo.Context) error {
	var input request.CreateHackathon
	if err := ctx.Bind(&input); err != nil {
		log.Println(err)
		return echo.ErrBadRequest
	}

	form, err := ctx.MultipartForm()
	if err != nil {
		log.Println(err)
		return echo.ErrBadRequest
	}
	var image *multipart.FileHeader

	imageFiles, ok := form.File["image"]
	if !ok {
		image = nil
	} else {
		image = imageFiles[0]
	}

	log.Println("statuses", input.Statuses)

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
// @Summary			Get Hackathon
// @Description	Get Hackathon
// @Tags				Hackathon
// @Produce			json
// @Param				hackathon_id						path			string									true			"request body"
// @Success			200											{object}	response.GetHackathon							"success response"
// @Failure			400											{object}	nil																"error response"
// @Failure			500											{object}	nil																"error response"
// @Router			/hackathons/{hackathon_id}				[GET]
func (hc *HackathonController) GetHackathon(ctx echo.Context) error {
	var input request.GetHackathon
	if ctx.Bind(&input) != nil {
		return echo.ErrBadRequest
	}

	return ctx.JSON(hc.input.GetHackathon(ctx.Request().Context(), input.HackathonID))
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
func (hc *HackathonController) ListHackathons(ctx echo.Context) error {
	var input request.ListHackathon
	if ctx.Bind(&input) != nil {
		return echo.ErrBadRequest
	}

	return ctx.JSON(hc.input.ListHackathon(ctx.Request().Context(),
		input.PageSize,
		input.PageID,
	))
}
