package db

import (
	"context"

	"github.com/hackhack-Geek-vol6/backend/domain"
)

type CreateAccountTxParams struct {
	// ユーザ登録部分
	Accounts
	// tech_tag登録用
	AccountTechTag []int32
	// FrameworkTag登録用
	AccountFrameworkTag []int32
}

// アカウント登録時のトランザクション
func (store *SQLStore) CreateAccountTx(ctx context.Context, arg CreateAccountTxParams) (domain.AccountResponses, error) {
	var result domain.AccountResponses
	err := store.execTx(ctx, func(q *Queries) error {
		var err error
		// アカウントを登録する
		account, err := q.CreateAccount(ctx, CreateAccountParams{
			UserID:          arg.UserID,
			Username:        arg.Username,
			Icon:            arg.Icon,
			ExplanatoryText: arg.ExplanatoryText,
			LocateID:        arg.LocateID,
			Rate:            arg.Rate,
			HashedPassword:  arg.HashedPassword,
			Email:           arg.Email,
			ShowLocate:      arg.ShowLocate,
			ShowRate:        arg.ShowRate,
		})

		if err != nil {
			return err
		}
		result = domain.AccountResponses{}

		// アカウントＩＤからテックタグのレコードを登録する
		for _, techTag := range arg.AccountTechTag {
			accountTag, err := q.CreateAccountTags(ctx, CreateAccountTagsParams{
				UserID:    arg.UserID,
				TechTagID: techTag,
			})
			if err != nil {
				return err
			}
			techtag, err := q.GetTechTagByID(ctx, accountTag.TechTagID)
			if err != nil {
				return err
			}
			result.AccountTechTags = append(result.AccountTechTags, techtag)
		}
		// アカウントＩＤからフレームワークのレコードを登録する
		for _, accountFrameworkTag := range arg.AccountFrameworkTag {
			accountFramework, err := q.CreateAccountFramework(ctx, CreateAccountFrameworkParams{
				UserID:      arg.UserID,
				FrameworkID: accountFrameworkTag,
			})
			if err != nil {
				return err
			}
			framework, err := q.GetFrameworksByID(ctx, accountFramework.FrameworkID)
			if err != nil {
				return err
			}
			result.AccountFrameworks = append(result.AccountFrameworks, framework)
		}

		return nil
	})
	return result, err
}
