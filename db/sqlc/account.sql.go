// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.18.0
// source: account.sql

package db

import (
	"context"
	"database/sql"
)

const createAccount = `-- name: CreateAccount :one
INSERT INTO accounts (
    user_id,
    username,
    icon,
    explanatory_text,
    locate_id,
    rate,
    hashed_password,
    email,
    show_locate,
    show_rate
)VALUES(
    $1,$2,$3,$4,$5,$6,$7,$8,$9,$10
)RETURNING user_id, username, icon, explanatory_text, locate_id, rate, hashed_password, email, create_at, show_locate, show_rate, update_at
`

type CreateAccountParams struct {
	UserID          string         `json:"user_id"`
	Username        string         `json:"username"`
	Icon            sql.NullString `json:"icon"`
	ExplanatoryText sql.NullString `json:"explanatory_text"`
	LocateID        int32          `json:"locate_id"`
	Rate            int32          `json:"rate"`
	HashedPassword  sql.NullString `json:"hashed_password"`
	Email           string         `json:"email"`
	ShowLocate      bool           `json:"show_locate"`
	ShowRate        bool           `json:"show_rate"`
}

func (q *Queries) CreateAccount(ctx context.Context, arg CreateAccountParams) (Accounts, error) {
	row := q.db.QueryRowContext(ctx, createAccount,
		arg.UserID,
		arg.Username,
		arg.Icon,
		arg.ExplanatoryText,
		arg.LocateID,
		arg.Rate,
		arg.HashedPassword,
		arg.Email,
		arg.ShowLocate,
		arg.ShowRate,
	)
	var i Accounts
	err := row.Scan(
		&i.UserID,
		&i.Username,
		&i.Icon,
		&i.ExplanatoryText,
		&i.LocateID,
		&i.Rate,
		&i.HashedPassword,
		&i.Email,
		&i.CreateAt,
		&i.ShowLocate,
		&i.ShowRate,
		&i.UpdateAt,
	)
	return i, err
}

const getAccount = `-- name: GetAccount :one
SELECT 
    user_id, username, icon, explanatory_text, locate_id, rate, hashed_password, email, create_at, show_locate, show_rate, update_at
FROM
    accounts
WHERE
    user_id = $1
`

func (q *Queries) GetAccount(ctx context.Context, userID string) (Accounts, error) {
	row := q.db.QueryRowContext(ctx, getAccount, userID)
	var i Accounts
	err := row.Scan(
		&i.UserID,
		&i.Username,
		&i.Icon,
		&i.ExplanatoryText,
		&i.LocateID,
		&i.Rate,
		&i.HashedPassword,
		&i.Email,
		&i.CreateAt,
		&i.ShowLocate,
		&i.ShowRate,
		&i.UpdateAt,
	)
	return i, err
}

const getAccountAuth = `-- name: GetAccountAuth :one
SELECT 
    user_id,
    hashed_password,
    email
FROM 
    accounts
WHERE
    user_id = $1
`

type GetAccountAuthRow struct {
	UserID         string         `json:"user_id"`
	HashedPassword sql.NullString `json:"hashed_password"`
	Email          string         `json:"email"`
}

func (q *Queries) GetAccountAuth(ctx context.Context, userID string) (GetAccountAuthRow, error) {
	row := q.db.QueryRowContext(ctx, getAccountAuth, userID)
	var i GetAccountAuthRow
	err := row.Scan(&i.UserID, &i.HashedPassword, &i.Email)
	return i, err
}

const getAccountbyEmail = `-- name: GetAccountbyEmail :one
SELECT 
    user_id, username, icon, explanatory_text, locate_id, rate, hashed_password, email, create_at, show_locate, show_rate, update_at
FROM
    accounts
WHERE
    email = $1
`

func (q *Queries) GetAccountbyEmail(ctx context.Context, email string) (Accounts, error) {
	row := q.db.QueryRowContext(ctx, getAccountbyEmail, email)
	var i Accounts
	err := row.Scan(
		&i.UserID,
		&i.Username,
		&i.Icon,
		&i.ExplanatoryText,
		&i.LocateID,
		&i.Rate,
		&i.HashedPassword,
		&i.Email,
		&i.CreateAt,
		&i.ShowLocate,
		&i.ShowRate,
		&i.UpdateAt,
	)
	return i, err
}

const listAccounts = `-- name: ListAccounts :many
SELECT
    user_id,
    username,
    icon,
    (
        SELECT 
            name 
        FROM 
            locates 
        WHERE 
            locate_id = accounts.locate_id
    ) as locate,
    rate,
    show_locate,
    show_rate
FROM
    accounts
WHERE username LIKE $1
LIMIT $2
OFFSET $3
`

type ListAccountsParams struct {
	Username string `json:"username"`
	Limit    int32  `json:"limit"`
	Offset   int32  `json:"offset"`
}

type ListAccountsRow struct {
	UserID     string         `json:"user_id"`
	Username   string         `json:"username"`
	Icon       sql.NullString `json:"icon"`
	Locate     string         `json:"locate"`
	Rate       int32          `json:"rate"`
	ShowLocate bool           `json:"show_locate"`
	ShowRate   bool           `json:"show_rate"`
}

func (q *Queries) ListAccounts(ctx context.Context, arg ListAccountsParams) ([]ListAccountsRow, error) {
	rows, err := q.db.QueryContext(ctx, listAccounts, arg.Username, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []ListAccountsRow{}
	for rows.Next() {
		var i ListAccountsRow
		if err := rows.Scan(
			&i.UserID,
			&i.Username,
			&i.Icon,
			&i.Locate,
			&i.Rate,
			&i.ShowLocate,
			&i.ShowRate,
		); err != nil {
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
