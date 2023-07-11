// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.18.0
// source: account.sql

package db

import (
	"context"
	"database/sql"
	"time"
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
	Icon            []byte         `json:"icon"`
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

const getAccountByID = `-- name: GetAccountByID :one
SELECT 
    user_id,
    username,
    icon,
    explanatory_text,
    (
        SELECT 
            name 
        FROM 
            locates 
        WHERE 
            locate_id = accounts.locate_id
    ) as locate,
    rate,
    hashed_password,
    email,
    show_locate,
    show_rate,
    create_at,
    update_at
FROM
    accounts
WHERE
    user_id = $1
`

type GetAccountByIDRow struct {
	UserID          string         `json:"user_id"`
	Username        string         `json:"username"`
	Icon            []byte         `json:"icon"`
	ExplanatoryText sql.NullString `json:"explanatory_text"`
	Locate          string         `json:"locate"`
	Rate            int32          `json:"rate"`
	HashedPassword  sql.NullString `json:"hashed_password"`
	Email           string         `json:"email"`
	ShowLocate      bool           `json:"show_locate"`
	ShowRate        bool           `json:"show_rate"`
	CreateAt        time.Time      `json:"create_at"`
	UpdateAt        time.Time      `json:"update_at"`
}

func (q *Queries) GetAccountByID(ctx context.Context, userID string) (GetAccountByIDRow, error) {
	row := q.db.QueryRowContext(ctx, getAccountByID, userID)
	var i GetAccountByIDRow
	err := row.Scan(
		&i.UserID,
		&i.Username,
		&i.Icon,
		&i.ExplanatoryText,
		&i.Locate,
		&i.Rate,
		&i.HashedPassword,
		&i.Email,
		&i.ShowLocate,
		&i.ShowRate,
		&i.CreateAt,
		&i.UpdateAt,
	)
	return i, err
}

const getAccountbyEmail = `-- name: GetAccountbyEmail :one
SELECT 
    user_id,
    username,
    icon,
    explanatory_text,
    (
        SELECT 
            name 
        FROM 
            locates 
        WHERE 
            locate_id = accounts.locate_id
    ) as locate,
    rate,
    hashed_password,
    email,
    show_locate,
    show_rate,
    create_at,
    update_at
FROM
    accounts
WHERE
    email = $1
`

type GetAccountbyEmailRow struct {
	UserID          string         `json:"user_id"`
	Username        string         `json:"username"`
	Icon            []byte         `json:"icon"`
	ExplanatoryText sql.NullString `json:"explanatory_text"`
	Locate          string         `json:"locate"`
	Rate            int32          `json:"rate"`
	HashedPassword  sql.NullString `json:"hashed_password"`
	Email           string         `json:"email"`
	ShowLocate      bool           `json:"show_locate"`
	ShowRate        bool           `json:"show_rate"`
	CreateAt        time.Time      `json:"create_at"`
	UpdateAt        time.Time      `json:"update_at"`
}

func (q *Queries) GetAccountbyEmail(ctx context.Context, email string) (GetAccountbyEmailRow, error) {
	row := q.db.QueryRowContext(ctx, getAccountbyEmail, email)
	var i GetAccountbyEmailRow
	err := row.Scan(
		&i.UserID,
		&i.Username,
		&i.Icon,
		&i.ExplanatoryText,
		&i.Locate,
		&i.Rate,
		&i.HashedPassword,
		&i.Email,
		&i.ShowLocate,
		&i.ShowRate,
		&i.CreateAt,
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
	UserID     string `json:"user_id"`
	Username   string `json:"username"`
	Icon       []byte `json:"icon"`
	Locate     string `json:"locate"`
	Rate       int32  `json:"rate"`
	ShowLocate bool   `json:"show_locate"`
	ShowRate   bool   `json:"show_rate"`
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
