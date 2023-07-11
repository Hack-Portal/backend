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
	UserID          string  `json:"user_id" binding:"required"`
	Username        string  `json:"username" binding:"required"`
	Icon            string  `json:"icon"`
	ExplanatoryText string  `json:"explanatory_text"`
	LocateID        int32   `json:"locate_id" binding:"required"`
	Password        string  `json:"password"`
	ShowLocate      bool    `json:"show_locate"`
	ShowRate        bool    `json:"show_rate"`
	TechTags        []int32 `json:"tech_tags"`
	Frameworks      []int32 `json:"frameworks"`
}

// アカウントに関するレスポンス
type AccountResponses struct {
	UserID          string     `json:"user_id"`
	Username        string     `json:"username"`
	Icon            string     `json:"icon"`
	ExplanatoryText string     `json:"explanatory_text"`
	Rate            int32      `json:"rate"`
	Locate          db.Locates `json:"locate"`
	ShowLocate      bool       `json:"show_locate"`
	ShowRate        bool       `json:"show_rate"`

	TechTags   []db.TechTags
	Frameworks []db.Frameworks
}

func (server *Server) CreateAccount(ctx *gin.Context) {
	var request CreateAccountRequestParam
	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	payload := ctx.MustGet(AuthorizationClaimsKey).(*token.FireBaseCustomToken)
	if request.Password == "" {
		var err error
		request.Password, err = util.HashPassword(request.Password)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, errorResponse(err))
			return
		}
	}

	args := db.CreateAccountTxParams{
		Accounts: db.Accounts{
			UserID:   request.UserID,
			Username: request.Username,
			Icon: sql.NullString{
				String: request.Icon,
				Valid:  true,
			},
			ExplanatoryText: sql.NullString{
				String: request.ExplanatoryText,
				Valid:  true,
			},
			LocateID: request.LocateID,
			Rate:     0,
			HashedPassword: sql.NullString{
				String: request.Password,
				Valid:  true,
			},
			Email:      payload.Email,
			ShowLocate: request.ShowLocate,
			ShowRate:   request.ShowRate,
		},
		AccountTechTag:      request.TechTags,
		AccountFrameworkTag: request.Frameworks,
	}

	result, err := server.store.CreateAccountTx(ctx, args)
	if err != nil {
		// ToDo: すでに登録されている時のエラーの分岐を作る
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	locate, err := server.store.GetLocate(ctx, result.Account.LocateID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	response := AccountResponses{
		UserID:          result.Account.UserID,
		Username:        result.Account.Username,
		Icon:            result.Account.Icon.String,
		ExplanatoryText: result.Account.ExplanatoryText.String,
		Locate:          locate,
		Rate:            result.Account.Rate,
		ShowLocate:      result.Account.ShowLocate,
		ShowRate:        result.Account.ShowRate,
		TechTags:        result.AccountTechTags,
		Frameworks:      result.AccountFrameworks,
	}
	ctx.JSON(http.StatusOK, response)
}

// アカウントを取得する際のパラメータ
type GetAccountRequestParams struct {
	ID string `uri:"id"`
}

// アカウントを取得する
// 認証必須
func (server *Server) GetAccount(ctx *gin.Context) {
	var request GetAccountRequestParams
	if err := ctx.ShouldBindUri(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	payload := ctx.MustGet(AuthorizationClaimsKey).(*token.FireBaseCustomToken)

	// アカウント取得
	account, err := server.store.GetAccount(ctx, request.ID)
	if err != nil {
		// ToDo: IDがなかったときの分岐を作る
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
	}
	//
	locate, err := server.store.GetLocate(ctx, account.LocateID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	techTags, err := server.store.GetAccountTags(ctx, account.UserID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	frameworks, err := server.store.ListAccountFrameworks(ctx, account.UserID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	var accountTechTags []db.TechTags

	// 技術タグを取得する
	for _, tags := range techTags {
		techtag, err := server.store.GetTechTag(ctx, tags.TechTagID.Int32)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, errorResponse(err))
			return
		}
		accountTechTags = append(accountTechTags, techtag)
	}
	// フレームワークを取得する
	var accountFrameworks []db.Frameworks
	for _, framework := range frameworks {
		fw, err := server.store.GetFrameworks(ctx, framework.FrameworkID.Int32)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, errorResponse(err))
			return
		}
		accountFrameworks = append(accountFrameworks, fw)
	}

	// 本人のリクエストの時は、すべての情報を返す
	// そうでないときはShowLocateとShowRateの情報に沿って返す
	var response AccountResponses
	if payload.Email == account.Email {
		response = AccountResponses{
			UserID:          account.UserID,
			Username:        account.Username,
			Icon:            account.Icon.String,
			ExplanatoryText: account.ExplanatoryText.String,
			Locate:          locate,
			Rate:            account.Rate,
			ShowLocate:      account.ShowLocate,
			ShowRate:        account.ShowRate,
			TechTags:        accountTechTags,
			Frameworks:      accountFrameworks,
		}
	} else {
		response = AccountResponses{
			UserID:          account.UserID,
			Username:        account.Username,
			Icon:            account.Icon.String,
			ExplanatoryText: account.ExplanatoryText.String,
			ShowLocate:      account.ShowLocate,
			ShowRate:        account.ShowRate,
			TechTags:        accountTechTags,
			Frameworks:      accountFrameworks,
		}
		if account.ShowLocate {
			response.Locate = locate
		}
		if account.ShowRate {
			response.Rate = account.Rate
		}
	}
	ctx.JSON(http.StatusOK, response)
}
