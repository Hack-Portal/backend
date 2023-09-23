package transaction

import (
	"context"
	"time"

	repository "github.com/hackhack-Geek-vol6/backend/pkg/adapter/gateways/repository/datasource"
	"github.com/hackhack-Geek-vol6/backend/pkg/domain/params"
	dbutil "github.com/hackhack-Geek-vol6/backend/pkg/util/db"
	util "github.com/hackhack-Geek-vol6/backend/pkg/util/etc"
)

func createAccountTags(ctx context.Context, q *repository.Queries, id string, techTags []int32) error {
	for _, techTag := range techTags {
		_, err := q.CreateAccountTags(ctx, repository.CreateAccountTagsParams{
			AccountID: id,
			TechTagID: techTag,
		})
		if err != nil {
			return err
		}
	}
	return nil
}

func createAccountFrameworks(ctx context.Context, q *repository.Queries, id string, frameworks []int32) error {
	for _, framework := range frameworks {
		_, err := q.CreateAccountFrameworks(ctx, repository.CreateAccountFrameworksParams{
			AccountID:   id,
			FrameworkID: framework,
		})
		if err != nil {
			return err
		}
	}
	return nil
}

func compAccount(request repository.Account, latest repository.Account) (result repository.UpdateAccountsParams) {
	result = repository.UpdateAccountsParams{
		AccountID: latest.AccountID,
		Rate:      latest.Rate,

		Username:        latest.Username,
		ExplanatoryText: latest.ExplanatoryText,
		Character:       latest.Character,
		Icon:            latest.Icon,
		LocateID:        latest.LocateID,
		ShowLocate:      latest.ShowLocate,
		ShowRate:        latest.ShowRate,

		TwitterLink: latest.TwitterLink,
		GithubLink:  latest.GithubLink,
		DiscordLink: latest.DiscordLink,

		UpdateAt: time.Now(),
	}

	if util.CheckDiff(latest.Username, request.Username) {
		result.Username = request.Username
	}

	if util.CheckDiff(latest.ExplanatoryText.String, request.ExplanatoryText.String) {
		result.ExplanatoryText = dbutil.ToSqlNullString(request.ExplanatoryText.String)
	}

	if util.CheckDiff(latest.Icon.String, request.Icon.String) {
		result.Icon = dbutil.ToSqlNullString(request.Icon.String)
	}

	if util.CheckDiff(latest.TwitterLink.String, request.TwitterLink.String) {
		result.TwitterLink = dbutil.ToSqlNullString(request.TwitterLink.String)
	}

	if util.CheckDiff(latest.GithubLink.String, request.GithubLink.String) {
		result.GithubLink = dbutil.ToSqlNullString(request.GithubLink.String)
	}

	if util.CheckDiff(latest.DiscordLink.String, request.DiscordLink.String) {
		result.DiscordLink = dbutil.ToSqlNullString(request.DiscordLink.String)
	}

	if request.LocateID != 0 {
		if latest.LocateID != request.LocateID {
			result.LocateID = request.LocateID
		}
	}

	if request.Character.Int32 != 0 {
		if latest.Character != request.Character {
			result.Character = request.Character
		}
	}

	if latest.ShowLocate != request.ShowLocate {
		result.ShowLocate = request.ShowLocate
	}

	if latest.ShowRate != request.ShowRate {
		result.ShowRate = request.ShowRate
	}

	return
}

func (store *SQLStore) CreateAccountTx(ctx context.Context, args params.CreateAccountParams) (repository.Account, error) {
	var account repository.Account
	err := store.execTx(ctx, func(q *repository.Queries) error {
		var err error
		account, err = q.CreateAccounts(ctx, args.AccountInfo)
		if err != nil {
			return err
		}

		if err := createAccountTags(ctx, q, args.AccountInfo.AccountID, args.AccountTechTag); err != nil {
			return err
		}

		if err := createAccountFrameworks(ctx, q, args.AccountInfo.AccountID, args.AccountFrameworkTag); err != nil {
			return err
		}

		return nil
	})
	return account, err
}

func (store *SQLStore) UpdateAccountTx(ctx context.Context, args params.UpdateAccountParams) (repository.Account, error) {
	var account repository.Account
	err := store.execTx(ctx, func(q *repository.Queries) error {
		latest, err := q.GetAccountsByID(ctx, args.AccountInfo.AccountID)
		if err != nil {
			return err
		}

		account, err = q.UpdateAccounts(ctx, compAccount(args.AccountInfo, latest))
		if err != nil {
			return err
		}

		// 以下タグ部分
		err = q.DeleteAccountTagsByUserID(ctx, latest.AccountID)
		if err != nil {
			return err
		}

		err = q.DeleteAccountFrameworkByUserID(ctx, latest.AccountID)
		if err != nil {
			return err
		}

		if err := createAccountTags(ctx, q, args.AccountInfo.AccountID, args.AccountTechTag); err != nil {
			return err
		}

		if err := createAccountFrameworks(ctx, q, args.AccountInfo.AccountID, args.AccountFrameworkTag); err != nil {
			return err
		}

		return nil
	})
	return account, err
}
