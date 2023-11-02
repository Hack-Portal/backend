package controllers

import (
	"temp/src/datastructs/input"
	"temp/src/usecases"
	"temp/src/usecases/dai"
	"temp/src/usecases/inputport"
	"temp/src/usecases/outputport"

	"github.com/gin-gonic/gin"
)

// ここでいうControllerとは、任意の構造にデータをバインドし、Usecase Interactorに渡すことを指す

type hackathonController struct {
	Interactor inputport.HackathonInputPort
}

func NewHackathonController(out outputport.HackathonOutputPort, repository dai.HackathonRepository) *hackathonController {
	return &hackathonController{
		Interactor: usecases.NewHackathonInterface(out, repository),
	}
}

func (hc *hackathonController) Create(ctx *gin.Context) {
	var reqBody input.HackathonCreate

	ctx.BindJSON(&reqBody)

	image, _ := CheckFile(ctx.Request)

	hc.Interactor.Create(reqBody, image)
}
