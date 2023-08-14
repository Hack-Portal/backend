// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.19.1
// source: bookmark.sql

package repository

import (
	"context"
)

const createBookmarks = `-- name: CreateBookmarks :one
INSERT INTO
    bookmarks(opus, account_id)
VALUES
    ($1, $2) RETURNING opus, account_id, create_at, is_delete
`

type CreateBookmarksParams struct {
	Opus      int32  `json:"opus"`
	AccountID string `json:"account_id"`
}

func (q *Queries) CreateBookmarks(ctx context.Context, arg CreateBookmarksParams) (Bookmark, error) {
	row := q.db.QueryRowContext(ctx, createBookmarks, arg.Opus, arg.AccountID)
	var i Bookmark
	err := row.Scan(
		&i.Opus,
		&i.AccountID,
		&i.CreateAt,
		&i.IsDelete,
	)
	return i, err
}

const deleteBookmarksByID = `-- name: DeleteBookmarksByID :one
UPDATE
    bookmarks
SET
    is_delete = true
WHERE
    account_id = $1
    AND opus = $2 RETURNING opus, account_id, create_at, is_delete
`

type DeleteBookmarksByIDParams struct {
	AccountID string `json:"account_id"`
	Opus      int32  `json:"opus"`
}

func (q *Queries) DeleteBookmarksByID(ctx context.Context, arg DeleteBookmarksByIDParams) (Bookmark, error) {
	row := q.db.QueryRowContext(ctx, deleteBookmarksByID, arg.AccountID, arg.Opus)
	var i Bookmark
	err := row.Scan(
		&i.Opus,
		&i.AccountID,
		&i.CreateAt,
		&i.IsDelete,
	)
	return i, err
}

const listBookmarksByID = `-- name: ListBookmarksByID :many
SELECT
    opus, account_id, create_at, is_delete
FROM
    bookmarks
WHERE
    account_id = $1 AND is_delete = false
`

func (q *Queries) ListBookmarksByID(ctx context.Context, accountID string) ([]Bookmark, error) {
	rows, err := q.db.QueryContext(ctx, listBookmarksByID, accountID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Bookmark{}
	for rows.Next() {
		var i Bookmark
		if err := rows.Scan(
			&i.Opus,
			&i.AccountID,
			&i.CreateAt,
			&i.IsDelete,
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
