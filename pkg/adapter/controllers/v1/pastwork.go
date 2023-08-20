package controller

import (
	"bytes"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hackhack-Geek-vol6/backend/pkg/bootstrap"
	"github.com/hackhack-Geek-vol6/backend/pkg/domain"
	"github.com/hackhack-Geek-vol6/backend/pkg/usecase/inputport"
)

type PastWorkController struct {
	PastWorkUsecase inputport.PastworkUsecase
	Env             *bootstrap.Env
}

func (pc *PastWorkController) CreatePastWork(ctx *gin.Context) {
	var (
		reqBody domain.CreatePastWorkRequestBody
		image   []byte
	)
	if err := ctx.ShouldBindJSON(&reqBody); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	file, _, err := ctx.Request.FormFile(ImageKey)
	if err != nil {
		switch err.Error() {
		case MultiPartNextPartEoF:
			ctx.JSON(400, errorResponse(err))
			return
		case HttpNoSuchFile:
			ctx.JSON(400, errorResponse(err))
			return
		case RequestContentTypeIsnt:
			break
		default:
			ctx.JSON(400, errorResponse(err))
			return
		}
	} else {
		icon := bytes.NewBuffer(nil)
		if _, err := io.Copy(icon, file); err != nil {
			ctx.JSON(http.StatusInternalServerError, errorResponse(err))
			return
		}
		image = icon.Bytes()
	}

	response, err := pc.PastWorkUsecase.CreatePastWork(ctx, domain.CreatePastWorkParams{
		Name:               reqBody.Name,
		ThumbnailImage:     string(image),
		ExplanatoryText:    reqBody.ExplanatoryText,
		PastWorkTags:       reqBody.PastWorkTags,
		PastWorkFrameworks: reqBody.PastWorkFrameworks,
		AccountPastWorks:   reqBody.AccountPastWorks,
	}, image)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, response)
}
