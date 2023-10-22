package usecase

import (
	"context"
	"time"

	"github.com/hackhack-Geek-vol6/backend/pkg/adapter/gateways/repository/transaction"
	"github.com/hackhack-Geek-vol6/backend/src/domain/response"
	"github.com/hackhack-Geek-vol6/backend/src/repository"
	"github.com/hackhack-Geek-vol6/backend/src/usecases/inputport"
)

type followUsecase struct {
	store          transaction.Store
	contextTimeout time.Duration
}

func NewFollowUsercase(store transaction.Store, timeout time.Duration) inputport.FollowUsecase {
	return &followUsecase{
		store:          store,
		contextTimeout: timeout,
	}
}

func (fu *followUsecase) CreateFollow(ctx context.Context, body repository.CreateFollowsParams) (result response.Follow, err error) {
	ctx, cancel := context.WithTimeout(ctx, fu.contextTimeout)
	defer cancel()

	follow, err := fu.store.CreateFollows(ctx, body)
	if err != nil {
		return
	}

	account, err := fu.store.GetAccountsByID(ctx, follow.ToAccountID)
	if err != nil {
		return
	}
	result = response.Follow{AccountID: account.AccountID, Username: account.Username, Icon: account.Icon.String}
	return
}

func (fu *followUsecase) RemoveFollow(ctx context.Context, body repository.DeleteFollowsParams) error {
	ctx, cancel := context.WithTimeout(ctx, fu.contextTimeout)
	defer cancel()

	return fu.store.DeleteFollows(ctx, body)
}

func (fu *followUsecase) GetFollowByID(ctx context.Context, ID string, mode bool) (result []response.Follow, err error) {
	ctx, cancel := context.WithTimeout(ctx, fu.contextTimeout)
	defer cancel()

	var follows []repository.Follow

	if mode {
		follows, err = fu.store.ListFollowsByToUserID(ctx, ID)
		if err != nil {
			return
		}
	} else {
		follows, err = fu.store.ListFollowsByFromUserID(ctx, ID)
		if err != nil {
			return
		}
	}

	for _, follow := range follows {
		account, err := fu.store.GetAccountsByID(ctx, follow.FromAccountID)
		if err != nil {
			return nil, err
		}
		result = append(result, response.Follow{
			AccountID: account.AccountID,
			Username:  account.Username,
			Icon:      account.Icon.String,
		})
	}
	return
}
