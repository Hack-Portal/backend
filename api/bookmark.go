package api

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	db "github.com/hackhack-Geek-vol6/backend/db/sqlc"
	"github.com/hackhack-Geek-vol6/backend/util/token"
)

type CreateBookmarkRequestBody struct {
	UserID      string `json:"user_id"`
	HackathonID int32  `json:"hackathon_id"`
}

type BookmarkResponse struct {
	HackathonID int32     `json:"hackathon_id"`
	Name        string    `json:"name"`
	Icon        string    `json:"icon"`
	Description string    `json:"description"`
	Link        string    `json:"link"`
	Expired     time.Time `json:"expired"`
	StartDate   time.Time `json:"start_date"`
	Term        int32     `json:"term"`
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
func (server *Server) CreateBookmark(ctx *gin.Context) {
	var request CreateBookmarkRequestBody
	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	payload := ctx.MustGet(AuthorizationClaimsKey).(*token.FireBaseCustomToken)
	account, err := server.store.GetAccountByID(ctx, request.UserID)
	if err != nil {
		// UIDについてのカラムがない場合の処理を作る必要がある (badRequest)
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	// 認証ヘッダとUID登録先のEmailが一致しない場合
	if payload.Email != account.Email {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	bookmark, err := server.store.CreateBookmark(ctx, db.CreateBookmarkParams{
		HackathonID: request.HackathonID,
		UserID:      account.UserID,
	})
	if err != nil {
		// NoRowSQLエラーの場合の処理を作る必要がある (badRequest)
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	hackathon, err := server.store.GetHackathonByID(ctx, bookmark.HackathonID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	response := BookmarkResponse{
		HackathonID: hackathon.HackathonID,
		Name:        hackathon.Name,
		Icon:        hackathon.Icon.String,
		Description: hackathon.Description,
		Link:        hackathon.Link,
		Expired:     hackathon.Expired,
		StartDate:   hackathon.StartDate,
		Term:        hackathon.Term,
	}
	ctx.JSON(http.StatusOK, response)
}

type RemoveBookmarkRequestURI struct {
	HackathonID int32 `uri:"hackathon_id"`
}

// RemoveBookmark	godoc
// @Summary			Delete bookmark
// @Description		Delete the bookmark of the specified hackathon ID
// @Tags			Bookmark
// @Produce			json
// @Param			hackathon_id 	path 			string				true	"Delete Bookmark Request Body"
// @Success			200				{object}		BookmarkResponse	"delete success response"
// @Failure 		400				{object}		ErrorResponse		"bad request response"
// @Failure 		500				{object}		ErrorResponse		"server error response"
// @Router       	/bookmarks/{hackathon_id} 		[delete]
func (server *Server) RemoveBookmark(ctx *gin.Context) {
	var request RemoveBookmarkRequestURI
	if err := ctx.ShouldBindUri(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	payload := ctx.MustGet(AuthorizationClaimsKey).(*token.FireBaseCustomToken)
	account, err := server.store.GetAccountByEmail(ctx, payload.Email)
	if err != nil {
		// TODO: UIDについてのカラムがない場合の処理を作る必要がある (badRequest)
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	hackathon, err := server.store.GetHackathonByID(ctx, request.HackathonID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	_, err = server.store.SoftRemoveBookmark(ctx, db.SoftRemoveBookmarkParams{
		UserID:      account.UserID,
		HackathonID: request.HackathonID,
	})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	response := BookmarkResponse{
		HackathonID: hackathon.HackathonID,
		Name:        hackathon.Name,
		Icon:        hackathon.Icon.String,
		Description: hackathon.Description,
		Link:        hackathon.Link,
		Expired:     hackathon.Expired,
		StartDate:   hackathon.StartDate,
		Term:        hackathon.Term,
	}

	ctx.JSON(http.StatusOK, response)
}

type ListBookmarkRequestQueries struct {
	PageSize int32 `form:"page_size" binding:"required"`
	PageID   int32 `form:"page_id" binding:"required"`
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
// @Router       	/bookmarks/{hackathon_id} 	[get]
func (server *Server) ListBookmarkToHackathon(ctx *gin.Context) {
	var request ListBookmarkRequestQueries
	if err := ctx.ShouldBindQuery(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	payload := ctx.MustGet(AuthorizationClaimsKey).(*token.FireBaseCustomToken)
	account, err := server.store.GetAccountByEmail(ctx, payload.Email)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	bookmarks, err := server.store.ListBookmarkByUserID(ctx, account.UserID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	var response []BookmarkResponse
	for _, bookmark := range bookmarks {
		hackathon, err := server.store.GetHackathonByID(ctx, bookmark.HackathonID)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, errorResponse(err))
			return
		}
		response = append(response, BookmarkResponse{
			HackathonID: hackathon.HackathonID,
			Name:        hackathon.Name,
			Icon:        hackathon.Icon.String,
			Description: hackathon.Description,
			Link:        hackathon.Link,
			Expired:     hackathon.Expired,
			StartDate:   hackathon.StartDate,
			Term:        hackathon.Term,
		})
	}
	ctx.JSON(http.StatusOK, response)
}
