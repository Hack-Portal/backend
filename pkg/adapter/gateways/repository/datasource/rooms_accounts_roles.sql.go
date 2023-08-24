// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.19.1
// source: rooms_accounts_roles.sql

package repository

import (
	"context"
)

const createRoomsAccountsRoles = `-- name: CreateRoomsAccountsRoles :one
INSERT INTO rooms_accounts_roles (rooms_account_id, role_id)
VALUES ($1, $2)
RETURNING rooms_account_id, role_id
`

type CreateRoomsAccountsRolesParams struct {
	RoomsAccountID int32 `json:"rooms_account_id"`
	RoleID         int32 `json:"role_id"`
}

func (q *Queries) CreateRoomsAccountsRoles(ctx context.Context, arg CreateRoomsAccountsRolesParams) (RoomsAccountsRole, error) {
	row := q.db.QueryRowContext(ctx, createRoomsAccountsRoles, arg.RoomsAccountID, arg.RoleID)
	var i RoomsAccountsRole
	err := row.Scan(&i.RoomsAccountID, &i.RoleID)
	return i, err
}

const deleteRoomsAccountsRolesByID = `-- name: DeleteRoomsAccountsRolesByID :exec
DELETE FROM rooms_accounts_roles
WHERE rooms_account_id = $1
`

func (q *Queries) DeleteRoomsAccountsRolesByID(ctx context.Context, roomsAccountID int32) error {
	_, err := q.db.ExecContext(ctx, deleteRoomsAccountsRolesByID, roomsAccountID)
	return err
}

const listRoomsAccountsRolesByID = `-- name: ListRoomsAccountsRolesByID :many
SELECT rooms_account_id, role_id
FROM rooms_accounts_roles
WHERE rooms_account_id = $1
LIMIT $2 OFFSET $3
`

type ListRoomsAccountsRolesByIDParams struct {
	RoomsAccountID int32 `json:"rooms_account_id"`
	Limit          int32 `json:"limit"`
	Offset         int32 `json:"offset"`
}

func (q *Queries) ListRoomsAccountsRolesByID(ctx context.Context, arg ListRoomsAccountsRolesByIDParams) ([]RoomsAccountsRole, error) {
	rows, err := q.db.QueryContext(ctx, listRoomsAccountsRolesByID, arg.RoomsAccountID, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []RoomsAccountsRole{}
	for rows.Next() {
		var i RoomsAccountsRole
		if err := rows.Scan(&i.RoomsAccountID, &i.RoleID); err != nil {
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
