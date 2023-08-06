package usecase

import (
	"context"
	"time"

	db "github.com/hackhack-Geek-vol6/backend/db/sqlc"
	"github.com/hackhack-Geek-vol6/backend/domain"
)

type rateUsecase struct {
	store          db.Store
	contextTimeout time.Duration
}

func NewRateUsercase(store db.Store, timeout time.Duration) domain.RateUsecase {
	return &rateUsecase{
		store:          store,
		contextTimeout: timeout,
	}
}

func (ru *rateUsecase) CreateRateEntry(ctx context.Context, body db.CreateRateParams) (db.RateEntries, error) {
	ctx, cancel := context.WithTimeout(ctx, ru.contextTimeout)
	defer cancel()

	rate, err := ru.store.CreateRate(ctx, body)
	if err != nil {
		return db.RateEntries{}, err
	}

	_, err = ru.store.UpdateRateByUserID(ctx, db.UpdateRateByUserIDParams{UserID: body.UserID, Rate: body.Rate})
	if err != nil {
		return db.RateEntries{}, err
	}

	return rate, nil
}

func (ru *rateUsecase) ListRateEntry(ctx context.Context, id string, query domain.ListRateParams) ([]db.RateEntries, error) {
	ctx, cancel := context.WithTimeout(ctx, ru.contextTimeout)
	defer cancel()

	rates, err := ru.store.ListRate(ctx, db.ListRateParams{
		UserID: id,
		Limit:  query.PageSize,
		Offset: (query.PageId - 1) * query.PageSize,
	})
	if err != nil {
		return nil, err
	}
	return rates, nil
}
