package db

import (
	"context"
)

type CreatePastWorkTxParams struct {
	Name               string `json:"name"`
	ThumbnailImage     []byte `json:"thumbnail_image"`
	ExplanatoryText    string `json:"explanatory_text"`
	PastWorkTags       []int32
	PastWorkFrameworks []int32
	AccountPastWorks   []string
}
type CreatePastWorkTxResult struct {
	PastWorks
	PastWorkTags       []PastWorkTags
	PastWorkFrameworks []PastWorkFrameworks
	AccountPastWorks   []AccountPastWorks
}

// 過去作品を登録時のトランザクション
func (store *SQLStore) CreatePastWorkTx(ctx context.Context, arg CreatePastWorkTxParams) (CreatePastWorkTxResult, error) {
	var result CreatePastWorkTxResult
	err := store.execTx(ctx, func(q *Queries) error {
		var err error
		// 過去作品を登録する
		result.PastWorks, err = q.CreatePastWorks(ctx, CreatePastWorksParams{
			Name:            arg.Name,
			ThumbnailImage:  arg.ThumbnailImage,
			ExplanatoryText: arg.ExplanatoryText,
		})
		if err != nil {
			return err
		}
		// 過去作品IDからタグのレコードを登録する
		for _, tag_id := range arg.PastWorkTags {
			_, err = q.CreatePastWorkTag(ctx, CreatePastWorkTagParams{
				Opus:      result.Opus,
				TechTagID: tag_id,
			})
			if err != nil {
				return err
			}
		}
		pastTag, err := q.GetPastWorkTagsByOpus(ctx, result.Opus)
		if err != nil {
			return err
		}
		result.PastWorkTags = pastTag

		// 過去作品IDからフレームワークのレコードを登録する
		for _, framework_id := range arg.PastWorkFrameworks {
			_, err = q.CreatePastWorkFrameworks(ctx, CreatePastWorkFrameworksParams{
				Opus:        result.Opus,
				FrameworkID: framework_id,
			})
			if err != nil {
				return err
			}
		}
		framework, err := q.GetPastWorkFrameworksByOpus(ctx, result.Opus)
		if err != nil {
			return err
		}
		result.PastWorkFrameworks = framework

		// 過去作品IDからアカウントのレコードを登録する
		for _, account := range arg.AccountPastWorks {
			_, err = q.CreateAccountPastWorks(ctx, CreateAccountPastWorksParams{
				Opus:   result.Opus,
				UserID: account,
			})
			if err != nil {
				return err
			}
		}
		account, err := q.GetAccountPastWorksByOpus(ctx, result.Opus)
		if err != nil {
			return err
		}
		result.AccountPastWorks = account
		return nil
	})
	return result, err
}
