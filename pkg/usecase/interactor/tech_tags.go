package usecase

import (
	"context"

	repository "github.com/hackhack-Geek-vol6/backend/pkg/adapter/gateways/repository/datasource"
	"github.com/hackhack-Geek-vol6/backend/pkg/adapter/gateways/repository/transaction"
)

func parseTechTags(ctx context.Context, store transaction.Store, tags []int32) (result []repository.TechTag, err error) {
	for _, tag := range tags {
		techTag, err := store.GetTechTagsByID(ctx, tag)
		if err != nil {
			return result, err
		}
		result = append(result, repository.TechTag{TechTagID: techTag.TechTagID, Language: techTag.Language})
	}
	return
}

func accountTechTagsStruct(tags []repository.ListAccountTagsByUserIDRow) []int32 {
	var result []int32
	for _, tag := range tags {
		result = append(result, tag.TechTagID.Int32)
	}
	return result
}
