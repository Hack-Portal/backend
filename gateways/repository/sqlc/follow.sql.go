// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.19.0
// source: follow.sql

package db

import (
	"context"
)

const createFollow = `-- name: CreateFollow :one
INSERT INTO
  follows (to_user_id, from_user_id)
VALUES
($1, $2) RETURNING to_user_id, from_user_id, create_at
`

type CreateFollowParams struct {
	ToUserID   string `json:"to_user_id"`
	FromUserID string `json:"from_user_id"`
}

func (q *Queries) CreateFollow(ctx context.Context, arg CreateFollowParams) (Follows, error) {
	row := q.db.QueryRowContext(ctx, createFollow, arg.ToUserID, arg.FromUserID)
	var i Follows
	err := row.Scan(&i.ToUserID, &i.FromUserID, &i.CreateAt)
	return i, err
}

const listFollowByToUserID = `-- name: ListFollowByToUserID :many
SELECT
  to_user_id, from_user_id, create_at
FROM
  follows
WHERE
  to_user_id = $1
`

func (q *Queries) ListFollowByToUserID(ctx context.Context, toUserID string) ([]Follows, error) {
	rows, err := q.db.QueryContext(ctx, listFollowByToUserID, toUserID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Follows{}
	for rows.Next() {
		var i Follows
		if err := rows.Scan(&i.ToUserID, &i.FromUserID, &i.CreateAt); err != nil {
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

const removeFollow = `-- name: RemoveFollow :exec
DELETE FROM
  follows
WHERE
  to_user_id = $1
  AND from_user_id = $2
`

type RemoveFollowParams struct {
	ToUserID   string `json:"to_user_id"`
	FromUserID string `json:"from_user_id"`
}

func (q *Queries) RemoveFollow(ctx context.Context, arg RemoveFollowParams) error {
	_, err := q.db.ExecContext(ctx, removeFollow, arg.ToUserID, arg.FromUserID)
	return err
}
