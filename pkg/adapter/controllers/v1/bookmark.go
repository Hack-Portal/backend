package controller

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	repository "github.com/hackhack-Geek-vol6/backend/pkg/adapter/gateways/repository/datasource"
	"github.com/hackhack-Geek-vol6/backend/pkg/bootstrap"
	"github.com/hackhack-Geek-vol6/backend/pkg/domain"
	"github.com/hackhack-Geek-vol6/backend/pkg/usecase/inputport"
)

type LikeController struct {
	LikeUsecase inputport.LikeUsecase
	Env         *bootstrap.Env
}

func (bc *LikeController) CreateBookmark(ctx *gin.Context) {
	var reqBody domain.CreateBookmarkRequest
	if err := ctx.ShouldBindJSON(&reqBody); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	response, err := bc.LikeUsecase.CreateLike(ctx, repository.CreateLikesParams{Opus: reqBody.Opus, AccountID: reqBody.AccountID})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, response)
}

func (bc *LikeController) RemoveBookmark(ctx *gin.Context) {
	var (
		reqURI  domain.BookmarkRequestWildCard
		reqBody domain.RemoveBookmarkRequestQueries
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

func (bc *LikeController) ListBookmark(ctx *gin.Context) {
	var (
		reqURI  domain.BookmarkRequestWildCard
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
