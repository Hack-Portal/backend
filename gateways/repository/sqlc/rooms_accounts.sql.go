// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.19.0
// source: rooms_accounts.sql

package db

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

func (q *Queries) CreateRoomsAccounts(ctx context.Context, arg CreateRoomsAccountsParams) (RoomsAccounts, error) {
	row := q.db.QueryRowContext(ctx, createRoomsAccounts, arg.UserID, arg.RoomID, arg.IsOwner)
	var i RoomsAccounts
	err := row.Scan(
		&i.UserID,
		&i.RoomID,
		&i.IsOwner,
		&i.CreateAt,
	)
	return i, err
}

const getRoomsAccountsByRoomID = `-- name: GetRoomsAccountsByRoomID :many
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

type GetRoomsAccountsByRoomIDRow struct {
	UserID  sql.NullString `json:"user_id"`
	Icon    sql.NullString `json:"icon"`
	IsOwner bool           `json:"is_owner"`
}

func (q *Queries) GetRoomsAccountsByRoomID(ctx context.Context, roomID uuid.UUID) ([]GetRoomsAccountsByRoomIDRow, error) {
	rows, err := q.db.QueryContext(ctx, getRoomsAccountsByRoomID, roomID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []GetRoomsAccountsByRoomIDRow{}
	for rows.Next() {
		var i GetRoomsAccountsByRoomIDRow
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

const removeAccountInRoom = `-- name: RemoveAccountInRoom :exec
DELETE FROM rooms_accounts WHERE room_id = $1 AND user_id = $2
`

type RemoveAccountInRoomParams struct {
	RoomID uuid.UUID `json:"room_id"`
	UserID string    `json:"user_id"`
}

func (q *Queries) RemoveAccountInRoom(ctx context.Context, arg RemoveAccountInRoomParams) error {
	_, err := q.db.ExecContext(ctx, removeAccountInRoom, arg.RoomID, arg.UserID)
	return err
}