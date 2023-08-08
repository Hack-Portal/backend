package transaction

import (
	"context"
	"database/sql"
	"fmt"

	"cloud.google.com/go/firestore"
	fb "firebase.google.com/go"
	"github.com/google/uuid"
	"github.com/hackhack-Geek-vol6/backend/pkg/domain"
	repository "github.com/hackhack-Geek-vol6/backend/pkg/gateways/repository/datasource"
)

const (
	ForeignKeyViolation = "foreign_key_violation"
	UniqueViolation     = "unique_violation"
)

type Store interface {
	repository.Querier
	// Account Tx
	CreateAccountTx(ctx context.Context, args domain.CreateAccountParams) (domain.AccountResponses, error)
	GetAccountTxByID(ctx context.Context, ID string) (domain.AccountResponses, error)
	GetAccountTxByEmail(ctx context.Context, email string) (domain.AccountResponses, error)
	UpdateAccountTx(ctx context.Context, args domain.UpdateAccountParam) (domain.AccountResponses, error)
	// Room Tx
	CreateRoomTx(ctx context.Context, args domain.CreateRoomParam) (domain.GetRoomResponse, error)
	GetRoomTx(ctx context.Context, id uuid.UUID) (domain.GetRoomResponse, error)
	ListRoomTx(ctx context.Context, query domain.ListRoomsRequest) ([]domain.ListRoomResponse, error)
	UpdateRoomTx(ctx context.Context, body domain.UpdateRoomParam) (domain.GetRoomResponse, error)
	DeleteRoomTx(ctx context.Context, args domain.DeleteRoomParam) error
	AddAccountInRoom(ctx context.Context, args domain.AddAccountInRoomParam) error

	// Firebase
	InitChatRoom(ctx context.Context, roomID string) (*firestore.WriteResult, error)
	WriteFireStore(ctx context.Context, arg domain.WriteFireStoreParam) (*firestore.WriteResult, error)
	ReadDocsByRoomID(ctx context.Context, RoomID string) (map[string]domain.ChatRoomsWrite, error)
	UploadImage(ctx context.Context, file []byte) (string, error)
}
type SQLStore struct {
	*repository.Queries
	db  *sql.DB
	App *fb.App
}

func NewStore(db *sql.DB, app *fb.App) *SQLStore {
	return &SQLStore{
		db:      db,
		Queries: repository.New(db),
		App:     app,
	}
}

// トランザクションを実行する用の雛形
func (store *SQLStore) execTx(ctx context.Context, fn func(*repository.Queries) error) error {
	tx, err := store.db.BeginTx(ctx, &sql.TxOptions{})
	if err != nil {
		return err
	}
	q := repository.New(tx)

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
