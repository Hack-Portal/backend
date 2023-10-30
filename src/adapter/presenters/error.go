package presenters

import (
	"net/http"
	"temp/src/usecase/output"

	"github.com/gin-gonic/gin"
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
