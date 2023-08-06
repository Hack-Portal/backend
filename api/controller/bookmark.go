package controller

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hackhack-Geek-vol6/backend/bootstrap"
	db "github.com/hackhack-Geek-vol6/backend/db/sqlc"
	"github.com/hackhack-Geek-vol6/backend/domain"
)

type BookmarkController struct {
	BookmarkUsecase domain.BookmarkUsecase
	Env             *bootstrap.Env
}

// CreateBookmark	godoc
// @Summary			Create new bookmark
// @Description		Create a bookmark from the specified hackathon ID
// @Tags			Bookmark
// @Produce			json
// @Param			CreateBookmarkRequestBody 	body 		CreateBookmarkRequestBody	true	"Create Bookmark Request Body"
// @Success			200							{object}	BookmarkResponse			"create success response"
// @Failure 		400							{object}	ErrorResponse				"bad request response"
// @Failure 		500							{object}	ErrorResponse				"server error response"
// @Router       	/bookmarks 					[post]
func (bc *BookmarkController) CreateBookmark(ctx *gin.Context) {
	var reqBody domain.CreateBookmarkRequest
	if err := ctx.ShouldBindJSON(&reqBody); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	response, err := bc.BookmarkUsecase.CreateBookmark(ctx, db.CreateBookmarkParams{HackathonID: reqBody.HackathonID, UserID: reqBody.UserID})
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
// @Param			user_id		 	path 			string				true	"Delete Bookmark Request Body"
// @Success			200				{object}		BookmarkResponse	"delete success response"
// @Failure 		400				{object}		ErrorResponse		"bad request response"
// @Failure 		500				{object}		ErrorResponse		"server error response"
// @Router       	/bookmarks/:user_id 		[delete]
func (bc *BookmarkController) RemoveBookmark(ctx *gin.Context) {
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

	if err := bc.BookmarkUsecase.RemoveBookmark(ctx, reqURI.UserID, reqBody.HackathonID); err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, DeleteResponse{Result: fmt.Sprintf("delete successful")})
}

// ListBookmarkToHackathon	godoc
// @Summary			Get bookmarks
// @Description		Get my bookmarks
// @Tags			Bookmark
// @Produce			json
// @Param			ListBookmarkRequestQueries 	formData 		string				true	"Delete Bookmark Request Body"
// @Success			200							{array}			BookmarkResponse	"delete success response"
// @Failure 		400							{object}		ErrorResponse		"bad request response"
// @Failure 		500							{object}		ErrorResponse		"server error response"
// @Router       	/bookmarks/:user_id  		[get]
func (bc *BookmarkController) ListBookmarkToHackathon(ctx *gin.Context) {
	var (
		reqURI  domain.BookmarkRequestWildCard
		reqBody domain.ListBookmarkRequestQueries
	)
	if err := ctx.ShouldBindUri(&reqURI); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	if err := ctx.ShouldBindQuery(&reqBody); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	response, err := bc.BookmarkUsecase.GetBookmarks(ctx, reqURI.UserID, reqBody)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, response)
}
