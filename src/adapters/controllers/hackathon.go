package controllers

import (
	"log"
	"mime/multipart"

	"github.com/Hack-Portal/backend/src/datastructure/request"
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
