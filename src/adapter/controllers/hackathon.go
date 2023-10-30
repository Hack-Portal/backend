package controllers

import (
	"strconv"
	"temp/src/entities/request"
	"temp/src/usecase/input"
	"temp/src/usecase/output"

	"github.com/gin-gonic/gin"
)

type hackathonController struct {
	input  input.HackathonInputPort
	output output.ErrorOutputPort
}

func NewHackathonController(input input.HackathonInputPort, output output.ErrorOutputPort) *hackathonController {
	return &hackathonController{
		input:  input,
		output: output,
	}
}

func (h *hackathonController) Create(ctx *gin.Context) {
	var reqBody request.CreateHackathon

	if err := ctx.ShouldBind(&reqBody); err != nil {
		h.output.BadRequest(err)
		return
	}

	image, err := CheckFile(ctx.Request)
	if err != nil {
		h.output.BadRequest(err)
		return
	}

	if err := h.input.Create(reqBody, image); err != nil {
		switch err.(type) {
		case *strconv.NumError:
			h.output.BadRequest(err)
		default:
			h.output.InternalServerError(err)
		}
		return
	}
}
