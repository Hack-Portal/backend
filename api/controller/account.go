package controller

import (
	"bytes"
	"database/sql"
	"errors"
	"io"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/hackhack-Geek-vol6/backend/bootstrap"
	db "github.com/hackhack-Geek-vol6/backend/db/sqlc"
	"github.com/hackhack-Geek-vol6/backend/util"
	"github.com/hackhack-Geek-vol6/backend/util/token"
	"github.com/lib/pq"
)

type AccountController struct {
	AccountUsecase domain.LoginUsecase
	Env            *bootstrap.Env
}

// アカウントを作る時のリクエストパラメータ
type CreateAccountRequestBody struct {
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
type CreateAccountResponses struct {
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

// CreateAccount	godoc
// @Summary			Create new account
// @Description		Create an account from the requested body
// @Tags			Accounts
// @Produce			json
// @Param			CreateAccountRequestBody 	body 			CreateAccountRequestBody	true	"Create Account Request Body"
// @Success			200							{object}		CreateAccountResponses		"create succsss response"
// @Failure 		400							{object}		ErrorResponse				"bad request response"
// @Failure 		500							{object}		ErrorResponse				"server error response"
// @Router       	/accounts 	[post]
func (server *Server) CreateAccount(ctx *gin.Context) {
	var request CreateAccountRequestBody
	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	//
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
	response := CreateAccountResponses{
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
type AccountRequestWildCard struct {
	ID string `uri:"id"`
}

type GetAccountResponses struct {
	UserID          string `json:"user_id"`
	Username        string `json:"username"`
	Icon            string `json:"icon"`
	ExplanatoryText string `json:"explanatory_text"`
	Rate            int32  `json:"rate"`
	Email           string `json:"email"`
	Locate          string `json:"locate"`
	ShowLocate      bool   `json:"show_locate"`
	ShowRate        bool   `json:"show_rate"`

	TechTags   []db.TechTags   `json:"tech_tags"`
	Frameworks []db.Frameworks `json:"frameworks"`
}

// GetAccount		godoc
// @Summary			Get account
// @Description		Return a user from the id specified in the path
// @Tags			Accounts
// @Produce			json
// @Param			user_id 	path			string				true	"Accounts API wildcard"
// @Success			200			{object}		GetAccountResponses	"Get success response"
// @Failure 		400			{object}		ErrorResponse		"bad request response"
// @Failure 		500			{object}		ErrorResponse		"server error response"
// @Router       	/accounts/:user_id 			[get]
func (server *Server) GetAccount(ctx *gin.Context) {
	var request AccountRequestWildCard
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

	locate, err := server.store.GetLocateByID(ctx, account.LocateID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
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
	var response GetAccountResponses
	if payload.Email == account.Email {
		response = GetAccountResponses{
			UserID:          account.UserID,
			Username:        account.Username,
			Icon:            account.Icon.String,
			ExplanatoryText: account.ExplanatoryText.String,
			Locate:          locate.Name,
			Email:           account.Email,
			Rate:            account.Rate,
			ShowLocate:      account.ShowLocate,
			ShowRate:        account.ShowRate,
			TechTags:        accountTechTags,
			Frameworks:      accountFrameworks,
		}
	} else {
		response = GetAccountResponses{
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
			response.Locate = locate.Name
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
	Icon            string    `json:"icon"`
	Locate          string    `json:"locate"`
	Rate            int32     `json:"rate"`
	HashedPassword  string    `json:"hashed_password"`
	ShowLocate      bool      `json:"show_locate"`
	ShowRate        bool      `json:"show_rate"`
	CreatedAt       time.Time `json:"created_at"`
}

// UpdateAccount	godoc
// @Summary			Update Account
// @Description		Update user info from requested body
// @Tags			Accounts
// @Produce			json
// @Param			user_id 					path		string						true	"Accounts API wildcard"
// @Param			UpdateAccountRequestBody 	body		UpdateAccountRequestBody	true	"Update Account Request Body"
// @Success			200							{object}	UpdateAccountResponse		"Update success response"
// @Failure 		400							{object}	ErrorResponse				"bad request response"
// @Failure 		500							{object}	ErrorResponse				"server error response"
// @Router       	/accounts/:user_id 			[put]
func (server *Server) UpdateAccount(ctx *gin.Context) {
	var (
		requestBody UpdateAccountRequestBody
		requestURI  AccountRequestWildCard
		imageURL    string
	)
	if err := ctx.ShouldBindUri(&requestURI); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	if err := ctx.ShouldBindJSON(&requestBody); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	file, _, err := ctx.Request.FormFile(ImageKey)
	if err != nil {
		switch err.Error() {
		case MultiPartNextPartEoF:
			ctx.JSON(400, errorResponse(err))
			return
		case HttpNoSuchFile:
			ctx.JSON(400, errorResponse(err))
			return
		case RequestContentTypeIsnt:
			break
		default:
			ctx.JSON(400, errorResponse(err))
			return
		}
	} else {
		// 画像がある場合
		buf := bytes.NewBuffer(nil)
		if _, err := io.Copy(buf, file); err != nil {
			ctx.JSON(500, errorResponse(err))
			return
		}
		imageURL, err = server.store.UploadImage(ctx, buf.Bytes())
		if err != nil {
			ctx.JSON(500, errorResponse(err))
			return
		}
	}

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
	result, err := server.store.UpdateAccount(ctx, parseUpdateAccountParam(account, requestBody))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	locate, err := server.store.GetLocateByID(ctx, result.LocateID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	response := UpdateAccountResponse{
		Username:        result.Username,
		ExplanatoryText: requestBody.ExplanatoryText,
		Icon:            imageURL,
		Locate:          locate.Name,
		Rate:            result.Rate,
		HashedPassword:  result.HashedPassword.String,
		ShowLocate:      result.ShowLocate,
		ShowRate:        result.ShowRate,
	}

	ctx.JSON(http.StatusOK, response)
}

// 型変換
func parseUpdateAccountParam(account db.GetAccountByEmailRow, body UpdateAccountRequestBody) (result db.UpdateAccountParams) {
	result.UserID = account.UserID
	if util.StringLength(body.Username) != 0 {
		if util.EqualString(account.Username, body.Username) {
			result.Username = account.Username
		}
		result.Username = body.Username
	}

	if !util.EqualString(account.ExplanatoryText.String, body.ExplanatoryText) {
		if util.StringLength(body.ExplanatoryText) != 0 {
			result.ExplanatoryText = sql.NullString{
				String: body.ExplanatoryText,
				Valid:  true,
			}
		} else {
			result.ExplanatoryText = sql.NullString{
				Valid: false,
			}
		}

	} else {
		result.ExplanatoryText = account.ExplanatoryText
	}

	if body.LocateID != 0 {
		if util.Equalint(int(account.LocateID), int(body.LocateID)) {
			result.LocateID = account.LocateID
		} else {
			result.LocateID = body.LocateID
		}
	}

	if body.Rate != 0 {
		if util.Equalint(int(account.Rate), int(body.Rate)) {
			result.Rate = account.Rate
		} else {
			result.Rate = body.Rate
		}
	}

	if util.StringLength(body.HashedPassword) != 0 {
		var hashedPassword string
		hashedPassword, _ = util.HashPassword(body.HashedPassword)
		if util.EqualString(account.HashedPassword.String, hashedPassword) {
			result.HashedPassword = account.HashedPassword
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

// DeleteAccount	godoc
// @Summary			Remove Account
// @Description		Only you can delete your account (logical delete)
// @Tags			Accounts
// @Produce			json
// @Param			user_id 	path			string			true	"Accounts API wildcard"
// @Success			200			{object}		DeleteResponse	"delete success response"
// @Failure 		400			{object}		ErrorResponse	"bad request response"
// @Failure 		500			{object}		ErrorResponse	"server error response"
// @Router       	/accounts/:user_id 		[delete]
func (server *Server) DeleteAccount(ctx *gin.Context) {
	var request AccountRequestWildCard
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
	ctx.JSON(http.StatusOK, DeleteResponse{
		Result: "success",
	})

}
