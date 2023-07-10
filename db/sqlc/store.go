package db

import (
	"context"
	"database/sql"
	"fmt"
)

type Store interface {
	Querier
	CreateAccountTx(ctx context.Context, arg CreateAccountTxParams) (CreateAccountTxResult, error)
}

type SQLStore struct {
	*Queries
	db *sql.DB
}

func NewStore(db *sql.DB) *SQLStore {
	return &SQLStore{
		db:      db,
		Queries: New(db),
	}
}

// トランザクションを実行する用の雛形
func (store *SQLStore) execTx(ctx context.Context, fn func(*Queries) error) error {
	tx, err := store.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	q := New(tx)

	err = fn(q)
	if err != nil {
		//トランザクションにエラーが発生したときのロールバック処理
		if rbErr := tx.Rollback(); rbErr != nil {
			return fmt.Errorf("tx err : %v , rb err: %v", err, rbErr)
		}
		return err
	}
	return tx.Commit()
}

type CreateAccountTxParams struct {
	// ユーザ登録部分
	Accounts
	// tech_tag登録用
	AccountTechTag []int32
	// FrameworkTag登録用
	AccountFrameworkTag []int32
}
type CreateAccountTxResult struct {
	Account           Accounts
	AccountTechTags   []AccountTags
	AccountFrameworks []AccountFrameworks
}

// アカウント登録時のトランザクション
func (store *SQLStore) CreateAccountTx(ctx context.Context, arg CreateAccountTxParams) (CreateAccountTxResult, error) {
	var result CreateAccountTxResult
	err := store.execTx(ctx, func(q *Queries) error {
		var err error
		// アカウントを登録する
		result.Account, err = q.CreateAccount(ctx, CreateAccountParams{
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

		for _, techtag := range arg.AccountTechTag {
			accountTag, err := q.CreataAccountTags(ctx, CreataAccountTagsParams{
				UserID:    arg.UserID,
				TechTagID: techtag,
			})
			if err != nil {
				return err
			}
			result.AccountTechTags = append(result.AccountTechTags, accountTag)
		}

		for _, accountFrameworkTag := range arg.AccountFrameworkTag {
			accountFramework, err := q.CreateAccountFramework(ctx, CreateAccountFrameworkParams{
				AccountID:   arg.UserID,
				FrameworkID: accountFrameworkTag,
			})
			if err != nil {
				return err
			}
			result.AccountFrameworks = append(result.AccountFrameworks, accountFramework)
		}
		return nil
	})
	return result, err
}
