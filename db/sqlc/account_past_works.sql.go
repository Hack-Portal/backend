// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.19.1
// source: account_past_works.sql

package db

import (
	"context"
)

const createAccountPastWorks = `-- name: CreateAccountPastWorks :one
INSERT INTO account_past_works (opus, user_id)
VALUES ($1, $2)
RETURNING opus, user_id
`

type CreateAccountPastWorksParams struct {
	Opus   int32  `json:"opus"`
	UserID string `json:"user_id"`
}

func (q *Queries) CreateAccountPastWorks(ctx context.Context, arg CreateAccountPastWorksParams) (AccountPastWorks, error) {
	row := q.db.QueryRowContext(ctx, createAccountPastWorks, arg.Opus, arg.UserID)
	var i AccountPastWorks
	err := row.Scan(&i.Opus, &i.UserID)
	return i, err
}

const deleteAccountPastWorksByOpus = `-- name: DeleteAccountPastWorksByOpus :exec
DELETE FROM account_past_works
WHERE opus = $1
`

func (q *Queries) DeleteAccountPastWorksByOpus(ctx context.Context, opus int32) error {
	_, err := q.db.ExecContext(ctx, deleteAccountPastWorksByOpus, opus)
	return err
}

const getAccountPastWorksByOpus = `-- name: GetAccountPastWorksByOpus :many
SELECT opus, user_id
FROM account_past_works
WHERE opus = $1
`

func (q *Queries) GetAccountPastWorksByOpus(ctx context.Context, opus int32) ([]AccountPastWorks, error) {
	rows, err := q.db.QueryContext(ctx, getAccountPastWorksByOpus, opus)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []AccountPastWorks{}
	for rows.Next() {
		var i AccountPastWorks
		if err := rows.Scan(&i.Opus, &i.UserID); err != nil {
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

const listAccountPastWorks = `-- name: ListAccountPastWorks :many
SELECT opus, user_id
FROM account_past_works
`

func (q *Queries) ListAccountPastWorks(ctx context.Context) ([]AccountPastWorks, error) {
	rows, err := q.db.QueryContext(ctx, listAccountPastWorks)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []AccountPastWorks{}
	for rows.Next() {
		var i AccountPastWorks
		if err := rows.Scan(&i.Opus, &i.UserID); err != nil {
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