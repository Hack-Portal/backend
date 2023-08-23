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

func (ru *rateUsecase) CreateRateEntry(ctx context.Context, body repository.CreateRateEntitiesParams) error {
	ctx, cancel := context.WithTimeout(ctx, ru.contextTimeout)
	defer cancel()

	if err := ru.store.CreateRateEntityTx(ctx, body); err != nil {
		return err
	}

	return nil
}

func (ru *rateUsecase) ListRateEntry(ctx context.Context, id string, query domain.ListRequest) ([]repository.RateEntity, error) {
	ctx, cancel := context.WithTimeout(ctx, ru.contextTimeout)
	defer cancel()

	rates, err := ru.store.ListRateEntities(ctx, repository.ListRateEntitiesParams{
		AccountID: id,
		Limit:     query.PageSize,
		Offset:    (query.PageID - 1) * query.PageSize,
	})
	if err != nil {
		return nil, err
	}
	return rates, nil
}

func (au *rateUsecase) ListAccountRate(ctx context.Context, args domain.ListRequest) (result []domain.AccountRateResponse, err error) {
	ctx, cancel := context.WithTimeout(ctx, au.contextTimeout)
	defer cancel()

	accounts, err := au.store.ListAccounts(ctx, repository.ListAccountsParams{
		Limit:  args.PageSize,
		Offset: (args.PageID - 1) * args.PageSize,
	})
	if err != nil {
		return
	}

	return parseAccountRateResponse(accounts), nil
}

func parseAccountRateResponse(accounts []repository.Account) (result []domain.AccountRateResponse) {
	for _, account := range accounts {
		result = append(result, domain.AccountRateResponse{
			AccountID: account.AccountID,
			Username:  account.Username,
			Icon:      account.Icon.String,
			Rate:      account.Rate,
		})
	}
	return
}
