package transaction

import (
	"context"

	"github.com/hackhack-Geek-vol6/backend/pkg/repository"
)

func (store *SQLStore) CreateAccount(ctx context.Context) {
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
