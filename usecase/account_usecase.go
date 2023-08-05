package usecase

import (
	"context"
	"time"

	db "github.com/hackhack-Geek-vol6/backend/db/sqlc"
	"github.com/hackhack-Geek-vol6/backend/domain"
)

type accountUsecase struct {
	store          db.Store
	contextTimeout time.Duration
}

func NewAccountUsercase(accountRepository domain.AccountRepository, timeout time.Duration) domain.AccountUsecase {
	return &accountUsecase{
		accountRepository: accountRepository,
		contextTimeout:    timeout,
	}
}

func (au *accountUsecase) GetAccountByID(ctx context.Context, ID string) (domain.AccountResponses, error) {
	ctx, cancel := context.WithTimeout(ctx, au.contextTimeout)
	defer cancel()
	return au.store.GetAccountByID(ctx, ID)
}

func (au *accountUsecase) CreateAccount(ctx context.Context, body db.CreateAccountTxParams) (domain.AccountResponses, error) {
	ctx, cancel := context.WithTimeout(ctx, au.contextTimeout)
	defer cancel()
	return au.store.CreateAccountTx(ctx, body)
}
