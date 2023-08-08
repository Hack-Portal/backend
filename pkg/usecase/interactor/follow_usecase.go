package usecase

import (
	"context"
	"time"

	repository "github.com/hackhack-Geek-vol6/backend/pkg/adapter/gateways/repository/datasource"
	"github.com/hackhack-Geek-vol6/backend/pkg/adapter/gateways/repository/transaction"
	"github.com/hackhack-Geek-vol6/backend/pkg/domain"
	"github.com/hackhack-Geek-vol6/backend/pkg/usecase/inputport"
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

func (fu *followUsecase) CreateFollow(ctx context.Context, body repository.CreateFollowsParams) (result domain.FollowResponse, err error) {
	ctx, cancel := context.WithTimeout(ctx, fu.contextTimeout)
	defer cancel()

	follow, err := fu.store.CreateFollows(ctx, body)
	if err != nil {
		return
	}

	account, err := fu.store.GetAccountsByID(ctx, follow.ToUserID)
	if err != nil {
		return
	}
	result = domain.FollowResponse{UserID: account.UserID, Username: account.Username, Icon: account.Icon.String}
	return
}

func (fu *followUsecase) RemoveFollow(ctx context.Context, body repository.DeleteFollowsParams) error {
	ctx, cancel := context.WithTimeout(ctx, fu.contextTimeout)
	defer cancel()

	return fu.store.DeleteFollows(ctx, body)
}

func (fu *followUsecase) GetFollowByToID(ctx context.Context, ID string) (result []domain.FollowResponse, err error) {
	ctx, cancel := context.WithTimeout(ctx, fu.contextTimeout)
	defer cancel()

	follows, err := fu.store.ListFollowsByToUserID(ctx, ID)
	if err != nil {
		return
	}
	for _, follow := range follows {
		account, err := fu.store.GetAccountsByID(ctx, follow.FromUserID)
		if err != nil {
			return nil, err
		}
		result = append(result, domain.FollowResponse{
			UserID:   account.UserID,
			Username: account.Username,
			Icon:     account.Icon.String,
		})
	}
	return
}
