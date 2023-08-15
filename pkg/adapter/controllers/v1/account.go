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
	"github.com/hackhack-Geek-vol6/backend/pkg/util/jwt"
	"github.com/lib/pq"
)

type AccountController struct {
	AccountUsecase inputport.AccountUsecase
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
	var (
		reqBody domain.CreateAccountRequest
		image   []byte
	)
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
		icon := bytes.NewBuffer(nil)
		if _, err := io.Copy(icon, file); err != nil {
			ctx.JSON(http.StatusInternalServerError, errorResponse(err))
			return
		}
		image = icon.Bytes()
	}

	payload := ctx.MustGet(middleware.AuthorizationClaimsKey).(*jwt.FireBaseCustomToken)
	// ここでエラーが起きてる

	response, err := ac.AccountUsecase.CreateAccount(ctx, reqBody, image, payload.Email)

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

	response, err := ac.AccountUsecase.GetAccountByID(ctx, reqUri.AccountID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, response)
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
		reqBody domain.UpdateAccountRequest
		reqURI  domain.AccountRequestWildCard
		image   []byte
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
		icon := bytes.NewBuffer(nil)
		if _, err := io.Copy(icon, file); err != nil {
			ctx.JSON(http.StatusInternalServerError, errorResponse(err))
			return
		}
		image = icon.Bytes()
	}

	response, err := ac.AccountUsecase.UpdateAccount(
		ctx,
		domain.UpdateAccountParam{
			AccountInfo: repository.Account{
				AccountID: reqURI.AccountID,
				UserID:    reqBody.UserID,
				Username:  reqBody.Username,
				ExplanatoryText: sql.NullString{
					String: reqBody.ExplanatoryText,
					Valid:  true,
				},
				LocateID:   reqBody.LocateID,
				ShowLocate: reqBody.ShowLocate,
				ShowRate:   reqBody.ShowRate,
			},
			AccountTechTag:      reqBody.TechTags,
			AccountFrameworkTag: reqBody.Frameworks,
		},
		image)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, response)
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
