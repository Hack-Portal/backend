// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.19.1
// source: rate.sql

package repository

import (
	"context"
)

const createRate = `-- name: CreateRate :one
INSERT INTO rate_entries (user_id, rate)
VALUES($1, $2)
RETURNING user_id, rate, create_at
`

type CreateRateParams struct {
	UserID string `json:"user_id"`
	Rate   int32  `json:"rate"`
}

func (q *Queries) CreateRate(ctx context.Context, arg CreateRateParams) (RateEntry, error) {
	row := q.db.QueryRowContext(ctx, createRate, arg.UserID, arg.Rate)
	var i RateEntry
	err := row.Scan(&i.UserID, &i.Rate, &i.CreateAt)
	return i, err
}

const listRate = `-- name: ListRate :many
SELECT user_id, rate, create_at
FROM rate_entries
WHERE user_id = $1
LIMIT $2 OFFSET $3
`

type ListRateParams struct {
	UserID string `json:"user_id"`
	Limit  int32  `json:"limit"`
	Offset int32  `json:"offset"`
}

func (q *Queries) ListRate(ctx context.Context, arg ListRateParams) ([]RateEntry, error) {
	rows, err := q.db.QueryContext(ctx, listRate, arg.UserID, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []RateEntry{}
	for rows.Next() {
		var i RateEntry
		if err := rows.Scan(&i.UserID, &i.Rate, &i.CreateAt); err != nil {
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