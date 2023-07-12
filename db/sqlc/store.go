package db

import (
	"context"
	"database/sql"
	"fmt"
)

type Store interface {
	Querier
	CreateAccountTx(ctx context.Context, arg CreateAccountTxParams) (CreateAccountTxResult, error)
	CreateRoomTx(ctx context.Context, arg CreateRoomTxParams) (CraeteRoomTxResult, error)
	CreateHackathonTx(ctx context.Context, arg CreateHackathonTxParams) (CreateHackathonTxResult, error)
	ListRoomTx(ctx context.Context, arg ListRoomTxParam) ([]ListRoomTxResult, error)
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
	AccountTechTags   []TechTags
	AccountFrameworks []Frameworks
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

		// アカウントＩＤからテックタグのレコードを登録する
		for _, techtag := range arg.AccountTechTag {
			accountTag, err := q.CreataAccountTags(ctx, CreataAccountTagsParams{
				UserID:    arg.UserID,
				TechTagID: techtag,
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
				AccountID:   arg.UserID,
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

type CreateHackathonTxParams struct {
	// ハッカソン登録部分
	Hackathons
	// status_tag登録用
	HackathonStatusTag []int32
}

type CreateHackathonTxResult struct {
	Hackathons
	HackathonStatusTag []StatusTags
}

// ハッカソン登録時のトランザクション
func (store *SQLStore) CreateHackathonTx(ctx context.Context, arg CreateHackathonTxParams) (CreateHackathonTxResult, error) {
	var result CreateHackathonTxResult
	err := store.execTx(ctx, func(q *Queries) error {
		var err error
		// ハッカソンを登録する
		result.Hackathons, err = q.CreateHackathon(ctx, CreateHackathonParams{
			Name:        arg.Name,
			Icon:        arg.Icon,
			Description: arg.Description,
			Link:        arg.Link,
			Expired:     arg.Expired,
			StartDate:   arg.StartDate,
			Term:        arg.Term,
		})
		if err != nil {
			return err
		}
		// ハッカソンIDからステータスタグのレコードを登録する
		for _, status_id := range arg.HackathonStatusTag {
			_, err = q.CreateHackathonStatusTag(ctx, CreateHackathonStatusTagParams{
				HackathonID: result.HackathonID,
				StatusID:    status_id,
			})
			if err != nil {
				return err
			}
		}
		statusTag, err := q.GetStatusTags(ctx, result.HackathonID)
		result.HackathonStatusTag = statusTag
		return nil
	})
	return result, err
}
