package usecase

import (
	"context"
	"time"

	"github.com/hackhack-Geek-vol6/backend/pkg/domain"
	repository "github.com/hackhack-Geek-vol6/backend/pkg/gateways/repository/datasource"
	"github.com/hackhack-Geek-vol6/backend/pkg/gateways/repository/transaction"
	"github.com/hackhack-Geek-vol6/backend/pkg/usecase/inputport"
)

type rateUsecase struct {
	store          transaction.Store
	contextTimeout time.Duration
}

func NewRateUsercase(store transaction.Store, timeout time.Duration) inputport.RateUsecase {
	return &rateUsecase{
		store:          store,
		contextTimeout: timeout,
	}
}

func (ru *rateUsecase) CreateRateEntry(ctx context.Context, body repository.CreateRateParams) (repository.RateEntry, error) {
	ctx, cancel := context.WithTimeout(ctx, ru.contextTimeout)
	defer cancel()

	rate, err := ru.store.CreateRate(ctx, body)
	if err != nil {
		return repository.RateEntry{}, err
	}

	_, err = ru.store.UpdateRateByUserID(ctx, repository.UpdateRateByUserIDParams{UserID: body.UserID, Rate: body.Rate})
	if err != nil {
		return repository.RateEntry{}, err
	}

	return rate, nil
}

func (ru *rateUsecase) ListRateEntry(ctx context.Context, id string, query domain.ListRateParams) ([]repository.RateEntry, error) {
	ctx, cancel := context.WithTimeout(ctx, ru.contextTimeout)
	defer cancel()

	rates, err := ru.store.ListRate(ctx, repository.ListRateParams{
		UserID: id,
		Limit:  query.PageSize,
		Offset: (query.PageId - 1) * query.PageSize,
	})
	if err != nil {
		return nil, err
	}
	return rates, nil
}
