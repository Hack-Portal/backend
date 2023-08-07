// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.19.1
// source: past_work_tags.sql

package db

import (
	"context"
)

const createPastWorkTag = `-- name: CreatePastWorkTag :one
INSERT INTO past_work_tags (opus, tech_tag_id)
VALUES($1, $2)
RETURNING opus, tech_tag_id
`

type CreatePastWorkTagParams struct {
	Opus      int32 `json:"opus"`
	TechTagID int32 `json:"tech_tag_id"`
}

func (q *Queries) CreatePastWorkTag(ctx context.Context, arg CreatePastWorkTagParams) (PastWorkTags, error) {
	row := q.db.QueryRowContext(ctx, createPastWorkTag, arg.Opus, arg.TechTagID)
	var i PastWorkTags
	err := row.Scan(&i.Opus, &i.TechTagID)
	return i, err
}

const deletePastWorkTagsByOpus = `-- name: DeletePastWorkTagsByOpus :exec
DELETE FROM past_work_tags
WHERE opus = $1
`

func (q *Queries) DeletePastWorkTagsByOpus(ctx context.Context, opus int32) error {
	_, err := q.db.ExecContext(ctx, deletePastWorkTagsByOpus, opus)
	return err
}

const getPastWorkTagsByOpus = `-- name: GetPastWorkTagsByOpus :many
SELECT opus, tech_tag_id
FROM past_work_tags
WHERE opus = $1
`

func (q *Queries) GetPastWorkTagsByOpus(ctx context.Context, opus int32) ([]PastWorkTags, error) {
	rows, err := q.db.QueryContext(ctx, getPastWorkTagsByOpus, opus)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []PastWorkTags{}
	for rows.Next() {
		var i PastWorkTags
		if err := rows.Scan(&i.Opus, &i.TechTagID); err != nil {
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

const listPastWorkTags = `-- name: ListPastWorkTags :many
SELECT opus, tech_tag_id
FROM past_work_tags
`

func (q *Queries) ListPastWorkTags(ctx context.Context) ([]PastWorkTags, error) {
	rows, err := q.db.QueryContext(ctx, listPastWorkTags)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []PastWorkTags{}
	for rows.Next() {
		var i PastWorkTags
		if err := rows.Scan(&i.Opus, &i.TechTagID); err != nil {
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
