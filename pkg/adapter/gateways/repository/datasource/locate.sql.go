// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.19.1
// source: locate.sql

package repository

import (
	"context"
)

const getLocatesByID = `-- name: GetLocatesByID :one
SELECT locate_id, name FROM locates WHERE locate_id = $1
`

func (q *Queries) GetLocatesByID(ctx context.Context, locateID int32) (Locate, error) {
	row := q.db.QueryRowContext(ctx, getLocatesByID, locateID)
	var i Locate
	err := row.Scan(&i.LocateID, &i.Name)
	return i, err
}

const listLocates = `-- name: ListLocates :many
SELECT locate_id, name FROM locates
`

func (q *Queries) ListLocates(ctx context.Context) ([]Locate, error) {
	rows, err := q.db.QueryContext(ctx, listLocates)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Locate{}
	for rows.Next() {
		var i Locate
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
