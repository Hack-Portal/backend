package usecase

import (
	"context"
	"database/sql"
	"time"

	repository "github.com/hackhack-Geek-vol6/backend/pkg/adapter/gateways/repository/datasource"
	"github.com/hackhack-Geek-vol6/backend/pkg/adapter/gateways/repository/transaction"
	"github.com/hackhack-Geek-vol6/backend/pkg/domain"
	"github.com/hackhack-Geek-vol6/backend/pkg/usecase/inputport"
)

type accountUsecase struct {
	store          transaction.Store
	contextTimeout time.Duration
}

func NewAccountUsercase(store transaction.Store, timeout time.Duration) inputport.AccountUsecase {
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

func (au *accountUsecase) CreateAccount(ctx context.Context, body domain.CreateAccountRequest, image []byte) (domain.AccountResponses, error) {
	ctx, cancel := context.WithTimeout(ctx, au.contextTimeout)
	defer cancel()
	// 画像が空やないときに処理する
	var imageURL string
	if image != nil {
		var err error
		imageURL, err = au.store.UploadImage(ctx, image)
		if err != nil {
			return domain.AccountResponses{}, err
		}
	}

	return au.store.CreateAccountTx(ctx, domain.CreateAccountParams{
		AccountInfo: repository.CreateAccountsParams{
			UserID:   body.UserID,
			Username: body.Username,
			Icon: sql.NullString{
				String: imageURL,
				Valid:  true,
			},
			ExplanatoryText: sql.NullString{
				String: body.ExplanatoryText,
				Valid:  true,
			},
			LocateID:   body.LocateID,
			Rate:       0,
			ShowLocate: body.ShowLocate,
			ShowRate:   body.ShowRate,
		},
		AccountTechTag:      body.TechTags,
		AccountFrameworkTag: body.Frameworks,
	})
}

func (au *accountUsecase) UpdateAccount(ctx context.Context, body domain.UpdateAccountParam, image []byte) (domain.AccountResponses, error) {
	ctx, cancel := context.WithTimeout(ctx, au.contextTimeout)
	defer cancel()

	var imageURL string
	if image != nil {
		var err error
		imageURL, err = au.store.UploadImage(ctx, image)
		if err != nil {
			return domain.AccountResponses{}, err
		}

		body.AccountInfo.Icon = sql.NullString{
			String: imageURL,
			Valid:  true,
		}
	}
	return au.store.UpdateAccountTx(ctx, body)
}

func (au *accountUsecase) DeleteAccount(ctx context.Context, id string) error {
	ctx, cancel := context.WithTimeout(ctx, au.contextTimeout)
	defer cancel()
	_, err := au.store.DeleteAccounts(ctx, id)
	return err
}
