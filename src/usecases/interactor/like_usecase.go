package usecase

import (
	"context"
	"time"

	repository "github.com/hackhack-Geek-vol6/backend/pkg/adapter/gateways/repository/datasource"
	"github.com/hackhack-Geek-vol6/backend/pkg/adapter/gateways/repository/transaction"
	"github.com/hackhack-Geek-vol6/backend/pkg/domain/request"
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

func (lu *likeUsecase) CreateLike(ctx context.Context, body repository.CreateLikesParams) (repository.Like, error) {
	ctx, cancel := context.WithTimeout(ctx, lu.contextTimeout)
	defer cancel()

	return lu.store.CreateLikes(ctx, body)
}

func (lu *likeUsecase) GetLike(ctx context.Context, id string, query request.ListRequest) (result []repository.Like, err error) {
	ctx, cancel := context.WithTimeout(ctx, lu.contextTimeout)
	defer cancel()

	return lu.store.ListLikesByID(ctx, id)
}

func (lu *likeUsecase) RemoveLike(ctx context.Context, accountID string, opus int32) error {
	ctx, cancel := context.WithTimeout(ctx, lu.contextTimeout)
	defer cancel()

	return lu.store.DeleteLikesByID(ctx, repository.DeleteLikesByIDParams{AccountID: accountID, Opus: opus})
}
