// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.19.1
// source: rooms_accounts.sql

package repository

import (
	"context"
	"database/sql"

	"github.com/google/uuid"
)

const createRoomsAccounts = `-- name: CreateRoomsAccounts :one
INSERT INTO rooms_accounts (
    user_id,
    room_id,
    is_owner
)VALUES(
    $1,$2,$3
)RETURNING user_id, room_id, is_owner, create_at
`

type CreateRoomsAccountsParams struct {
	UserID  string    `json:"user_id"`
	RoomID  uuid.UUID `json:"room_id"`
	IsOwner bool      `json:"is_owner"`
}

func (q *Queries) CreateRoomsAccounts(ctx context.Context, arg CreateRoomsAccountsParams) (RoomsAccount, error) {
	row := q.db.QueryRowContext(ctx, createRoomsAccounts, arg.UserID, arg.RoomID, arg.IsOwner)
	var i RoomsAccount
	err := row.Scan(
		&i.UserID,
		&i.RoomID,
		&i.IsOwner,
		&i.CreateAt,
	)
	return i, err
}

const deleteRoomsAccountsByID = `-- name: DeleteRoomsAccountsByID :exec
DELETE FROM rooms_accounts WHERE room_id = $1 AND user_id = $2
`

type DeleteRoomsAccountsByIDParams struct {
	RoomID uuid.UUID `json:"room_id"`
	UserID string    `json:"user_id"`
}

func (q *Queries) DeleteRoomsAccountsByID(ctx context.Context, arg DeleteRoomsAccountsByIDParams) error {
	_, err := q.db.ExecContext(ctx, deleteRoomsAccountsByID, arg.RoomID, arg.UserID)
	return err
}

const getRoomsAccountsByID = `-- name: GetRoomsAccountsByID :many
SELECT 
    accounts.user_id, 
    accounts.icon,
    rooms_accounts.is_owner
FROM 
    rooms_accounts
LEFT OUTER JOIN 
    accounts 
ON 
    rooms_accounts.user_id = accounts.user_id 
WHERE 
    rooms_accounts.room_id = $1
`

type GetRoomsAccountsByIDRow struct {
	UserID  sql.NullString `json:"user_id"`
	Icon    sql.NullString `json:"icon"`
	IsOwner bool           `json:"is_owner"`
}

func (q *Queries) GetRoomsAccountsByID(ctx context.Context, roomID uuid.UUID) ([]GetRoomsAccountsByIDRow, error) {
	rows, err := q.db.QueryContext(ctx, getRoomsAccountsByID, roomID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []GetRoomsAccountsByIDRow{}
	for rows.Next() {
		var i GetRoomsAccountsByIDRow
		if err := rows.Scan(&i.UserID, &i.Icon, &i.IsOwner); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
