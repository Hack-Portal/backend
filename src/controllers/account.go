package controller

import (
	"bytes"
	"database/sql"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hackhack-Geek-vol6/backend/pkg/jwt"
	"github.com/hackhack-Geek-vol6/backend/pkg/logger"
	"github.com/hackhack-Geek-vol6/backend/pkg/utils"
	"github.com/hackhack-Geek-vol6/backend/src/domain/params"
	"github.com/hackhack-Geek-vol6/backend/src/domain/request"
	"github.com/hackhack-Geek-vol6/backend/src/infrastructure/middleware"
	"github.com/hackhack-Geek-vol6/backend/src/repository"
	"github.com/hackhack-Geek-vol6/backend/src/usecases/inputport"
	usecase "github.com/hackhack-Geek-vol6/backend/src/usecases/interactor"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/newrelic/go-agent/v3/integrations/nrgin"
)

type AccountController struct {
	AccountUsecase inputport.AccountUsecase
	l              logger.Logger
}

func NewAccountController(store repository.SQLStore, l logger.Logger) *AccountController {
	return &AccountController{
		AccountUsecase: usecase.NewAccountUsercase(store, l),
		l:              l,
	}
}

// CreateAccount	godoc
//
//	@Summary		Create new account
//	@Description	Create an account from the requested body
//	@Accept			multipart/form-data
//	@Tags			Accounts
//	@Produce		json
//	@Param			CreateAccountRequest	body		request.CreateAccount	true	"Create Account Request"
//	@Success		200						{object}	response.Account		"create success response"
//	@Failure		400						{object}	ErrorResponse			"bad request response"
//	@Failure		500						{object}	ErrorResponse			"server error response"
//	@Router			/accounts				[post]
func (ac *AccountController) CreateAccount(ctx *gin.Context) {
	txn := nrgin.Transaction(ctx)
	defer txn.End()

	var (
		reqBody    request.CreateAccount
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
		tags, err = utils.StringToArrayInt32(reqBody.TechTags)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, errorResponse(err))
			return
		}
	}

	if len(reqBody.Frameworks) != 0 {
		frameworks, err = utils.StringToArrayInt32(reqBody.Frameworks)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, errorResponse(err))
			return
		}
	}

	payload := ctx.MustGet(middleware.AuthorizationClaimsKey).(*jwt.FireBaseCustomToken)

	response, err := ac.AccountUsecase.CreateAccount(ctx, params.CreateAccount{
		AccountInfo: repository.CreateAccountsParams{
			AccountID:       reqBody.AccountID,
			Email:           payload.Email,
			Username:        reqBody.Username,
			LocateID:        reqBody.LocateID,
			ExplanatoryText: utils.ToPgText(reqBody.ExplanatoryText),
			ShowLocate:      reqBody.ShowLocate,
			ShowRate:        reqBody.ShowRate,
		},
		AccountTechTag:      tags,
		AccountFrameworkTag: frameworks,
	}, image)

	if err != nil {
		// すでに登録されている場合と参照エラーの処理
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
//	@Param			account_id				path		string				true	"Accounts API wildcard"
//	@Success		200						{object}	response.Account	"Get success response"
//	@Failure		400						{object}	ErrorResponse		"bad request response"
//	@Failure		403						{object}	ErrorResponse		"error response"
//	@Failure		500						{object}	ErrorResponse		"server error response"
//	@Router			/accounts/{account_id}	[get]
func (ac *AccountController) GetAccount(ctx *gin.Context) {
	txn := nrgin.Transaction(ctx)
	defer txn.End()
	var (
		payload *jwt.FireBaseCustomToken
		reqUri  request.AccountWildCard
	)

	if err := ctx.ShouldBindUri(&reqUri); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	if !ctx.MustGet(middleware.AuthorizationKeyNotInclude).(bool) {
		payload = ctx.MustGet(middleware.AuthorizationClaimsKey).(*jwt.FireBaseCustomToken)
	}

	log.Println(payload)
	log.Println(reqUri.AccountID)

	response, err := ac.AccountUsecase.GetAccountByID(ctx, reqUri.AccountID, payload)
	if err != nil {
		switch err.Error() {
		case sql.ErrNoRows.Error():
			err := errors.New("そんなユーザおらんがな")
			ctx.JSON(http.StatusForbidden, errorResponse(err))
		default:
			err := errors.New("すまんサーバエラーや")
			ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		}
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
//	@Param			account_id				path		string					true	"Accounts API wildcard"
//	@Param			UpdateAccountRequest	body		request.UpdateAccount	true	"Update Account Request Body"
//	@Success		200						{object}	response.Account		"Update success response"
//	@Failure		400						{object}	ErrorResponse			"bad request response"
//	@Failure		403						{object}	ErrorResponse			"error response"
//	@Failure		500						{object}	ErrorResponse			"server error response"
//	@Router			/accounts/{account_id} [put]
func (ac *AccountController) UpdateAccount(ctx *gin.Context) {
	txn := nrgin.Transaction(ctx)
	defer txn.End()
	var (
		reqBody    request.UpdateAccount
		reqURI     request.AccountWildCard
		image      []byte
		tags       []int32
		frameworks []int32
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

	if len(reqBody.TechTags) != 0 {
		tags, err = utils.StringToArrayInt32(reqBody.TechTags)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, errorResponse(err))
			return
		}
	}

	if len(reqBody.Frameworks) != 0 {
		frameworks, err = utils.StringToArrayInt32(reqBody.Frameworks)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, errorResponse(err))
			return
		}
	}

	response, err := ac.AccountUsecase.UpdateAccount(
		ctx,
		params.UpdateAccount{
			AccountInfo: repository.Account{
				AccountID: reqURI.AccountID,
				Username:  reqBody.Username,
				ExplanatoryText: pgtype.Text{
					String: reqBody.ExplanatoryText,
					Valid:  true,
				},
				LocateID:    reqBody.LocateID,
				ShowLocate:  reqBody.ShowLocate,
				ShowRate:    reqBody.ShowRate,
				TwitterLink: utils.ToPgText(reqBody.TwitterLink),
				GithubLink:  utils.ToPgText(reqBody.GithubLink),
				DiscordLink: utils.ToPgText(reqBody.DiscordLink),
			},
			AccountTechTag:      tags,
			AccountFrameworkTag: frameworks,
		},
		image)
	if err != nil {
		switch err.Error() {
		case sql.ErrNoRows.Error():
			err := errors.New("そんなユーザおらんがな")
			ctx.JSON(http.StatusForbidden, errorResponse(err))
		default:
			err := errors.New(fmt.Sprintf("すまんサーバエラーや%s", err))
			ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		}
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
//	@Failure		403						{object}	ErrorResponse	"error response"
//	@Failure		500						{object}	ErrorResponse	"server error response"
//	@Router			/accounts/{account_id}	[delete]
func (ac *AccountController) DeleteAccount(ctx *gin.Context) {
	txn := nrgin.Transaction(ctx)
	defer txn.End()
	var reqURI request.AccountWildCard
	if err := ctx.ShouldBindUri(&reqURI); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	err := ac.AccountUsecase.DeleteAccount(ctx, reqURI.AccountID)

	if err != nil {
		switch err.Error() {
		case sql.ErrNoRows.Error():
			err := errors.New("そんなユーザおらんがな")
			ctx.JSON(http.StatusForbidden, errorResponse(err))
		default:
			err := errors.New("すまんサーバエラーや")
			ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		}
		return
	}

	ctx.JSON(http.StatusOK, SuccessResponse{Result: fmt.Sprintf("delete successful")})
}

// GetJoinRoom	godoc
//
//	@Summary		Get Join Room
//	@Description	Get Join Room
//	@Tags			Accounts
//	@Produce		json
//	@Success		200								{array}		response.GetJoinRoom	"success response"
//	@Failure		400								{object}	ErrorResponse			"error response"
//	@Failure		403								{object}	ErrorResponse			"error response"
//	@Failure		500								{object}	ErrorResponse			"error response"
//	@Router			/accounts/{account_id}/rooms	[get]
func (ac *AccountController) GetJoinRoom(ctx *gin.Context) {
	txn := nrgin.Transaction(ctx)
	defer txn.End()

	var reqURI request.AccountWildCard
	if err := ctx.ShouldBindUri(&reqURI); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	response, err := ac.AccountUsecase.GetJoinRoom(ctx, reqURI.AccountID)
	if err != nil {
		switch err.Error() {
		case sql.ErrNoRows.Error():
			err := errors.New("そんなユーザおらんがな")
			ctx.JSON(http.StatusForbidden, errorResponse(err))
		default:
			err := errors.New("すまんサーバエラーや")
			ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		}
		return
	}

	ctx.JSON(http.StatusOK, response)
}
