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

// CreateBookmark	godoc
// @Summary			Create new bookmark
// @Description		Create a bookmark from the specified hackathon ID
// @Tags			Bookmark
// @Produce			json
// @Param			domain.CreateBookmarkRequest 	body 		domain.CreateBookmarkRequest	true	"Create Bookmark Request Body"
// @Success			200							{object}	domain.BookmarkResponse			"create success response"
// @Failure 		400							{object}	ErrorResponse				"bad request response"
// @Failure 		500							{object}	ErrorResponse				"server error response"
// @Router       	/bookmarks 					[post]
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

// RemoveBookmark	godoc
// @Summary			Delete bookmark
// @Description		Delete the bookmark of the specified hackathon ID
// @Tags			Bookmark
// @Produce			json
// @Param			account_id		 	path 			string				true	"Delete Bookmark Request Body"
// @Success			200				{object}		domain.BookmarkResponse	"delete success response"
// @Failure 		400				{object}		ErrorResponse		"bad request response"
// @Failure 		500				{object}		ErrorResponse		"server error response"
// @Router       	/bookmarks/:user_id 		[delete]
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

// ListBookmarkToHackathon	godoc
// @Summary			Get bookmarks
// @Description		Get my bookmarks
// @Tags			Bookmark
// @Produce			json
// @Param			domain.ListRequest 	formData 		string				true	"Delete Bookmark Request Body"
// @Success			200							{array}			domain.BookmarkResponse	"delete success response"
// @Failure 		400							{object}		ErrorResponse		"bad request response"
// @Failure 		500							{object}		ErrorResponse		"server error response"
// @Router       	/bookmarks/:user_id  		[get]
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
