package db

import (
	"context"
	"database/sql"
	"fmt"

	"cloud.google.com/go/firestore"
)

type Store interface {
	Querier
	CreateAccountTx(ctx context.Context, arg CreateAccountTxParams) (CreateAccountTxResult, error)
	CreateRoomTx(ctx context.Context, arg CreateRoomTxParams) (CraeteRoomTxResult, error)
	CreateHackathonTx(ctx context.Context, arg CreateHackathonTxParams) (CreateHackathonTxResult, error)
	ListRoomTx(ctx context.Context, arg ListRoomTxParam) ([]ListRoomTxResult, error)
	// Firebase
	InitChatRoom(ctx context.Context, roomID string) (*firestore.WriteResult, error)
	WriteFireStore(ctx context.Context, arg WriteFireStoreParam) (*firestore.WriteResult, error)
	ReadDocsByRoomID(ctx context.Context, RoomID string) (map[string]ChatRoomsWrite, error)
}

type SQLStore struct {
	*Queries
	db     *sql.DB
	client *firestore.Client
}

func NewStore(db *sql.DB, client *firestore.Client) *SQLStore {
	return &SQLStore{
		db:      db,
		Queries: New(db),
		client:  client,
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
		statusTag, err := q.GetStatusTagsByhackathonID(ctx, result.HackathonID)
		result.HackathonStatusTag = statusTag
		return nil
	})
	return result, err
}
