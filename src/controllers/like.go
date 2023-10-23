package controller

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hackhack-Geek-vol6/backend/pkg/logger"
	"github.com/hackhack-Geek-vol6/backend/pkg/repository"
	"github.com/hackhack-Geek-vol6/backend/src/domain/request"
	"github.com/hackhack-Geek-vol6/backend/src/transaction"
	"github.com/hackhack-Geek-vol6/backend/src/usecases/inputport"
	usecase "github.com/hackhack-Geek-vol6/backend/src/usecases/interactor"
)

type LikeController struct {
	LikeUsecase inputport.LikeUsecase
	l           logger.Logger
}

func NewLikeController(store transaction.SQLStore, l logger.Logger) *LikeController {
	return &LikeController{
		LikeUsecase: usecase.NewLikeUsercase(store, l),
		l:           l,
	}
}

// CreateLike	godoc
//
//	@Summary		Create new like
//	@Description	Create a like from the specified Account ID and hackathon ID
//	@Tags			Like
//	@Produce		json
//	@Param			CreateLikeRequest	body		request.CreateLike	true	"Create Like Request Body"
//	@Success		200					{object}	SuccessResponse			"create success response"
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
//	@Param			opus				query		request.RemoveLike				true	"opus"
//	@Success		200					{object}	SuccessResponse	"delete success response"
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
//	@Param			ListRequest			formData	request.ListRequest	true	"Like Request Body"
//	@Success		200					{array}		repository.Like	"success response"
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
