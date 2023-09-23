package controller

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	repository "github.com/hackhack-Geek-vol6/backend/pkg/adapter/gateways/repository/datasource"
	"github.com/hackhack-Geek-vol6/backend/pkg/bootstrap"
	"github.com/hackhack-Geek-vol6/backend/pkg/domain/request"
	"github.com/hackhack-Geek-vol6/backend/pkg/usecase/inputport"
)

type LikeController struct {
	LikeUsecase inputport.LikeUsecase
	Env         *bootstrap.Env
}

// CreateLike	godoc
//
//	@Summary		Create new like
//	@Description	Create a like from the specified Account ID and hackathon ID
//	@Tags			Like
//	@Produce		json
//	@Param			CreateLikeRequest	body		domain.CreateLikeRequest	true	"Create Like Request Body"
//	@Success		200					{object}	domain.LikeResponse			"create success response"
//	@Failure		400					{object}	ErrorResponse				"bad request response"
//	@Failure		500					{object}	ErrorResponse				"server error response"
//	@Router			/like	[post]
func (lc *LikeController) CreateLike(ctx *gin.Context) {
	var reqBody request.CreateLike
	if err := ctx.ShouldBindJSON(&reqBody); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	_, err := lc.LikeUsecase.CreateLike(ctx, repository.CreateLikesParams{
		Opus:      reqBody.Opus,
		AccountID: reqBody.AccountID,
	})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, SuccessResponse{Result: fmt.Sprintf("create successful")})
}

// RemoveLike	godoc
//
//	@Summary		Delete like
//	@Description	Delete a like from the specified Account ID
//	@Tags			Like
//	@Produce		json
//	@Param			account_id			path		string				true	"Delete Like Request Body"
//	@Param			opus				query		int32				true	"opus"
//	@Success		200					{object}	domain.LikeResponse	"delete success response"
//	@Failure		400					{object}	ErrorResponse		"bad request response"
//	@Failure		500					{object}	ErrorResponse		"server error response"
//	@Router			/like/{account_id} 	[delete]
func (lc *LikeController) RemoveLike(ctx *gin.Context) {
	var (
		reqURI  request.AccountWildCard
		reqBody request.RemoveLike
	)
	if err := ctx.ShouldBindUri(&reqURI); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	if err := ctx.ShouldBindQuery(&reqBody); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	if err := lc.LikeUsecase.RemoveLike(ctx, reqURI.AccountID, reqBody.Opus); err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, SuccessResponse{Result: fmt.Sprintf("delete successful")})
}

// ListLike	godoc
//
//	@Summary		Get likes
//	@Description	Get my likes
//	@Tags			Like
//	@Produce		json
//	@Param			account_id			path		string				true	"account_id"
//	@Param			ListRequest			formData	domain.ListRequest	true	"Like Request Body"
//	@Success		200					{array}		domain.LikeResponse	"success response"
//	@Failure		400					{object}	ErrorResponse		"bad request response"
//	@Failure		500					{object}	ErrorResponse		"server error response"
//	@Router			/like/{account_id} 	[get]
func (lc *LikeController) ListLike(ctx *gin.Context) {
	var (
		reqURI  request.AccountWildCard
		reqBody request.ListRequest
	)
	if err := ctx.ShouldBindUri(&reqURI); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	if err := ctx.ShouldBindQuery(&reqBody); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	response, err := lc.LikeUsecase.GetLike(ctx, reqURI.AccountID, reqBody)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, response)
}
