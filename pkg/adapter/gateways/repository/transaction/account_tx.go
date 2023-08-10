package transaction

import (
	"context"
	"database/sql"

	repository "github.com/hackhack-Geek-vol6/backend/pkg/adapter/gateways/repository/datasource"
	"github.com/hackhack-Geek-vol6/backend/pkg/domain"
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
		Icon:      latest.Icon,
		Rate:      latest.Rate,
	}

	if len(request.Username) != 0 {
		if latest.Username != request.Username {
			result.Username = request.Username
		}
	} else {
		result.Username = latest.Username
	}

	if len(request.ExplanatoryText.String) != 0 {
		if latest.ExplanatoryText.String != request.ExplanatoryText.String {
			result.ExplanatoryText = sql.NullString{
				String: request.ExplanatoryText.String,
				Valid:  true,
			}
		}
	} else {
		result.ExplanatoryText = latest.ExplanatoryText
	}

	if len(request.Icon.String) != 0 {
		if latest.Icon.String != request.Icon.String {
			result.Icon = sql.NullString{
				String: request.Icon.String,
				Valid:  true,
			}
		}
	} else {
		result.Icon = latest.Icon
	}

	if request.LocateID != 0 {
		if latest.LocateID != request.LocateID {
			result.LocateID = request.LocateID
		}
	} else {
		result.LocateID = latest.LocateID
	}

	if latest.ShowLocate != request.ShowLocate {
		result.ShowLocate = request.ShowLocate
	} else {
		result.ShowLocate = latest.ShowLocate
	}

	if latest.ShowRate != request.ShowRate {
		result.ShowRate = request.ShowRate
	} else {
		result.ShowRate = latest.ShowRate
	}

	return
}

func (store *SQLStore) CreateAccountTx(ctx context.Context, args domain.CreateAccountParams) (repository.Account, error) {
	var account repository.Account
	err := store.execTx(ctx, func(q *repository.Queries) error {
		var err error
		account, err = q.CreateAccounts(ctx, args.AccountInfo)
		if err != nil {
			return err
		}

		if err := createAccountFrameworks(ctx, q, args.AccountInfo.UserID, args.AccountTechTag); err != nil {
			return err
		}

		if err := createAccountFrameworks(ctx, q, args.AccountInfo.UserID, args.AccountFrameworkTag); err != nil {
			return err
		}

		return nil
	})
	return account, err
}

func (store *SQLStore) UpdateAccountTx(ctx context.Context, args domain.UpdateAccountParam) (repository.Account, error) {
	var account repository.Account
	err := store.execTx(ctx, func(q *repository.Queries) error {
		latest, err := q.GetAccountsByID(ctx, args.AccountInfo.UserID)
		if err != nil {
			return err
		}

		account, err = q.UpdateAccounts(ctx, compAccount(args.AccountInfo, latest))
		if err != nil {
			return err
		}

		// 以下タグ部分
		err = q.DeleteAccountTagsByUserID(ctx, latest.UserID)
		if err != nil {
			return err
		}

		err = q.DeleteAccountFrameworkByUserID(ctx, latest.UserID)
		if err != nil {
			return err
		}

		if err := createAccountFrameworks(ctx, q, args.AccountInfo.UserID, args.AccountTechTag); err != nil {
			return err
		}

		if err := createAccountFrameworks(ctx, q, args.AccountInfo.UserID, args.AccountFrameworkTag); err != nil {
			return err
		}

		return nil
	})
	return account, err
}
