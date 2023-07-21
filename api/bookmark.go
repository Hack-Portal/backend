package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	db "github.com/hackhack-Geek-vol6/backend/db/sqlc"
	"github.com/hackhack-Geek-vol6/backend/util/token"
)

type CreateBookmarkRequestBody struct {
	UserID      string `json:"user_id"`
	HackathonID int32  `json:"hackathon_id"`
}

// CreateBookmark	godoc
// @Summary			Create new bookmark
// @Description		Create new bookmark
// @Tags			Bookmark
// @Produce			json
// @Param			CreateBookmarkRequestBody 	body 	CreateBookmarkRequestBody	true	"New Bookmark Request Body"
// @Success			200			{object}		db.Hackathons	"create succsss response"
// @Failure 		400			{object}		ErrorResponse	"bad request response"
// @Failure 		500			{object}		ErrorResponse	"server error response"
// @Router       	/bookmarks 	[post]

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

	ctx.JSON(http.StatusOK, hackathon)
}

// RemoveBookmark	godoc
// @Summary			delete bookmark
// @Description		delete bookmark
// @Tags			Bookmark
// @Produce			json
// @Param			hackathon_id 	path 		string	true	"Delete Bookmark Request Body"
// @Success			200			{object}		db.Hackathons	"delete succsss response"
// @Failure 		400			{object}		ErrorResponse	"bad request response"
// @Failure 		500			{object}		ErrorResponse	"server error response"
// @Router       	/bookmarks/{hackathon_id} 	[delete]

type RemoveBookmarkRequestURI struct {
	HackathonID int32 `uri:"hackathon_id"`
}

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

	ctx.JSON(http.StatusOK, hackathon)
}

// ListBookmarkToHackathon	godoc
// @Summary			Get my bookmark
// @Description		Get bookmark
// @Tags			Bookmark
// @Produce			json
// @Param			ListBookmarkRequestQueries 	formData 		string	true	"Delete Bookmark Request Body"
// @Success			200			{object}		db.Hackathons	"delete succsss response"
// @Failure 		400			{object}		ErrorResponse	"bad request response"
// @Failure 		500			{object}		ErrorResponse	"server error response"
// @Router       	/bookmarks/{hackathon_id} 	[get]

type ListBookmarkRequestQueries struct {
	PageSize int32 `form:"page_size"`
	PageID   int32 `form:"page_id"`
}

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
	var hackathons []db.Hackathons
	for _, bookmark := range bookmarks {
		hackathon, err := server.store.GetHackathonByID(ctx, bookmark.HackathonID)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, errorResponse(err))
			return
		}
		hackathons = append(hackathons, hackathon)
	}
	ctx.JSON(http.StatusOK, hackathons)
}
