package presenters

import (
	"github.com/gin-gonic/gin"
	"github.com/hackhack-Geek-vol6/backend/src/entities"
	"github.com/hackhack-Geek-vol6/backend/src/usecase/output"
)

type hackathonPresenter struct {
	ctx *gin.Context
}

func NewHackathonPresenter() output.HackathonOutputPort {
	return &hackathonPresenter{}
}

func (h *hackathonPresenter) RenderCreate() error {
	h.ctx.JSON(200, nil)
	return nil
}

func (h *hackathonPresenter) RenderGet(hackathons []*entities.Hackathon) error
func (h *hackathonPresenter) RenderUpdateByID(hackathon *entities.Hackathon) error
func (h *hackathonPresenter) RenderDeleteByID(hackathonID int32) error

func (h *hackathonPresenter) RenderApprove(hackathonID int32) error

func (h *hackathonPresenter) RenderError(err error) error
