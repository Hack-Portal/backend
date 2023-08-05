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

func NewAccountUsercase(store db.Store, timeout time.Duration) domain.AccountUsecase {
	return &accountUsecase{
		store:          store,
		contextTimeout: timeout,
	}
}

func (au *accountUsecase) GetAccountByID(ctx context.Context, id string) (domain.AccountResponses, error) {
	ctx, cancel := context.WithTimeout(ctx, au.contextTimeout)
	defer cancel()

	return au.store.GetAccountTxByID(ctx, id)
}

func (au *accountUsecase) GetAccountByEmail(ctx context.Context, email string) (domain.AccountResponses, error) {
	ctx, cancel := context.WithTimeout(ctx, au.contextTimeout)
	defer cancel()

	return au.store.GetAccountTxByEmail(ctx, email)
}

func (au *accountUsecase) CreateAccount(ctx context.Context, body db.CreateAccountTxParams) (domain.AccountResponses, error) {
	ctx, cancel := context.WithTimeout(ctx, au.contextTimeout)
	defer cancel()

	return au.store.CreateAccountTx(ctx, body)
}

func (au *accountUsecase) UpdateAccount(ctx context.Context, body db.UpdateAccountTxParams) (domain.AccountResponses, error) {
	ctx, cancel := context.WithTimeout(ctx, au.contextTimeout)
	defer cancel()

	return au.store.UpdateAccountTx(ctx, body)
}

func (au *accountUsecase) DeleteAccount(ctx context.Context, id string) error {
	ctx, cancel := context.WithTimeout(ctx, au.contextTimeout)
	defer cancel()
	_, err := au.store.SoftDeleteAccount(ctx, id)
	return err
}

func (au *accountUsecase) UploadImage(ctx context.Context, body []byte) (string, error) {
	ctx, cancel := context.WithTimeout(ctx, au.contextTimeout)
	defer cancel()

	return au.store.UploadImage(ctx, body)
}
