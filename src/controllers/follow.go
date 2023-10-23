package controller

import (
	"database/sql"
	"errors"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hackhack-Geek-vol6/backend/pkg/logger"
	"github.com/hackhack-Geek-vol6/backend/pkg/repository"
	"github.com/hackhack-Geek-vol6/backend/src/domain/request"
	"github.com/hackhack-Geek-vol6/backend/src/transaction"
	"github.com/hackhack-Geek-vol6/backend/src/usecases/inputport"
	usecase "github.com/hackhack-Geek-vol6/backend/src/usecases/interactor"
	"github.com/newrelic/go-agent/v3/integrations/nrgin"
)

type FollowController struct {
	FollowUsecase inputport.FollowUsecase
	l             logger.Logger
}

func NewFollowController(store transaction.SQLStore, l logger.Logger) *FollowController {
	return &FollowController{
		FollowUsecase: usecase.NewFollowUsercase(store, l),
		l:             l,
	}
}

// TODO:レスポンス変更 => accounts
// CreateFollow	godoc
//
//	@Summary		Create Follow
//	@Description	Follow!!!!!!!!
//	@Tags			Accounts
//	@Produce		json
//	@Param			from_account_id						path		string					true	"Accounts API wildcard"
//	@Param			CreateFollowRequest					body		request.CreateFollow	true	"create Follow Request Body"
//	@Success		200									{array}		response.Follow			"success response"
//	@Failure		400									{object}	ErrorResponse			"error response"
//	@Failure		403									{object}	ErrorResponse			"error response"
//	@Failure		500									{object}	ErrorResponse			"error response"
//	@Router			/accounts/{from_account_id}/follow	[post]
func (fc *FollowController) CreateFollow(ctx *gin.Context) {
	txn := nrgin.Transaction(ctx)
	defer txn.End()
	var (
		reqURI  request.AccountWildCard
		reqBody request.CreateFollow
	)
	if err := ctx.ShouldBindUri(&reqURI); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	if err := ctx.ShouldBindJSON(&reqBody); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	response, err := fc.FollowUsecase.CreateFollow(ctx, repository.CreateFollowsParams{
		ToAccountID:   reqBody.ToAccountID,
		FromAccountID: reqURI.AccountID,
	})
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

// TODO:レスポンス修正
// RemoveFollow	godoc
//
//	@Summary		Remove follow
//	@Description	Remove follow account
//	@Tags			Accounts
//	@Produce		json
//	@Param			from_account_id						path		string					true	"Accounts API wildcard"
//	@Param			RemoveFollowRequestQueries			query		request.RemoveFollow	true	"Remove Follow Request Body"
//	@Success		200									{object}	SuccessResponse			"success response"
//	@Failure		400									{object}	ErrorResponse			"error response"
//	@Failure		500									{object}	ErrorResponse			"error response"
//	@Router			/accounts/{from_account_id}/follow	[delete]
func (fc *FollowController) RemoveFollow(ctx *gin.Context) {
	txn := nrgin.Transaction(ctx)
	defer txn.End()
	var (
		reqURI   request.AccountWildCard
		reqQuery request.RemoveFollow
	)
	if err := ctx.ShouldBindUri(&reqURI); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	if err := ctx.ShouldBindQuery(&reqQuery); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	fmt.Println("query", reqQuery)

	if err := fc.FollowUsecase.RemoveFollow(ctx, repository.DeleteFollowsParams{ToAccountID: reqQuery.AccountID, FromAccountID: reqURI.AccountID}); err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, SuccessResponse{Result: "Delete Successful"})
}

// GetFollow	godoc
//
//	@Summary		Get follow
//	@Description	Get follow account
//	@Tags			Accounts
//	@Produce		json
//	@Param			from_account_id						path		string				true	"Accounts API wildcard"
//	@Param			GetFollowRequestQueries				query		request.GetFollow	true	"Get Follow Request Body"
//	@Success		200									{object}	[]response.Follow	"success response"
//	@Failure		400									{object}	ErrorResponse		"error response"
//	@Failure		403									{object}	ErrorResponse		"error response"
//	@Failure		500									{object}	ErrorResponse		"error response"
//	@Router			/accounts/{from_account_id}/follow	[get]
func (fc *FollowController) GetFollow(ctx *gin.Context) {
	txn := nrgin.Transaction(ctx)
	defer txn.End()
	var (
		reqURI   request.AccountWildCard
		reqQuery request.GetFollow
	)
	if err := ctx.ShouldBindUri(&reqURI); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	if err := ctx.ShouldBindQuery(&reqQuery); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	result, err := fc.FollowUsecase.GetFollowByID(ctx, reqURI.AccountID, reqQuery.Mode)
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
	ctx.JSON(http.StatusOK, result)
}
