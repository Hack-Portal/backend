package controller

import (
	"bytes"
	"database/sql"
	"errors"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hackhack-Geek-vol6/backend/api/middleware"
	"github.com/hackhack-Geek-vol6/backend/bootstrap"
	db "github.com/hackhack-Geek-vol6/backend/db/sqlc"
	"github.com/hackhack-Geek-vol6/backend/domain"
	"github.com/hackhack-Geek-vol6/backend/util"
	"github.com/hackhack-Geek-vol6/backend/util/token"
	"github.com/lib/pq"
)

type AccountController struct {
	AccountUsecase domain.AccountRepository
	Env            *bootstrap.Env
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
func (ac *AccountController) CreateAccount(ctx *gin.Context) {
	var request domain.CreateAccountRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	result, err := ac.AccountUsecase.CreateAccount(ctx, db.CreateAccountTxParams{
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
			ShowLocate: request.ShowLocate,
			ShowRate:   request.ShowRate,
		},
		AccountTechTag:      request.TechTags,
		AccountFrameworkTag: request.Frameworks,
	})

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
	ctx.JSON(http.StatusOK, result)
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
func (ac *AccountController) GetAccount(ctx *gin.Context) {
	var reqUri domain.AccountRequestWildCard
	if err := ctx.ShouldBindUri(&reqUri); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	// アカウント取得
	result, err := ac.AccountUsecase.GetAccountByID(ctx, reqUri.ID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, result)
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
func (ac *AccountController) UpdateAccount(ctx *gin.Context) {
	var (
		reqBody  domain.UpdateAccountRequest
		reqURI   domain.AccountRequestWildCard
		imageURL string
	)
	if err := ctx.ShouldBindUri(&reqURI); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	if err := ctx.ShouldBindJSON(&reqBody); err != nil {
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
		imageURL, err = ac.AccountUsecase.UploadImage(ctx, buf.Bytes())
		if err != nil {
			ctx.JSON(500, errorResponse(err))
			return
		}
	}
	account, err := ac.AccountUsecase.GetAccountByID(ctx, reqURI.ID)

	result, err := ac.AccountUsecase.UpdateAccount(ctx, parseUpdateAccountParam(account, reqBody))
	result.Icon = imageURL

	ctx.JSON(http.StatusOK, result)
}

// 型変換
func parseUpdateAccountParam(account db.GetAccountByEmailRow, body domain.UpdateAccountRequest) (result db.UpdateAccountParams) {
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
func (ac *AccountController) DeleteAccount(ctx *gin.Context) {
	var request AccountRequestWildCard
	if err := ctx.ShouldBindUri(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	payload := ctx.MustGet(middleware.AuthorizationClaimsKey).(*token.FireBaseCustomToken)
	account, err := ac.AccountUsecase.GetAccountByEmail(ctx, payload.Email)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	if account.UserID != request.ID {
		err := errors.New("valid user id")
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	err = ac.AccountUsecase.DeleteAccount(ctx, request.ID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, DeleteResponse{
		Result: "success",
	})

}
