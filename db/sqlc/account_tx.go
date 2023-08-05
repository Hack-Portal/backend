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

		locate, err := q.GetLocateByID(ctx, account.LocateID)
		if err != nil {
			return err
		}

		result = domain.AccountResponses{
			UserID:          account.UserID,
			Username:        account.Username,
			Icon:            account.Icon.String,
			ExplanatoryText: account.ExplanatoryText.String,
			Rate:            account.Rate,
			Email:           account.Email,
			Locate:          locate.Name,
			ShowLocate:      account.ShowLocate,
			ShowRate:        account.ShowRate,
			CreateAt:        account.CreateAt,
		}

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
			result.TechTags = append(result.TechTags, techtag)
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
			result.Frameworks = append(result.Frameworks, framework)
		}
		return nil
	})

	return result, err
}

func (store *SQLStore) GetAccountTxByID(ctx context.Context, id string) (domain.AccountResponses, error) {
	var result domain.AccountResponses
	err := store.execTx(ctx, func(q *Queries) error {
		var err error
		// アカウントを登録する
		account, err := q.GetAccountByID(ctx, id)
		if err != nil {
			return err
		}

		locate, err := q.GetLocateByID(ctx, account.LocateID)
		if err != nil {
			return err
		}

		result = domain.AccountResponses{
			UserID:          account.UserID,
			Username:        account.Username,
			Icon:            account.Icon.String,
			ExplanatoryText: account.ExplanatoryText.String,
			Rate:            account.Rate,
			Email:           account.Email,
			Locate:          locate.Name,
			ShowLocate:      account.ShowLocate,
			ShowRate:        account.ShowRate,
			CreateAt:        account.CreateAt,
		}

		techTags, err := q.ListAccountTagsByUserID(ctx, id)
		if err != nil {
			return err
		}
		for _, techTag := range techTags {
			techtag, err := q.GetTechTagByID(ctx, techTag.TechTagID.Int32)
			if err != nil {
				return err
			}
			result.TechTags = append(result.TechTags, techtag)
		}

		frameworks, err := q.ListAccountFrameworksByUserID(ctx, id)
		if err != nil {
			return nil
		}
		for _, framework := range frameworks {
			techtag, err := q.GetFrameworksByID(ctx, framework.FrameworkID.Int32)
			if err != nil {
				return err
			}
			result.Frameworks = append(result.Frameworks, techtag)
		}
		return nil
	})

	return result, err
}

func (store *SQLStore) GetAccountTxByEmail(ctx context.Context, email string) (domain.AccountResponses, error) {
	var result domain.AccountResponses
	err := store.execTx(ctx, func(q *Queries) error {
		var err error
		// アカウントを登録する
		account, err := q.GetAccountByEmail(ctx, email)
		if err != nil {
			return err
		}

		locate, err := q.GetLocateByID(ctx, account.LocateID)
		if err != nil {
			return err
		}

		result = domain.AccountResponses{
			UserID:          account.UserID,
			Username:        account.Username,
			Icon:            account.Icon.String,
			ExplanatoryText: account.ExplanatoryText.String,
			Rate:            account.Rate,
			Email:           account.Email,
			Locate:          locate.Name,
			ShowLocate:      account.ShowLocate,
			ShowRate:        account.ShowRate,
			CreateAt:        account.CreateAt,
		}

		techTags, err := q.ListAccountTagsByUserID(ctx, account.UserID)
		if err != nil {
			return err
		}
		for _, techTag := range techTags {
			techtag, err := q.GetTechTagByID(ctx, techTag.TechTagID.Int32)
			if err != nil {
				return err
			}
			result.TechTags = append(result.TechTags, techtag)
		}

		frameworks, err := q.ListAccountFrameworksByUserID(ctx, account.UserID)
		if err != nil {
			return nil
		}
		for _, framework := range frameworks {
			techtag, err := q.GetFrameworksByID(ctx, framework.FrameworkID.Int32)
			if err != nil {
				return err
			}
			result.Frameworks = append(result.Frameworks, techtag)
		}
		return nil
	})

	return result, err
}

type UpdateAccountTxParams struct {
	// ユーザ登録部分
	UpdateAccountParams
	// tech_tag登録用
	AccountTechTag []int32
	// FrameworkTag登録用
	AccountFrameworkTag []int32
}

func (store *SQLStore) UpdateAccountTx(ctx context.Context, arg UpdateAccountTxParams) (domain.AccountResponses, error) {
	var result domain.AccountResponses
	err := store.execTx(ctx, func(q *Queries) error {
		var err error
		// アカウントを登録する
		account, err := q.UpdateAccount(ctx, arg.UpdateAccountParams)
		if err != nil {
			return err
		}

		locate, err := q.GetLocateByID(ctx, account.LocateID)
		if err != nil {
			return err
		}

		result = domain.AccountResponses{
			UserID:          account.UserID,
			Username:        account.Username,
			Icon:            account.Icon.String,
			ExplanatoryText: account.ExplanatoryText.String,
			Rate:            account.Rate,
			Email:           account.Email,
			Locate:          locate.Name,
			ShowLocate:      account.ShowLocate,
			ShowRate:        account.ShowRate,
			CreateAt:        account.CreateAt,
		}

		err = q.DeleteAccounttagsByUserID(ctx, account.UserID)
		if err != nil {
			return err
		}

		err = q.DeleteAccountFrameworksByUserID(ctx, account.UserID)
		if err != nil {
			return err
		}

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
			result.TechTags = append(result.TechTags, techtag)
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
			result.Frameworks = append(result.Frameworks, framework)
		}
		return nil
	})

	return result, err
}
