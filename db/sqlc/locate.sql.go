// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.19.1
// source: locate.sql

package db

import (
	"context"
)

const getLocateByID = `-- name: GetLocateByID :one
SELECT locate_id, name FROM locates WHERE locate_id = $1
`

func (q *Queries) GetLocateByID(ctx context.Context, locateID int32) (Locates, error) {
	row := q.db.QueryRowContext(ctx, getLocateByID, locateID)
	var i Locates
	err := row.Scan(&i.LocateID, &i.Name)
	return i, err
}

const listLocates = `-- name: ListLocates :many
SELECT locate_id, name FROM locates
`

func (q *Queries) ListLocates(ctx context.Context) ([]Locates, error) {
	rows, err := q.db.QueryContext(ctx, listLocates)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Locates{}
	for rows.Next() {
		var i Locates
		if err := rows.Scan(&i.LocateID, &i.Name); err != nil {
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
