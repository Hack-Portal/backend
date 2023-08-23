package controller

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hackhack-Geek-vol6/backend/pkg/bootstrap"
	"github.com/hackhack-Geek-vol6/backend/pkg/domain"
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
func (bc *LikeController) CreateLike(ctx *gin.Context) {
	var reqBody domain.CreateLikeRequest
	if err := ctx.ShouldBindJSON(&reqBody); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
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
//	@Success		200					{object}	domain.LikeResponse	"delete success response"
//	@Failure		400					{object}	ErrorResponse		"bad request response"
//	@Failure		500					{object}	ErrorResponse		"server error response"
//	@Router			/like/{account_id} 	[delete]
func (bc *LikeController) RemoveLike(ctx *gin.Context) {
	var (
		reqURI  domain.LikeRequestWildCard
		reqBody domain.RemoveLikeRequestQueries
	)
	if err := ctx.ShouldBindUri(&reqURI); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	if err := ctx.ShouldBindQuery(&reqBody); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	if err := bc.LikeUsecase.RemoveLike(ctx, reqURI.AccountID, reqBody.Opus); err != nil {
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
func (bc *LikeController) ListLike(ctx *gin.Context) {
	var (
		reqURI  domain.LikeRequestWildCard
		reqBody domain.ListRequest
	)
	if err := ctx.ShouldBindUri(&reqURI); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	if err := ctx.ShouldBindQuery(&reqBody); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	response, err := bc.LikeUsecase.GetLike(ctx, reqURI.AccountID, reqBody)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, response)
}
