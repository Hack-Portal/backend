package usecase

import (
	"context"

	repository "github.com/hackhack-Geek-vol6/backend/pkg/adapter/gateways/repository/datasource"
	"github.com/hackhack-Geek-vol6/backend/pkg/adapter/gateways/repository/transaction"
)

func parseFrameworks(ctx context.Context, store transaction.Store, accountID string) (result []repository.Framework, err error) {
	frameworks, err := store.ListAccountFrameworksByUserID(ctx, accountID)
	if err != nil {
		return
	}
	for _, framework := range frameworks {
		result = append(result, repository.Framework{
			FrameworkID: framework.FrameworkID.Int32,
			TechTagID:   framework.TechTagID.Int32,
			Framework:   framework.Framework.String,
		})
	}
	return
}
