package controller

import (
	"bytes"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
	repository "github.com/hackhack-Geek-vol6/backend/pkg/adapter/gateways/repository/datasource"
	"github.com/hackhack-Geek-vol6/backend/pkg/adapter/gateways/repository/transaction"
	"github.com/hackhack-Geek-vol6/backend/pkg/bootstrap"
	"github.com/hackhack-Geek-vol6/backend/pkg/domain"
	"github.com/hackhack-Geek-vol6/backend/pkg/usecase/inputport"
	"github.com/lib/pq"
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
		ExplanatoryText:    reqBody.ExplanatoryText,
		PastWorkTags:       reqBody.PastWorkTags,
		PastWorkFrameworks: reqBody.PastWorkFrameworks,
		AccountPastWorks:   reqBody.AccountPastWorks,
	}, image)
	if err != nil {
		if err != nil {
			// すでに登録されている場合と参照エラーの処理
			if pqErr, ok := err.(*pq.Error); ok {
				switch pqErr.Code.Name() {
				case transaction.ForeignKeyViolation, transaction.UniqueViolation:
					ctx.JSON(http.StatusForbidden, errorResponse(err))
					return
				}
			}
			ctx.JSON(http.StatusInternalServerError, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, response)
}

func (pc *PastWorkController) GetPastWork(ctx *gin.Context) {
	var reqURI domain.PastWorksRequestWildCard
	if err := ctx.ShouldBindUri(&reqURI); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	response, err := pc.PastWorkUsecase.GetPastWork(ctx, reqURI.Opus)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, response)
}

func (pc *PastWorkController) ListPastWork(ctx *gin.Context) {
	var reqQuery domain.ListRequest

	if err := ctx.ShouldBindQuery(&reqQuery); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	response, err := pc.PastWorkUsecase.ListPastWork(ctx, reqQuery)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, response)
}

func (pc *PastWorkController) UpdatePastWork(ctx *gin.Context) {
	var (
		reqBody domain.CreatePastWorkRequestBody
		reqURI  domain.PastWorksRequestWildCard
	)
	if err := ctx.ShouldBindUri(&reqURI); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	if err := ctx.ShouldBindJSON(&reqBody); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	response, err := pc.PastWorkUsecase.UpdatePastWork(ctx, domain.UpdatePastWorkRequestBody{
		Name:               reqBody.Name,
		ExplanatoryText:    reqBody.ExplanatoryText,
		PastWorkTags:       reqBody.PastWorkTags,
		PastWorkFrameworks: reqBody.PastWorkFrameworks,
		AccountPastWorks:   reqBody.AccountPastWorks,
	})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, response)
}

func (pc *PastWorkController) DeletePastWork(ctx *gin.Context) {
	var reqURI domain.PastWorksRequestWildCard
	if err := ctx.ShouldBindUri(&reqURI); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	if err := pc.PastWorkUsecase.DeletePastWork(ctx, repository.DeletePastWorksByIDParams{
		Opus:     reqURI.Opus,
		IsDelete: true,
	}); err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, SuccessResponse{Result: "Delete Successful"})
}
