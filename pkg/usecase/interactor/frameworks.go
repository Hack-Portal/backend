package usecase

import (
	"context"

	repository "github.com/hackhack-Geek-vol6/backend/pkg/adapter/gateways/repository/datasource"
	"github.com/hackhack-Geek-vol6/backend/pkg/adapter/gateways/repository/transaction"
)

func parseFrameworks(ctx context.Context, store transaction.Store, frameworks []int32) (result []repository.Framework, err error) {
	for _, framework := range frameworks {
		fw, err := store.GetFrameworksByID(ctx, framework)
		if err != nil {
			return result, err
		}
		result = append(result, repository.Framework{TechTagID: fw.TechTagID, FrameworkID: fw.FrameworkID, Framework: fw.Framework})
	}
	return
}

func accountFWStruct(fws []repository.ListAccountFrameworksByUserIDRow) []int32 {
	var result []int32
	for _, fw := range fws {
		result = append(result, fw.FrameworkID.Int32)
	}
	return result
}
