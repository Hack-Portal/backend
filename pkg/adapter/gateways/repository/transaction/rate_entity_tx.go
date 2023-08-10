package transaction

import (
	"context"

	repository "github.com/hackhack-Geek-vol6/backend/pkg/adapter/gateways/repository/datasource"
)

func (store *SQLStore) CreateRateEntityTx(ctx context.Context, arg repository.CreateRateEntriesParams) error {
	err := store.execTx(ctx, func(q *repository.Queries) error {
		var err error
		_, err = q.CreateRateEntries(ctx, arg)
		if err != nil {
			return err
		}

		_, err = q.UpdateRateByID(ctx, repository.UpdateRateByIDParams{AccountID: arg.AccountID, Rate: arg.Rate})
		if err != nil {
			return err
		}

		return nil
	})
	return err
}
