package presenters

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hackhack-Geek-vol6/backend/src/usecase/output"
)

type errorPresenter struct {
	ctx *gin.Context
}

func NewErrorPresenter() output.ErrorOutputPort {
	return &errorPresenter{}
}

func (e *errorPresenter) BadRequest(err error) {
	e.ctx.JSON(http.StatusBadRequest, nil)

}

func (e *errorPresenter) InternalServerError(err error)
func (e *errorPresenter) NotFound(err error)
func (e *errorPresenter) Unauthorized(err error)
