package transaction

import (
	"context"
	"time"

	repository "github.com/hackhack-Geek-vol6/backend/pkg/adapter/gateways/repository/datasource"
	"github.com/hackhack-Geek-vol6/backend/pkg/domain/params"
	util "github.com/hackhack-Geek-vol6/backend/pkg/util/etc"
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
		_, err := q.CreateAccountPastWorks(ctx, repository.CreateAccountPastWorksParams{Opus: opus, AccountID: userid})
		if err != nil {
			return err
		}
	}
	return nil
}

func compPastWork(request repository.UpdatePastWorksByIDParams, latest repository.PastWork) (result repository.UpdatePastWorksByIDParams) {
	result = repository.UpdatePastWorksByIDParams{
		Opus:            request.Opus,
		Name:            latest.Name,
		ThumbnailImage:  latest.ThumbnailImage,
		ExplanatoryText: latest.ExplanatoryText,
		UpdateAt:        time.Now(),
	}

	if util.CheckDiff(latest.Name, request.Name) {
		result.Name = request.Name
	}

	if util.CheckDiff(latest.ThumbnailImage, request.ThumbnailImage) {
		result.ThumbnailImage = request.ThumbnailImage
	}

	if util.CheckDiff(latest.ExplanatoryText, request.ExplanatoryText) {
		result.ExplanatoryText = request.ExplanatoryText
	}

	return
}

func (store *SQLStore) CreatePastWorkTx(ctx context.Context, arg params.CreatePastWork) (repository.PastWork, error) {
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

		if err := createPastWorkFrameworks(ctx, q, pastwork.Opus, arg.PastWorkFrameworks); err != nil {
			return err
		}

		if err := createPastWorkMembers(ctx, q, pastwork.Opus, arg.AccountPastWorks); err != nil {
			return err
		}
		return nil
	})
	return pastwork, err
}

func (store *SQLStore) UpdatePastWorkTx(ctx context.Context, arg params.UpdatePastWork) (repository.PastWork, error) {
	var pastwork repository.PastWork
	err := store.execTx(ctx, func(q *repository.Queries) error {
		latest, err := q.GetPastWorksByOpus(ctx, arg.Opus)
		if err != nil {
			return err
		}

		pastwork, err = q.UpdatePastWorksByID(ctx, compPastWork(repository.UpdatePastWorksByIDParams{
			Name:            arg.Name,
			ExplanatoryText: arg.ExplanatoryText,
			Opus:            arg.Opus,
		}, latest))
		if err != nil {
			return err
		}

		if err := q.DeletePastWorkTagsByOpus(ctx, arg.Opus); err != nil {
			return err
		}

		if err := q.DeletePastWorkFrameworksByOpus(ctx, arg.Opus); err != nil {
			return err
		}

		if err := q.DeleteAccountPastWorksByOpus(ctx, arg.Opus); err != nil {
			return err
		}

		if err := createPastWorkTags(ctx, q, pastwork.Opus, arg.PastWorkTags); err != nil {
			return err
		}

		if err := createPastWorkFrameworks(ctx, q, pastwork.Opus, arg.PastWorkFrameworks); err != nil {
			return err
		}

		if err := createPastWorkMembers(ctx, q, pastwork.Opus, arg.AccountPastWorks); err != nil {
			return err
		}
		return nil
	})
	return pastwork, err
}
