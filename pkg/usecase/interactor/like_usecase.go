package usecase

import (
	"context"
	"time"

	repository "github.com/hackhack-Geek-vol6/backend/pkg/adapter/gateways/repository/datasource"
	"github.com/hackhack-Geek-vol6/backend/pkg/adapter/gateways/repository/transaction"
	"github.com/hackhack-Geek-vol6/backend/pkg/domain"
	"github.com/hackhack-Geek-vol6/backend/pkg/usecase/inputport"
)

type likeUsecase struct {
	store          transaction.Store
	contextTimeout time.Duration
}

func NewLikeUsercase(store transaction.Store, timeout time.Duration) inputport.LikeUsecase {
	return &likeUsecase{
		store:          store,
		contextTimeout: timeout,
	}
}

func (bu *likeUsecase) CreateLike(ctx context.Context, body repository.CreateLikesParams) (repository.Like, error) {
	ctx, cancel := context.WithTimeout(ctx, bu.contextTimeout)
	defer cancel()

	return bu.store.CreateLikes(ctx, body)
}

func (bu *likeUsecase) GetLike(ctx context.Context, id string, query domain.ListRequest) ([]repository.Like, error) {
	ctx, cancel := context.WithTimeout(ctx, bu.contextTimeout)
	defer cancel()

	return bu.store.ListLikesByID(ctx, id)
}

func (bu *likeUsecase) RemoveLike(ctx context.Context, accountID string, opus int32) error {
	ctx, cancel := context.WithTimeout(ctx, bu.contextTimeout)
	defer cancel()

	_, err := bu.store.DeleteLikesByID(ctx, repository.DeleteLikesByIDParams{AccountID: accountID, Opus: opus})
	return err
}
