package usecase

import (
	"context"
	"database/sql"
	"time"

	repository "github.com/hackhack-Geek-vol6/backend/pkg/adapter/gateways/repository/datasource"
	"github.com/hackhack-Geek-vol6/backend/pkg/adapter/gateways/repository/transaction"
	"github.com/hackhack-Geek-vol6/backend/pkg/domain"
	"github.com/hackhack-Geek-vol6/backend/pkg/usecase/inputport"
	dbutil "github.com/hackhack-Geek-vol6/backend/pkg/util/db"
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

func (au *accountUsecase) GetAccountByID(ctx context.Context, id string, email string) (result domain.AccountResponses, err error) {
	ctx, cancel := context.WithTimeout(ctx, au.contextTimeout)
	defer cancel()

	account, err := au.store.GetAccountsByID(ctx, id)
	if err != nil {
		return
	}

	locate, err := au.store.GetLocatesByID(ctx, account.LocateID)
	if err != nil {
		return
	}

	tags, err := au.store.ListAccountTagsByUserID(ctx, id)
	if err != nil {
		return
	}

	fws, err := au.store.ListAccountFrameworksByUserID(ctx, id)
	if err != nil {
		return
	}

	techTags, err := parseTechTags(ctx, au.store, accountTechTagsStruct(tags))
	if err != nil {
		return
	}

	frameworks, err := parseFrameworks(ctx, au.store, accountFWStruct(fws))
	if err != nil {
		return
	}

	if len(email) == 0 {

	} else {

	}

	result = parseAccountResponse(repository.Account{
		AccountID:       account.AccountID,
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
	)
	return
}

func (au *accountUsecase) GetAccountByEmail(ctx context.Context, email string) (result domain.AccountResponses, err error) {
	ctx, cancel := context.WithTimeout(ctx, au.contextTimeout)
	defer cancel()

	account, err := au.store.GetAccountsByEmail(ctx, email)
	if err != nil {
		return
	}

	locate, err := au.store.GetLocatesByID(ctx, account.LocateID)
	if err != nil {
		return
	}

	tags, err := au.store.ListAccountTagsByUserID(ctx, account.AccountID)
	if err != nil {
		return
	}

	fws, err := au.store.ListAccountFrameworksByUserID(ctx, account.AccountID)
	if err != nil {
		return
	}

	techTags, err := parseTechTags(ctx, au.store, accountTechTagsStruct(tags))
	if err != nil {
		return
	}

	frameworks, err := parseFrameworks(ctx, au.store, accountFWStruct(fws))
	if err != nil {
		return
	}

	result = parseAccountResponse(repository.Account{
		AccountID:       account.AccountID,
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
	)

	return
}

func (au *accountUsecase) CreateAccount(ctx context.Context, body domain.CreateAccount, image []byte, email string) (result domain.AccountResponses, err error) {
	ctx, cancel := context.WithTimeout(ctx, au.contextTimeout)
	defer cancel()
	// 画像が空やないときに処理する

	var imageURL string
	if image != nil {
		_, imageURL, err = au.store.UploadImage(ctx, image)
		if err != nil {
			return
		}
	}

	account, err := au.store.CreateAccountTx(ctx, domain.CreateAccountParams{
		AccountInfo: repository.CreateAccountsParams{
			AccountID: body.ReqBody.AccountID,
			Username:  body.ReqBody.Username,
			Email:     email,
			Icon: sql.NullString{
				String: imageURL,
				Valid:  true,
			},
			ExplanatoryText: sql.NullString{
				String: body.ReqBody.ExplanatoryText,
				Valid:  true,
			},
			LocateID:   body.ReqBody.LocateID,
			Rate:       0,
			ShowLocate: body.ReqBody.ShowLocate,
			ShowRate:   body.ReqBody.ShowRate,
		},
		AccountTechTag:      body.TechTags,
		AccountFrameworkTag: body.Frameworks,
	})
	if err != nil {
		return
	}

	locate, err := au.store.GetLocatesByID(ctx, account.LocateID)
	if err != nil {
		return
	}

	techTags, err := parseTechTags(ctx, au.store, body.TechTags)
	if err != nil {
		return
	}

	frameworks, err := parseFrameworks(ctx, au.store, body.Frameworks)
	if err != nil {
		return
	}
	result = parseAccountResponse(account, locate.Name, techTags, frameworks)
	return
}

func (au *accountUsecase) UpdateAccount(ctx context.Context, body domain.UpdateAccountParam, image []byte) (result domain.AccountResponses, err error) {
	ctx, cancel := context.WithTimeout(ctx, au.contextTimeout)
	defer cancel()

	var imageURL string
	if image != nil {
		_, imageURL, err = au.store.UploadImage(ctx, image)
		if err != nil {
			return
		}
	}

	body.AccountInfo.Icon = dbutil.ToSqlNullString(imageURL)

	account, err := au.store.UpdateAccountTx(ctx, body)
	if err != nil {
		return
	}

	locate, err := au.store.GetLocatesByID(ctx, account.LocateID)
	if err != nil {
		return
	}

	techTags, err := parseTechTags(ctx, au.store, body.AccountTechTag)
	if err != nil {
		return
	}

	frameworks, err := parseFrameworks(ctx, au.store, body.AccountFrameworkTag)
	if err != nil {
		return
	}

	result = parseAccountResponse(account, locate.Name, techTags, frameworks)

	return
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
