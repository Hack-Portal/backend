package transaction

import (
	"context"
	"database/sql"
	"fmt"

	"cloud.google.com/go/firestore"
	fb "firebase.google.com/go"
	repository "github.com/hackhack-Geek-vol6/backend/pkg/adapter/gateways/repository/datasource"
	"github.com/hackhack-Geek-vol6/backend/pkg/domain/params"
)

const (
	ForeignKeyViolation = "foreign_key_violation"
	UniqueViolation     = "unique_violation"
)

type Store interface {
	repository.Querier
	// Account Tx
	CreateAccountTx(ctx context.Context, args params.CreateAccount) (repository.Account, error)
	UpdateAccountTx(ctx context.Context, args params.UpdateAccount) (repository.Account, error)
	// Room Tx
	CreateRoomTx(ctx context.Context, args params.CreateRoom) (repository.Room, error)
	UpdateRoomTx(ctx context.Context, body params.UpdateRoom) (repository.Room, error)
	DeleteRoomTx(ctx context.Context, args params.DeleteRoom) error
	AddAccountInRoom(ctx context.Context, args params.AddAccountInRoom) error
	CloseRoom(ctx context.Context, args params.CloseRoom) error
	AddRoomAccountRoleByID(ctx context.Context, args params.RoomAccountRole) error
	UpdateRoomsAccountRoleByID(ctx context.Context, args params.RoomAccountRole) error
	// Hackathon Tx
	CreateHackathonTx(ctx context.Context, args params.CreateHackathon) (repository.Hackathon, error)

	// PastWork Tx
	CreatePastWorkTx(ctx context.Context, arg params.CreatePastWork) (repository.PastWork, error)
	UpdatePastWorkTx(ctx context.Context, arg params.UpdatePastWork) (repository.PastWork, error)

	// Rate Entities Tx
	CreateRateEntityTx(ctx context.Context, arg repository.CreateRateEntitiesParams) error

	// Firebase
	InitChatRoom(ctx context.Context, roomID string) (*firestore.WriteResult, error)
	CreateSubCollection(ctx context.Context, arg params.WriteFireStore) (*firestore.WriteResult, error)
	ReadDocsByRoomID(ctx context.Context, roomID string) (int, error)
	UploadImage(ctx context.Context, file []byte) (string, string, error)
	DeleteImage(ctx context.Context, file string) error
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
