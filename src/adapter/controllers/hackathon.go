package controllers

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/hackhack-Geek-vol6/backend/src/entities/request"
	"github.com/hackhack-Geek-vol6/backend/src/usecase/input"
	"github.com/hackhack-Geek-vol6/backend/src/usecase/output"
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
