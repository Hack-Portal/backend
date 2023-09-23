package usecase

import (
	"context"
	"time"

	repository "github.com/hackhack-Geek-vol6/backend/pkg/adapter/gateways/repository/datasource"
	"github.com/hackhack-Geek-vol6/backend/pkg/adapter/gateways/repository/transaction"
	"github.com/hackhack-Geek-vol6/backend/pkg/domain/params"
	"github.com/hackhack-Geek-vol6/backend/pkg/domain/response"
	"github.com/hackhack-Geek-vol6/backend/pkg/usecase/inputport"
	dbutil "github.com/hackhack-Geek-vol6/backend/pkg/util/db"
	"github.com/hackhack-Geek-vol6/backend/pkg/util/jwt"
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

func (au *accountUsecase) GetAccountByID(ctx context.Context, id string, token *jwt.FireBaseCustomToken) (result response.AccountResponse, err error) {
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

	techTags, err := parseTechTags(ctx, au.store, account.AccountID)
	if err != nil {
		return
	}

	frameworks, err := parseFrameworks(ctx, au.store, account.AccountID)
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

	if token == nil {
		if result.ShowLocate {
			result.Locate = locate.Name
		}
		if result.ShowRate {
			result.Rate = account.Rate
		}
		return
	}

	cnt, err := au.store.CheckAccount(ctx, repository.CheckAccountParams{
		AccountID: account.AccountID,
		Email:     token.Email,
	})
	if err != nil {
		return
	}

	if cnt >= 1 {
		result.Locate = locate.Name
		result.Rate = account.Rate
	} else {
		if result.ShowLocate {
			result.Locate = locate.Name
		}
		if result.ShowRate {
			result.Rate = account.Rate
		}
	}

	return
}

func (au *accountUsecase) GetAccountByEmail(ctx context.Context, email string) (result response.AccountResponse, err error) {
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

	techTags, err := parseTechTags(ctx, au.store, account.AccountID)
	if err != nil {
		return
	}

	frameworks, err := parseFrameworks(ctx, au.store, account.AccountID)
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

func (au *accountUsecase) CreateAccount(ctx context.Context, body params.CreateAccount, image []byte) (result response.AccountResponse, err error) {
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
	body.AccountInfo.Icon = dbutil.ToSqlNullString(imageURL)

	account, err := au.store.CreateAccountTx(ctx, body)
	if err != nil {
		return
	}

	locate, err := au.store.GetLocatesByID(ctx, account.LocateID)
	if err != nil {
		return
	}

	techTags, err := parseTechTags(ctx, au.store, account.AccountID)
	if err != nil {
		return
	}

	frameworks, err := parseFrameworks(ctx, au.store, account.AccountID)
	if err != nil {
		return
	}
	result = parseAccountResponse(account, locate.Name, techTags, frameworks)
	return
}

func (au *accountUsecase) UpdateAccount(ctx context.Context, body params.UpdateAccount, image []byte) (result response.AccountResponse, err error) {
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

	techTags, err := parseTechTags(ctx, au.store, account.AccountID)
	if err != nil {
		return
	}

	frameworks, err := parseFrameworks(ctx, au.store, account.AccountID)
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

func (au *accountUsecase) GetJoinRoom(ctx context.Context, accountID string) (result []response.GetJoinRoomResponse, err error) {
	ctx, cancel := context.WithTimeout(ctx, au.contextTimeout)
	defer cancel()

	rooms, err := au.store.ListJoinRoomByID(ctx, repository.ListJoinRoomByIDParams{
		AccountID: accountID,
		Expired:   time.Now().Add(-time.Hour * 24 * 30),
	})
	if err != nil {
		return
	}

	for _, room := range rooms {
		result = append(result, response.GetJoinRoomResponse{
			RoomID: room.RoomID,
			Title:  room.Title.String,
		})
	}
	return
}

func parseAccountResponse(account repository.Account, locate string, techTags []repository.TechTag, frameworks []repository.Framework) response.AccountResponse {
	return response.AccountResponse{
		AccountID:       account.AccountID,
		Username:        account.Username,
		Icon:            account.Icon.String,
		ExplanatoryText: account.ExplanatoryText.String,
		GithubLink:      account.GithubLink.String,
		TwitterLink:     account.TwitterLink.String,
		DiscordLink:     account.DiscordLink.String,
		ShowRate:        account.ShowRate,
		ShowLocate:      account.ShowLocate,
		TechTags:        techTags,
		Frameworks:      frameworks,
	}
}
