package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/hackhack-Geek-vol6/backend/src/datastructs/input"
	_ "github.com/hackhack-Geek-vol6/backend/src/datastructs/output"
	"github.com/hackhack-Geek-vol6/backend/src/usecases"
	"github.com/hackhack-Geek-vol6/backend/src/usecases/dai"
	"github.com/hackhack-Geek-vol6/backend/src/usecases/ports"
)

// ここでいうControllerとは、任意の構造にデータをバインドし、Usecase Interactorに渡すことを指す

type hackathonController struct {
	Interactor ports.HackathonInputBoundary
}

func NewHackathonController(out ports.HackathonOutputBoundary, repository dai.HackathonRepository, firebase dai.FirebaseRepository) *hackathonController {
	return &hackathonController{
		Interactor: usecases.NewHackathonInteractor(out, repository, firebase),
	}
}

// CreateHackathon	godoc
//
// @Summary			Create Hackathon
// @Description	Register a hackathon from given parameters
// @Tags				Hackathon
// @Produce			json
// @Param				CreateHackathonRequest		body		input.HackathonCreate	true	"create hackathon Request Body"
// @Success			200			{object}					output.CreateHackathon							"success response"
// @Failure			400			{object}					nil																	"error response"
// @Failure			500			{object}					nil																	"error response"
// @Router			/hackathons																																																																											[post]
func (hc *hackathonController) Create(ctx *gin.Context) {
	var reqBody input.HackathonCreate
	ctx.BindJSON(&reqBody)

	image, _ := CheckFile(ctx.Request)

	ctx.JSON(hc.Interactor.Create(reqBody, image))
}

func (hc *hackathonController) ReadAll(ctx *gin.Context) {

}

func (hc *hackathonController) Update(ctx *gin.Context) {

}

func (hc *hackathonController) DeleteAll(ctx *gin.Context) {

}

func (hc *hackathonController) Delete(ctx *gin.Context) {

}
