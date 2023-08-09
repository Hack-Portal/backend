package usecase

import (
	"context"
	"time"

	repository "github.com/hackhack-Geek-vol6/backend/pkg/adapter/gateways/repository/datasource"
	"github.com/hackhack-Geek-vol6/backend/pkg/adapter/gateways/repository/transaction"
	"github.com/hackhack-Geek-vol6/backend/pkg/domain"
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

func (ru *rateUsecase) CreateRateEntry(ctx context.Context, body repository.CreateRateEntriesParams) error {
	ctx, cancel := context.WithTimeout(ctx, ru.contextTimeout)
	defer cancel()

	if err := ru.store.CreateRateEntitieTx(ctx, body); err != nil {
		return err
	}

	return nil
}

func (ru *rateUsecase) ListRateEntry(ctx context.Context, id string, query domain.ListRateParams) ([]repository.RateEntry, error) {
	ctx, cancel := context.WithTimeout(ctx, ru.contextTimeout)
	defer cancel()

	rates, err := ru.store.ListRateEntries(ctx, repository.ListRateEntriesParams{
		UserID: id,
		Limit:  query.PageSize,
		Offset: (query.PageId - 1) * query.PageSize,
	})
	if err != nil {
		return nil, err
	}
	return rates, nil
}
