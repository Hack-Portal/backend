package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	db "github.com/hackhack-Geek-vol6/backend/db/sqlc"
	"github.com/hackhack-Geek-vol6/backend/util/token"
)

// ブックマークを作る
// 認証必須
type BookmarkRequest struct {
	UserID      string `json:"user_id"`
	HackathonID int32  `json:"hackathon_id"`
}

func (server *Server) CreateBookmark(ctx *gin.Context) {
	var request BookmarkRequest
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

// ブックマークを削除する
// 認証必須
type RemoveBookmarkRequest struct {
	HackathonID int32 `uri:"hackathon_id"`
}

func (server *Server) RemoveBookmark(ctx *gin.Context) {
	var request RemoveBookmarkRequest
	if err := ctx.ShouldBindUri(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	payload := ctx.MustGet(AuthorizationClaimsKey).(*token.FireBaseCustomToken)
	account, err := server.store.GetAccountByEmail(ctx, payload.Email)
	if err != nil {
		// UIDについてのカラムがない場合の処理を作る必要がある (badRequest)
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

// ブックマーク一覧からハッカソンを取得する
// 認証必須
type ListBookmarkRequest struct {
	PageSize int32 `form:"page_size"`
	PageID   int32 `form:"page_id"`
}

func (server *Server) ListBookmarkToHackathon(ctx *gin.Context) {
	var request ListBookmarkRequest
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
