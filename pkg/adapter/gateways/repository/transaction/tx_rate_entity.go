package transaction

import (
	"context"

	repository "github.com/hackhack-Geek-vol6/backend/pkg/adapter/gateways/repository/datasource"
)

func (store *SQLStore) CreateRateEntityTx(ctx context.Context, arg repository.CreateRateEntitiesParams) error {
	err := store.execTx(ctx, func(q *repository.Queries) error {
		var err error

		if _, err = q.CreateRateEntities(ctx, arg); err != nil {
			return err
		}

		if _, err = q.UpdateRateByID(ctx, repository.UpdateRateByIDParams{AccountID: arg.AccountID, Rate: arg.Rate}); err != nil {
			return err
		}

		return nil
	})
	return err
}
