package usecase

import (
	"context"
	"time"

	"github.com/hackhack-Geek-vol6/backend/cmd/config"
	"github.com/hackhack-Geek-vol6/backend/pkg/adapter/gateways/repository/transaction"
	"github.com/hackhack-Geek-vol6/backend/pkg/logger"
	"github.com/hackhack-Geek-vol6/backend/src/domain/request"
	"github.com/hackhack-Geek-vol6/backend/src/repository"
	"github.com/hackhack-Geek-vol6/backend/src/usecases/inputport"
)

type likeUsecase struct {
	store   transaction.Store
	l       logger.Logger
	timeout time.Duration
}

func NewLikeUsercase(store transaction.Store, l logger.Logger) inputport.LikeUsecase {
	return &likeUsecase{
		store:   store,
		l:       l,
		timeout: time.Duration(config.Config.Server.ContextTimeout),
	}
}

func (lu *likeUsecase) CreateLike(ctx context.Context, body repository.CreateLikesParams) (repository.Like, error) {
	ctx, cancel := context.WithTimeout(ctx, lu.timeout)
	defer cancel()

	return lu.store.CreateLikes(ctx, body)
}

func (lu *likeUsecase) GetLike(ctx context.Context, id string, query request.ListRequest) (result []repository.Like, err error) {
	ctx, cancel := context.WithTimeout(ctx, lu.timeout)
	defer cancel()

	return lu.store.ListLikesByID(ctx, id)
}

func (lu *likeUsecase) RemoveLike(ctx context.Context, accountID string, opus int32) error {
	ctx, cancel := context.WithTimeout(ctx, lu.timeout)
	defer cancel()

	return lu.store.DeleteLikesByID(ctx, repository.DeleteLikesByIDParams{AccountID: accountID, Opus: opus})
}
