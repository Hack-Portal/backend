package controller

import (
	"bytes"
	"database/sql"
	"errors"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hackhack-Geek-vol6/backend/pkg/bootstrap"
	"github.com/hackhack-Geek-vol6/backend/pkg/domain"
	"github.com/hackhack-Geek-vol6/backend/pkg/usecase/inputport"
	"github.com/newrelic/go-agent/v3/integrations/nrgin"
)

type HackathonController struct {
	HackathonUsecase inputport.HackathonUsecase
	Env              *bootstrap.Env
}

// CreateHackathon	godoc
//
//	@Summary		Create Hackathon
//	@Description	Register a hackathon from given parameters
//	@Tags			Hackathon
//	@Produce		json
//	@Param			CreateHackathonRequestBody	body		domain.CreateHackathonRequestBody	true	"create hackathon Request Body"
//	@Success		200							{object}	domain.HackathonResponses			"success response"
//	@Failure		400							{object}	ErrorResponse						"error response"
//	@Failure		500							{object}	ErrorResponse						"error response"
//	@Router			/hackathons																																																																											[post]
func (hc *HackathonController) CreateHackathon(ctx *gin.Context) {
	txn := nrgin.Transaction(ctx)
	defer txn.End()
	var (
		reqBody domain.CreateHackathonRequestBody
		image   []byte
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

	response, err := hc.HackathonUsecase.CreateHackathon(ctx, reqBody, image)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, response)
}

// GetHackathon	godoc
//
//	@Summary		Get Hackathon
//	@Description	Get Hackathon
//	@Tags			Hackathon
//	@Produce		json
//	@Param			hackathon_id				path		string						true	"Hackathons API wildcard"
//	@Success		200							{object}	domain.HackathonResponses	"success response"
//	@Failure		400							{object}	ErrorResponse				"error response"
//	@Failure		403							{object}	ErrorResponse				"error response"
//	@Failure		500							{object}	ErrorResponse				"error response"
//	@Router			/hackathons/{hackathon_id} 																																				[get]
func (hc *HackathonController) GetHackathon(ctx *gin.Context) {
	txn := nrgin.Transaction(ctx)
	defer txn.End()
	var reqURI domain.HackathonRequestWildCard
	if err := ctx.ShouldBindUri(&reqURI); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	response, err := hc.HackathonUsecase.GetHackathon(ctx, reqURI.HackathonID)
	if err != nil {
		switch err.Error() {
		case sql.ErrNoRows.Error():
			err := errors.New("そんなハッカソンはないわ")
			ctx.JSON(http.StatusForbidden, errorResponse(err))
		default:
			err := errors.New("すまんサーバエラーや")
			ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		}
		return
	}

	ctx.JSON(http.StatusOK, response)
}

// ハッカソン一覧取得
// ハッカソン一覧を取得する際のパラメータ

// ListHackathons	godoc
//
//	@Summary		List Hackathon
//	@Description	List Hackathon
//	@Tags			Hackathon
//	@Produce		json
//	@Param			ListHackathonsParams	query		domain.ListHackathonsParams	true	"List hackathon Request queries"
//	@Success		200						{array}		domain.HackathonResponses	"success response"
//	@Failure		400						{object}	ErrorResponse				"error response"
//	@Failure		500						{object}	ErrorResponse				"error response"
//	@Router			/hackathons 																																																					[get]
func (hc *HackathonController) ListHackathons(ctx *gin.Context) {
	txn := nrgin.Transaction(ctx)
	defer txn.End()
	var reqQuery domain.ListHackathonsParams
	if err := ctx.ShouldBindQuery(&reqQuery); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	response, err := hc.HackathonUsecase.ListHackathons(ctx, reqQuery)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, response)
}
