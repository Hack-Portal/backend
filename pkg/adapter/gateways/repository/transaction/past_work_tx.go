package transaction

import (
	"context"

	repository "github.com/hackhack-Geek-vol6/backend/pkg/adapter/gateways/repository/datasource"
	"github.com/hackhack-Geek-vol6/backend/pkg/domain"
)

func createPastWorkTags(ctx context.Context, q *repository.Queries, opus int32, tags []int32) error {
	for _, tag := range tags {
		_, err := q.CreatePastWorkTags(ctx, repository.CreatePastWorkTagsParams{Opus: opus, TechTagID: tag})
		if err != nil {
			return err
		}
	}
	return nil
}

func createPastWorkFrameworks(ctx context.Context, q *repository.Queries, opus int32, frameworks []int32) error {
	for _, framework := range frameworks {
		_, err := q.CreatePastWorkFrameworks(ctx, repository.CreatePastWorkFrameworksParams{Opus: opus, FrameworkID: framework})
		if err != nil {
			return err
		}
	}
	return nil
}

func createPastWorkMembers(ctx context.Context, q *repository.Queries, opus int32, id []string) error {
	for _, userid := range id {
		_, err := q.CreateAccountPastWorks(ctx, repository.CreateAccountPastWorksParams{Opus: opus, UserID: userid})
		if err != nil {
			return err
		}
	}
	return nil
}

func (store *SQLStore) CreatePastWorkTx(ctx context.Context, arg domain.CreatePastWorkParams) (repository.PastWork, error) {
	var pastwork repository.PastWork
	err := store.execTx(ctx, func(q *repository.Queries) error {
		var err error
		pastwork, err = q.CreatePastWorks(ctx, repository.CreatePastWorksParams{
			Name:            arg.Name,
			ThumbnailImage:  arg.ThumbnailImage,
			ExplanatoryText: arg.ExplanatoryText,
		})
		if err != nil {
			return err
		}

		if err := createPastWorkTags(ctx, q, pastwork.Opus, arg.PastWorkTags); err != nil {
			return err
		}

		if err := createPastWorkFrameworks(ctx, q, pastwork.Opus, arg.PastWorkTags); err != nil {
			return err
		}

		if err := createPastWorkMembers(ctx, q, pastwork.Opus, arg.AccountPastWorks); err != nil {
			return err
		}
		return nil
	})
	return pastwork, err
}
