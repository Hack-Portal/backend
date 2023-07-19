package api

import (
	"database/sql"
	"errors"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	db "github.com/hackhack-Geek-vol6/backend/db/sqlc"
	"github.com/hackhack-Geek-vol6/backend/util"
	"github.com/hackhack-Geek-vol6/backend/util/token"
	"github.com/lib/pq"
)

// アカウントを作る時のリクエストパラメータ
type CreateAccountRequestParam struct {
	UserID          string  `json:"user_id" binding:"required"`
	Username        string  `json:"username" binding:"required"`
	Icon            string  `json:"icon"`
	ExplanatoryText string  `json:"explanatory_text"`
	LocateID        int32   `json:"locate_id" binding:"required"`
	Password        string  `json:"password"`
	ShowLocate      bool    `json:"show_locate" binding:"required"`
	ShowRate        bool    `json:"show_rate" binding:"required"`
	TechTags        []int32 `json:"tech_tags"`
	Frameworks      []int32 `json:"frameworks"`
}

// アカウントに関するレスポンス
type AccountResponses struct {
	UserID          string `json:"user_id"`
	Username        string `json:"username"`
	Icon            string `json:"icon"`
	ExplanatoryText string `json:"explanatory_text"`
	Rate            int32  `json:"rate"`
	Locate          string `json:"locate"`
	ShowLocate      bool   `json:"show_locate"`
	ShowRate        bool   `json:"show_rate"`

	TechTags   []db.TechTags   `json:"tech_tags"`
	Frameworks []db.Frameworks `json:"frameworks"`
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
		// すでに登録されている場合と参照エラーの処理
		if pqErr, ok := err.(*pq.Error); ok {
			switch pqErr.Code.Name() {
			case db.ForeignKeyViolation, db.UniqueViolation:
				ctx.JSON(http.StatusForbidden, errorResponse(err))
				return
			}
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	locate, err := server.store.GetLocateByID(ctx, result.Account.LocateID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	response := AccountResponses{
		UserID:          result.Account.UserID,
		Username:        result.Account.Username,
		Icon:            result.Account.Icon.String,
		ExplanatoryText: result.Account.ExplanatoryText.String,
		Locate:          locate.Name,
		Rate:            result.Account.Rate,
		ShowLocate:      result.Account.ShowLocate,
		ShowRate:        result.Account.ShowRate,
		TechTags:        result.AccountTechTags,
		Frameworks:      result.AccountFrameworks,
	}
	ctx.JSON(http.StatusOK, response)
}

// アカウントを取得する際のパラメータ
type RequestParams struct {
	ID string `uri:"id"`
}

// アカウントを取得する
// 認証必須
func (server *Server) GetAccount(ctx *gin.Context) {
	var request RequestParams
	if err := ctx.ShouldBindUri(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	payload := ctx.MustGet(AuthorizationClaimsKey).(*token.FireBaseCustomToken)

	// アカウント取得
	account, err := server.store.GetAccountByID(ctx, request.ID)
	if err != nil {
		// ToDo: IDがなかったときの分岐を作る
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
	}

	techTags, err := server.store.ListAccountTagsByUserID(ctx, account.UserID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	frameworks, err := server.store.ListAccountFrameworksByUserID(ctx, account.UserID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	var accountTechTags []db.TechTags

	// 技術タグを取得する
	for _, tags := range techTags {
		techTag, err := server.store.GetTechTagByID(ctx, tags.TechTagID.Int32)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, errorResponse(err))
			return
		}
		accountTechTags = append(accountTechTags, techTag)
	}
	// フレームワークを取得する
	var accountFrameworks []db.Frameworks
	for _, framework := range frameworks {
		fw, err := server.store.GetFrameworksByID(ctx, framework.FrameworkID.Int32)
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
			Locate:          account.Locate,
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
			response.Locate = account.Locate
		}
		if account.ShowRate {
			response.Rate = account.Rate
		}
	}
	ctx.JSON(http.StatusOK, response)
}

type UpdateAccountRequestBody struct {
	Username        string `json:"username"`
	ExplanatoryText string `json:"explanatory_text"`
	LocateID        int32  `json:"locate_id"`
	Rate            int32  `json:"rate"`
	HashedPassword  string `json:"hashed_password"`
	ShowLocate      bool   `json:"show_locate"`
	ShowRate        bool   `json:"show_rate"`
}

type UpdateAccountResponse struct {
	Username        string    `json:"username"`
	ExplanatoryText string    `json:"explanatory_text"`
	icon            string    `json:"icon"`
	Locate          string    `json:"locate"`
	Rate            int32     `json:"rate"`
	HashedPassword  string    `json:"hashed_password"`
	ShowLocate      bool      `json:"show_locate"`
	ShowRate        bool      `json:"show_rate"`
	CreatedAt       time.Time `json:"created_at"`
}

// アカウント更新
// 認証必須
func (server *Server) UpdateAccount(ctx *gin.Context) {
	var (
		requestBody UpdateAccountRequestBody
		requestURI  RequestParams
	)
	if err := ctx.ShouldBindUri(&requestURI); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	if err := ctx.ShouldBindJSON(&requestBody); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	// TODO:条件次第でFormDataから画像を読み込む

	payload := ctx.MustGet(AuthorizationClaimsKey).(*token.FireBaseCustomToken)
	account, err := server.store.GetAccountByEmail(ctx, payload.Email)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	if account.UserID != requestURI.ID {
		err := errors.New("valid user id")
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	arg, err := parseUpdateAccountParam(account, requestBody)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	result, err := server.store.UpdateAccount(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	locate, err := server.store.GetLocateByID(ctx, result.LocateID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	// TODO:iconのPathを含む
	response := UpdateAccountResponse{
		Username:        result.Username,
		ExplanatoryText: requestBody.ExplanatoryText,
		icon:            result.Icon.String,
		Locate:          locate.Name,
		Rate:            result.Rate,
		HashedPassword:  result.HashedPassword.String,
		ShowLocate:      result.ShowLocate,
		ShowRate:        result.ShowRate,
	}

	ctx.JSON(http.StatusOK, response)
}

// 型変換
func parseUpdateAccountParam(account db.GetAccountByEmailRow, body UpdateAccountRequestBody) (result db.UpdateAccountParams, err error) {
	result.UserID = account.UserID
	if len(strings.TrimSpace(body.Username)) != 0 {
		result.Username = body.Username
	}

	if len(strings.TrimSpace(body.ExplanatoryText)) != 0 {
		result.ExplanatoryText = sql.NullString{
			String: body.ExplanatoryText,
			Valid:  true,
		}
	} else {
		result.ExplanatoryText = sql.NullString{
			Valid: false,
		}
	}
	if body.LocateID != 0 {
		result.LocateID = body.LocateID
	}

	if body.Rate != 0 {
		result.Rate = body.Rate
	}

	if len(strings.TrimSpace(body.HashedPassword)) != 0 {
		var hashedPassword string
		hashedPassword, err = util.HashPassword(body.HashedPassword)
		if err != nil {
			return
		}

		result.HashedPassword = sql.NullString{
			String: hashedPassword,
			Valid:  true,
		}
	} else {
		result.HashedPassword = sql.NullString{
			Valid: false,
		}
	}

	result.ShowLocate = body.ShowLocate
	result.ShowRate = body.ShowRate

	return
}

// ユーザ削除
func (server *Server) DeleteAccount(ctx *gin.Context) {
	var request RequestParams
	if err := ctx.ShouldBindUri(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	payload := ctx.MustGet(AuthorizationClaimsKey).(*token.FireBaseCustomToken)
	account, err := server.store.GetAccountByEmail(ctx, payload.Email)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	if account.UserID != request.ID {
		err := errors.New("valid user id")
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	_, err = server.store.SoftDeleteAccount(ctx, request.ID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"result": "success",
	})

}
