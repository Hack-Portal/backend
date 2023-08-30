package usecase

import (
	"context"

	repository "github.com/hackhack-Geek-vol6/backend/pkg/adapter/gateways/repository/datasource"
	"github.com/hackhack-Geek-vol6/backend/pkg/adapter/gateways/repository/transaction"
)

func parseTechTags(ctx context.Context, store transaction.Store, accountID string) (result []repository.TechTag, err error) {
	tags, err := store.ListAccountTagsByUserID(ctx, accountID)
	if err != nil {
		return
	}
	for _, tag := range tags {
		result = append(result, repository.TechTag{
			TechTagID: tag.TechTagID.Int32,
			Language:  tag.Language.String,
			Icon:      tag.Icon.String,
		})
	}
	return
}
