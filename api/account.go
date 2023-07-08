package api

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
	db "github.com/hackhack-Geek-vol6/backend/db/sqlc"
	"github.com/hackhack-Geek-vol6/backend/util"
	"github.com/hackhack-Geek-vol6/backend/util/token"
)

// アカウントを作る時のリクエストパラメータ
type CreateAccountRequestParam struct {
	Username        string `json:"username"`
	Icon            []byte `json:"icon"`
	ExplanatoryText string `json:"explanatory_text"`
	LocateID        int32  `json:"locate_id"`
	Password        string `json:"password"`
	ShowLocate      bool   `json:"show_locate"`
	ShowRate        bool   `json:"show_rate"`
}

// アカウントに関するレスポンス
type AccountResponses struct {
	UserID          string `json:"user_id"`
	Username        string `json:"username"`
	Icon            []byte `json:"icon"`
	ExplanatoryText string `json:"explanatory_text"`
	Locate          string `json:"locate"`
	ShowLocate      bool   `json:"show_locate"`
	ShowRate        bool   `json:"show_rate"`
}

// アカウント作成
func (server *Server) CreateAccount(ctx *gin.Context) {
	var request CreateAccountRequestParam
	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	claims := ctx.MustGet(AuthorizationClaimsKey).(*token.CustomClaims)
	if request.Password == "" {
		var err error
		request.Password, err = util.HashPassword(request.Password)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, errorResponse(err))
			return
		}
	}

	arg := db.CreateAccountParams{
		UserID:   claims.UserId,
		Username: request.Username,
		Icon:     request.Icon,
		ExplanatoryText: sql.NullString{
			String: request.ExplanatoryText,
			Valid:  false,
		},
		LocateID: request.LocateID,
		Rate:     0,
		HashedPassword: sql.NullString{
			String: request.Password,
			Valid:  false,
		},
		Email:      claims.Email,
		ShowLocate: request.ShowLocate,
		ShowRate:   request.ShowRate,
	}

	account, err := server.store.CreateAccount(ctx, arg)
	if err != nil {
		// ToDo: IDがなかったときの分岐を作る
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	locate, err := server.store.GetLocate(ctx, account.LocateID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	response := AccountResponses{
		UserID:          account.UserID,
		Username:        account.Username,
		Icon:            account.Icon,
		ExplanatoryText: account.ExplanatoryText.String,
		Locate:          locate.Name,
		ShowLocate:      account.ShowLocate,
		ShowRate:        account.ShowRate,
	}
	ctx.JSON(http.StatusOK, response)
}

// アカウントを取得するさいのパラメータ
type GetAccountRequestParams struct {
	ID string `uri:"id"`
}

// アカウントを取得する
func (server *Server) GetAccount(ctx *gin.Context) {
	var request GetAccountRequestParams
	if err := ctx.ShouldBindUri(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	// ToDo:単一アカウント取得にauthヘッダが必要か否か

	account, err := server.store.GetAccount(ctx, request.ID)
	if err != nil {
		// ToDo: IDがなかったときの分岐を作る
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
	}

	response := AccountResponses{
		UserID:          account.UserID,
		Username:        account.Username,
		Icon:            account.Icon,
		ExplanatoryText: account.ExplanatoryText.String,
		Locate:          account.Locate,
		ShowLocate:      account.ShowLocate,
		ShowRate:        account.ShowRate,
	}
	ctx.JSON(http.StatusOK, response)
}

type ListAccountRequestParam struct {
	Conditions string
	PageSize   int32
	PageID     int32
}

func (server *Server) ListAccount(ctx *gin.Context) {

}
