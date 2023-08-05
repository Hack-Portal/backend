package db

import (
	"context"
	"database/sql"
	"fmt"

	"cloud.google.com/go/firestore"
	fb "firebase.google.com/go"
	"github.com/hackhack-Geek-vol6/backend/bootstrap"
	"github.com/hackhack-Geek-vol6/backend/domain"
)

const (
	ForeignKeyViolation = "foreign_key_violation"
	UniqueViolation     = "unique_violation"
)

type Store interface {
	Querier
	CreateAccountTx(ctx context.Context, arg CreateAccountTxParams) (domain.AccountResponses, error)
	GetAccountTxByID(ctx context.Context, ID string) (domain.AccountResponses, error)
	GetAccountTxByEmail(ctx context.Context, email string) (domain.AccountResponses, error)
	UpdateAccountTx(ctx context.Context, arg UpdateAccountTxParams) (domain.AccountResponses, error)

	CreateRoomTx(ctx context.Context, arg CreateRoomTxParams) (CreateRoomTxResult, error)
	CreateHackathonTx(ctx context.Context, config *bootstrap.Env, arg CreateHackathonTxParams) (CreateHackathonTxResult, error)
	ListRoomTx(ctx context.Context, arg ListRoomTxParam) ([]ListRoomTxResult, error)
	// Firebase
	InitChatRoom(ctx context.Context, roomID string) (*firestore.WriteResult, error)
	WriteFireStore(ctx context.Context, arg WriteFireStoreParam) (*firestore.WriteResult, error)
	ReadDocsByRoomID(ctx context.Context, RoomID string) (map[string]ChatRoomsWrite, error)
	UploadImage(ctx context.Context, file []byte) (string, error)
}
type SQLStore struct {
	*Queries
	db  *sql.DB
	App *fb.App
}

func NewStore(db *sql.DB, app *fb.App) *SQLStore {
	return &SQLStore{
		db:      db,
		Queries: New(db),
		App:     app,
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
