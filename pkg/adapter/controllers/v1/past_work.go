package controller

import (
	"bytes"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
	repository "github.com/hackhack-Geek-vol6/backend/pkg/adapter/gateways/repository/datasource"
	"github.com/hackhack-Geek-vol6/backend/pkg/bootstrap"
	"github.com/hackhack-Geek-vol6/backend/pkg/domain"
	"github.com/hackhack-Geek-vol6/backend/pkg/usecase/inputport"
	util "github.com/hackhack-Geek-vol6/backend/pkg/util/etc"
)

type PastWorkController struct {
	PastWorkUsecase inputport.PastworkUsecase
	Env             *bootstrap.Env
}

// CreatePastWork	godoc
//
//	@Summary		Create new past work
//	@Description	Create a past work from the requested body
//	@Tags			PastWorks
//	@Produce		json
//	@Param			CreatePastWorkRequest	body		domain.PastWorkRequestBody	true	"Create PastWork Request"
//	@Success		200						{object}	domain.PastWorkResponse		"create success response"
//	@Failure		400						{object}	ErrorResponse				"bad request response"
//	@Failure		500						{object}	ErrorResponse				"server error response"
//	@Router			/pastworks	[post]
func (pc *PastWorkController) CreatePastWork(ctx *gin.Context) {
	var (
		reqBody    domain.PastWorkRequestBody
		image      []byte
		tags       []int32
		frameworks []int32
	)
	if err := ctx.ShouldBind(&reqBody); err != nil {
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
	if len(reqBody.PastWorkTags) != 0 {
		tags, err = util.StringToArrayInt32(reqBody.PastWorkTags)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, errorResponse(err))
			return
		}
	}

	if len(reqBody.PastWorkFrameworks) != 0 {
		frameworks, err = util.StringToArrayInt32(reqBody.PastWorkFrameworks)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, errorResponse(err))
			return
		}
	}

	response, err := pc.PastWorkUsecase.CreatePastWork(ctx, domain.CreatePastWorkParams{
		Name:               reqBody.Name,
		ExplanatoryText:    reqBody.ExplanatoryText,
		PastWorkTags:       tags,
		PastWorkFrameworks: frameworks,
		AccountPastWorks:   util.StringToArray(reqBody.AccountPastWorks),
	}, image)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, response)
}

// GetPastWork	godoc
//
//	@Summary		Get PastWork
//	@Description	Get PastWork
//	@Tags			PastWorks
//	@Produce		json
//	@Param			opus				path		string	true	"PastWorks API wildcard"
//	@Success		200					{object}	domain.PastWorkResponse
//	@Failure		400					{object}	ErrorResponse	"error response"
//	@Failure		500					{object}	ErrorResponse	"error response"
//	@Router			/pastworks/{opus}	[get]
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

// ListPastWork	godoc
//
//	@Summary		List PastWork
//	@Description	List PastWork
//	@Tags			PastWorks
//	@Produce		json
//	@Param			ListRequest	query		domain.ListRequest	true	"List PastWork Request"
//	@Success		200			{array}		domain.ListPastWorkResponse
//	@Failure		400			{object}	ErrorResponse	"error response"
//	@Failure		500			{object}	ErrorResponse	"error response"
//	@Router			/pastworks	[get]
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

// UpdatePastWork	godoc
//
//	@Summary		Update PastWork
//	@Description	Update PastWork
//	@Tags			PastWorks
//	@Produce		json
//	@Param			opus					path		string						true	"PastWorks API wildcard"
//	@Param			UpdatePastWorkRequest	body		domain.PastWorkRequestBody	true	"Update PastWork Request"
//	@Success		200						{object}	domain.PastWorkResponse
//	@Failure		400						{object}	ErrorResponse	"error response"
//	@Failure		500						{object}	ErrorResponse	"error response"
//	@Router			/pastworks/{opus}	[put]
func (pc *PastWorkController) UpdatePastWork(ctx *gin.Context) {
	var (
		reqBody    domain.PastWorkRequestBody
		reqURI     domain.PastWorksRequestWildCard
		tags       []int32
		frameworks []int32
		err        error
	)
	if err := ctx.ShouldBindUri(&reqURI); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	if err := ctx.ShouldBind(&reqBody); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	if len(reqBody.PastWorkTags) != 0 {
		tags, err = util.StringToArrayInt32(reqBody.PastWorkTags)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, errorResponse(err))
			return
		}
	}

	if len(reqBody.PastWorkFrameworks) != 0 {
		frameworks, err = util.StringToArrayInt32(reqBody.PastWorkFrameworks)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, errorResponse(err))
			return
		}
	}

	response, err := pc.PastWorkUsecase.UpdatePastWork(ctx, domain.UpdatePastWorkParams{
		Opus:               reqURI.Opus,
		Name:               reqBody.Name,
		ExplanatoryText:    reqBody.ExplanatoryText,
		PastWorkTags:       tags,
		PastWorkFrameworks: frameworks,
		AccountPastWorks:   util.StringToArray(reqBody.AccountPastWorks),
	})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, response)
}

// DeletePastWork	godoc
//
//	@Summary		Delete PastWork
//	@Description	Delete PastWork
//	@Tags			PastWorks
//	@Produce		json
//	@Param			opus				path		string	true	"PastWorks API wildcard"
//	@Success		200					{object}	SuccessResponse
//	@Failure		400					{object}	ErrorResponse	"error response"
//	@Failure		500					{object}	ErrorResponse	"error response"
//	@Router			/pastworks/{opus} 	[delete]
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
