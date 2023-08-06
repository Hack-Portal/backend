package usecase

import (
	"context"
	"time"

	db "github.com/hackhack-Geek-vol6/backend/db/sqlc"
	"github.com/hackhack-Geek-vol6/backend/domain"
)

type followUsecase struct {
	store          db.Store
	contextTimeout time.Duration
}

func NewFollowUsercase(store db.Store, timeout time.Duration) domain.FollowUsecase {
	return &followUsecase{
		store:          store,
		contextTimeout: timeout,
	}
}

func (fu *followUsecase) CreateFollow(ctx context.Context, body db.CreateFollowParams) (result domain.FollowResponse, err error) {
	ctx, cancel := context.WithTimeout(ctx, fu.contextTimeout)
	defer cancel()

	follow, err := fu.store.CreateFollow(ctx, body)
	if err != nil {
		return
	}

	account, err := fu.store.GetAccountByID(ctx, follow.ToUserID)
	if err != nil {
		return
	}
	result = domain.FollowResponse{UserID: account.UserID, Username: account.Username, Icon: account.Icon.String}
	return
}

func (fu *followUsecase) RemoveFollow(ctx context.Context, body db.SoftRemoveBookmarkParams) (err error) {
	ctx, cancel := context.WithTimeout(ctx, fu.contextTimeout)
	defer cancel()

	_, err = fu.store.SoftRemoveBookmark(ctx, body)
	return
}

func (fu *followUsecase) GetFollowByToID(ctx context.Context, ID string) (result []domain.FollowResponse, err error) {
	ctx, cancel := context.WithTimeout(ctx, fu.contextTimeout)
	defer cancel()

	follows, err := fu.store.ListFollowByToUserID(ctx, ID)
	if err != nil {
		return
	}
	for _, follow := range follows {
		account, err := fu.store.GetAccountByID(ctx, follow.FromUserID)
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
