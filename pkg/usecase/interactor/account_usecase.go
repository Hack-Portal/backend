package usecase

import (
	"context"
	"database/sql"
	"time"

	"github.com/google/uuid"
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

	account, err := au.store.GetAccountsByID(ctx, id)
	if err != nil {
		return domain.AccountResponses{}, err
	}

	locate, err := au.store.GetLocatesByID(ctx, account.LocateID)
	if err != nil {
		return domain.AccountResponses{}, err
	}

	tags, err := au.store.ListAccountTagsByUserID(ctx, id)
	if err != nil {
		return domain.AccountResponses{}, err
	}

	fws, err := au.store.ListAccountFrameworksByUserID(ctx, id)
	if err != nil {
		return domain.AccountResponses{}, err
	}

	techTags, err := parseTechTags(ctx, au.store, accountTechTagsStruct(tags))
	if err != nil {
		return domain.AccountResponses{}, err
	}

	frameworks, err := parseFrameworks(ctx, au.store, accountFWStruct(fws))
	if err != nil {
		return domain.AccountResponses{}, err
	}

	return parseAccountResponse(repository.Account{
		UserID:          account.UserID,
		Username:        account.Username,
		Icon:            account.Icon,
		ExplanatoryText: account.ExplanatoryText,
		Rate:            account.Rate,
		ShowRate:        account.ShowRate,
		ShowLocate:      account.ShowLocate,
	},
		locate.Name,
		techTags,
		frameworks,
	), nil
}

func (au *accountUsecase) ListAccountRate(ctx context.Context, id string) (domain.AccountRateResponse, error){
	ctx, cancel := context.WithTimeout(ctx, au.contextTimeout)
	defer cancel()

	accounts ,err := au.store.ListAccounts(ctx,repository.ListAccountsParams{
		Username: ,
	})
}

func (au *accountUsecase) GetAccountByEmail(ctx context.Context, email string) (domain.AccountResponses, error) {
	ctx, cancel := context.WithTimeout(ctx, au.contextTimeout)
	defer cancel()

	account, err := au.store.GetAccountsByEmail(ctx, sql.NullString{String: email, Valid: true})
	if err != nil {
		return domain.AccountResponses{}, err
	}

	locate, err := au.store.GetLocatesByID(ctx, account.LocateID)
	if err != nil {
		return domain.AccountResponses{}, err
	}

	tags, err := au.store.ListAccountTagsByUserID(ctx, account.UserID)
	if err != nil {
		return domain.AccountResponses{}, err
	}

	fws, err := au.store.ListAccountFrameworksByUserID(ctx, account.UserID)
	if err != nil {
		return domain.AccountResponses{}, err
	}

	techTags, err := parseTechTags(ctx, au.store, accountTechTagsStruct(tags))
	if err != nil {
		return domain.AccountResponses{}, err
	}

	frameworks, err := parseFrameworks(ctx, au.store, accountFWStruct(fws))
	if err != nil {
		return domain.AccountResponses{}, err
	}

	return parseAccountResponse(repository.Account{
		UserID:          account.UserID,
		Username:        account.Username,
		Icon:            account.Icon,
		ExplanatoryText: account.ExplanatoryText,
		Rate:            account.Rate,
		ShowRate:        account.ShowRate,
		ShowLocate:      account.ShowLocate,
	},
		locate.Name,
		techTags,
		frameworks,
	), nil
}

func (au *accountUsecase) CreateAccount(ctx context.Context, body domain.CreateAccountRequest, image []byte, email string) (domain.AccountResponses, error) {
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

	account, err := au.store.CreateAccountTx(ctx, domain.CreateAccountParams{
		AccountInfo: repository.CreateAccountsParams{
			AccountID: uuid.New().String(),
			UserID:    body.UserID,
			Username:  body.Username,
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
	if err != nil {
		return domain.AccountResponses{}, err
	}

	locate, err := au.store.GetLocatesByID(ctx, account.LocateID)
	if err != nil {
		return domain.AccountResponses{}, err
	}

	techTags, err := parseTechTags(ctx, au.store, body.TechTags)
	if err != nil {
		return domain.AccountResponses{}, err
	}

	frameworks, err := parseFrameworks(ctx, au.store, body.Frameworks)
	if err != nil {
		return domain.AccountResponses{}, err
	}
	return parseAccountResponse(account, locate.Name, techTags, frameworks), nil
}

func (au *accountUsecase) UpdateAccount(ctx context.Context, body domain.UpdateAccountParam, image []byte) (domain.AccountResponses, error) {
	ctx, cancel := context.WithTimeout(ctx, au.contextTimeout)
	defer cancel()

	imageURL, err := uploadImage(ctx, au.store, image)
	if err != nil {
		return domain.AccountResponses{}, err
	}

	body.AccountInfo.Icon = sql.NullString{
		String: imageURL,
		Valid:  true,
	}

	account, err := au.store.UpdateAccountTx(ctx, body)
	if err != nil {
		return domain.AccountResponses{}, err
	}

	locate, err := au.store.GetLocatesByID(ctx, account.LocateID)
	if err != nil {
		return domain.AccountResponses{}, err
	}

	techTags, err := parseTechTags(ctx, au.store, body.AccountTechTag)
	if err != nil {
		return domain.AccountResponses{}, err
	}

	frameworks, err := parseFrameworks(ctx, au.store, body.AccountFrameworkTag)
	if err != nil {
		return domain.AccountResponses{}, err
	}
	return parseAccountResponse(account, locate.Name, techTags, frameworks), nil
}

func (au *accountUsecase) DeleteAccount(ctx context.Context, id string) error {
	ctx, cancel := context.WithTimeout(ctx, au.contextTimeout)
	defer cancel()
	_, err := au.store.DeleteAccounts(ctx, id)
	return err
}

func parseAccountResponse(account repository.Account, locate string, techTags []repository.TechTag, frameworks []repository.Framework) domain.AccountResponses {
	return domain.AccountResponses{
		AccountID:       account.AccountID,
		Username:        account.Username,
		Icon:            account.Icon.String,
		ExplanatoryText: account.ExplanatoryText.String,
		Rate:            account.Rate,
		Locate:          locate,
		ShowRate:        account.ShowRate,
		ShowLocate:      account.ShowLocate,
		TechTags:        techTags,
		Frameworks:      frameworks,
	}
}

func uploadImage(ctx context.Context, store transaction.Store, image []byte) (string, error) {
	var imageURL string
	if image != nil {
		var err error
		imageURL, err = store.UploadImage(ctx, image)
		if err != nil {
			return "", err
		}
	}
	return imageURL, nil
}
