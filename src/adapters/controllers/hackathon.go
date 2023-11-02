package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/hackhack-Geek-vol6/backend/src/datastructs/input"
	"github.com/hackhack-Geek-vol6/backend/src/usecases"
	"github.com/hackhack-Geek-vol6/backend/src/usecases/dai"
	"github.com/hackhack-Geek-vol6/backend/src/usecases/inputport"
	"github.com/hackhack-Geek-vol6/backend/src/usecases/outputport"
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
