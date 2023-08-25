package controller

import (
	"bytes"
	"database/sql"
	"fmt"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hackhack-Geek-vol6/backend/pkg/adapter/gateways/infrastructure/httpserver/middleware"
	repository "github.com/hackhack-Geek-vol6/backend/pkg/adapter/gateways/repository/datasource"
	"github.com/hackhack-Geek-vol6/backend/pkg/adapter/gateways/repository/transaction"
	"github.com/hackhack-Geek-vol6/backend/pkg/bootstrap"
	"github.com/hackhack-Geek-vol6/backend/pkg/domain"
	"github.com/hackhack-Geek-vol6/backend/pkg/usecase/inputport"
	dbutil "github.com/hackhack-Geek-vol6/backend/pkg/util/db"
	util "github.com/hackhack-Geek-vol6/backend/pkg/util/etc"
	"github.com/hackhack-Geek-vol6/backend/pkg/util/jwt"
	"github.com/lib/pq"
	"github.com/newrelic/go-agent/v3/integrations/nrgin"
)

type AccountController struct {
	AccountUsecase inputport.AccountUsecase
	Env            *bootstrap.Env
}

// CreateAccount	godoc
//
//	@Summary		Create new account
//	@Description	Create an account from the requested body
//	@Accept			multipart/form-data
//	@Tags			Accounts
//	@Produce		json
//	@Param			CreateAccountRequest	body		domain.CreateAccountRequest	true	"Create Account Request"
//	@Success		200						{object}	domain.AccountResponses		"create success response"
//	@Failure		400						{object}	ErrorResponse				"bad request response"
//	@Failure		500						{object}	ErrorResponse				"server error response"
//	@Router			/accounts 																										[post]
func (ac *AccountController) CreateAccount(ctx *gin.Context) {
	txn := nrgin.Transaction(ctx)
	defer txn.End()

	var (
		reqBody    domain.CreateAccountRequest
		image      []byte
		tags       []int32
		frameworks []int32
	)
	if err := ctx.ShouldBind(&reqBody); err != nil {
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
			break
		case RequestContentTypeIsnt:
			break
		default:
			ctx.JSON(400, errorResponse(err))
			return
		}
	} else {
		icon := bytes.NewBuffer(nil)
		if _, err := io.Copy(icon, file); err != nil {
			ctx.JSON(http.StatusInternalServerError, errorResponse(err))
			return
		}
		image = icon.Bytes()
	}
	if len(reqBody.TechTags) != 0 {
		tags, err = util.StringToArrayInt32(reqBody.TechTags)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, errorResponse(err))
			return
		}
	}

	if len(reqBody.Frameworks) != 0 {
		frameworks, err = util.StringToArrayInt32(reqBody.Frameworks)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, errorResponse(err))
			return
		}
	}

	payload := ctx.MustGet(middleware.AuthorizationClaimsKey).(*jwt.FireBaseCustomToken)

	response, err := ac.AccountUsecase.CreateAccount(ctx, domain.CreateAccount{
		ReqBody:    reqBody,
		TechTags:   tags,
		Frameworks: frameworks,
	}, image, payload.Email)

	if err != nil {
		// すでに登録されている場合と参照エラーの処理
		if pqErr, ok := err.(*pq.Error); ok {
			switch pqErr.Code.Name() {
			case transaction.ForeignKeyViolation, transaction.UniqueViolation:
				ctx.JSON(http.StatusForbidden, errorResponse(err))
				return
			}
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, response)
}

// GetAccount		godoc
//
//	@Summary		Get account
//	@Description	Return a account from the id specified in the path
//	@Tags			Accounts
//	@Produce		json
//	@Param			account_id				path		string					true	"Accounts API wildcard"
//	@Success		200						{object}	domain.AccountResponses	"Get success response"
//	@Failure		400						{object}	ErrorResponse			"bad request response"
//	@Failure		500						{object}	ErrorResponse			"server error response"
//	@Router			/accounts/{account_id} 																								[get]
func (ac *AccountController) GetAccount(ctx *gin.Context) {
	txn := nrgin.Transaction(ctx)
	defer txn.End()
	var reqUri domain.AccountRequestWildCard
	if err := ctx.ShouldBindUri(&reqUri); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	response, err := ac.AccountUsecase.GetAccountByID(ctx, reqUri.AccountID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, response)
}

// UpdateAccount	godoc
//
//	@Summary		Update Account
//	@Description	Update account info from requested body
//	@Tags			Accounts
//	@Produce		json
//	@Param			account_id				path		string						true	"Accounts API wildcard"
//	@Param			UpdateAccountRequest	body		domain.UpdateAccountRequest	true	"Update Account Request Body"
//	@Success		200						{object}	domain.AccountResponses		"Update success response"
//	@Failure		400						{object}	ErrorResponse				"bad request response"
//	@Failure		500						{object}	ErrorResponse				"server error response"
//	@Router			/accounts/{account_id} 																											[put]
func (ac *AccountController) UpdateAccount(ctx *gin.Context) {
	txn := nrgin.Transaction(ctx)
	defer txn.End()
	var (
		reqBody domain.UpdateAccountRequest
		reqURI  domain.AccountRequestWildCard
		image   []byte
	)
	if err := ctx.ShouldBindUri(&reqURI); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	if err := ctx.ShouldBind(&reqBody); err != nil {
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
			break
		case RequestContentTypeIsnt:
			break
		default:
			ctx.JSON(400, errorResponse(err))
			return
		}
	} else {
		icon := bytes.NewBuffer(nil)
		if _, err := io.Copy(icon, file); err != nil {
			ctx.JSON(http.StatusInternalServerError, errorResponse(err))
			return
		}
		image = icon.Bytes()
	}

	tags, err := util.StringToArrayInt32(reqBody.TechTags)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	frameworks, err := util.StringToArrayInt32(reqBody.Frameworks)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	response, err := ac.AccountUsecase.UpdateAccount(
		ctx,
		domain.UpdateAccountParam{
			AccountInfo: repository.Account{
				AccountID: reqURI.AccountID,
				Username:  reqBody.Username,
				ExplanatoryText: sql.NullString{
					String: reqBody.ExplanatoryText,
					Valid:  true,
				},
				LocateID:    reqBody.LocateID,
				ShowLocate:  reqBody.ShowLocate,
				ShowRate:    reqBody.ShowRate,
				TwitterLink: dbutil.ToSqlNullString(reqBody.TwitterLink),
				GithubLink:  dbutil.ToSqlNullString(reqBody.GithubLink),
				DiscordLink: dbutil.ToSqlNullString(reqBody.DiscordLink),
			},
			AccountTechTag:      tags,
			AccountFrameworkTag: frameworks,
		},
		image)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, response)
}

// DeleteAccount	godoc
//
//	@Summary		Remove Account
//	@Description	Only you can delete your account (logical delete)
//	@Tags			Accounts
//	@Produce		json
//	@Param			account_id				path		string			true	"Accounts API wildcard"
//	@Success		200						{object}	SuccessResponse	"delete success response"
//	@Failure		400						{object}	ErrorResponse	"bad request response"
//	@Failure		500						{object}	ErrorResponse	"server error response"
//	@Router			/accounts/{account_id} 														[delete]
func (ac *AccountController) DeleteAccount(ctx *gin.Context) {
	txn := nrgin.Transaction(ctx)
	defer txn.End()
	var reqURI domain.AccountRequestWildCard
	if err := ctx.ShouldBindUri(&reqURI); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	err := ac.AccountUsecase.DeleteAccount(ctx, reqURI.AccountID)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, SuccessResponse{Result: fmt.Sprintf("delete successful")})
}
