// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.20.0
// source: hackathon_status_tags.sql

package repository

import (
	"context"
)

const createHackathonStatusTags = `-- name: CreateHackathonStatusTags :one
INSERT INTO hackathon_status_tags (
    hackathon_id,
    status_id
  )VALUES(
    $1,
    $2
  )
RETURNING hackathon_id, status_id
`

type CreateHackathonStatusTagsParams struct {
	HackathonID int32 `json:"hackathon_id"`
	StatusID    int32 `json:"status_id"`
}

func (q *Queries) CreateHackathonStatusTags(ctx context.Context, arg CreateHackathonStatusTagsParams) (HackathonStatusTag, error) {
	row := q.db.QueryRowContext(ctx, createHackathonStatusTags, arg.HackathonID, arg.StatusID)
	var i HackathonStatusTag
	err := row.Scan(&i.HackathonID, &i.StatusID)
	return i, err
}

const deleteHackathonStatusTagsByID = `-- name: DeleteHackathonStatusTagsByID :exec
DELETE FROM hackathon_status_tags WHERE hackathon_id = $1
`

func (q *Queries) DeleteHackathonStatusTagsByID(ctx context.Context, hackathonID int32) error {
	_, err := q.db.ExecContext(ctx, deleteHackathonStatusTagsByID, hackathonID)
	return err
}

const listHackathonStatusTagsByID = `-- name: ListHackathonStatusTagsByID :many
SELECT hackathon_id, status_id
FROM hackathon_status_tags
WHERE hackathon_id = $1
`

func (q *Queries) ListHackathonStatusTagsByID(ctx context.Context, hackathonID int32) ([]HackathonStatusTag, error) {
	rows, err := q.db.QueryContext(ctx, listHackathonStatusTagsByID, hackathonID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []HackathonStatusTag{}
	for rows.Next() {
		var i HackathonStatusTag
		if err := rows.Scan(&i.HackathonID, &i.StatusID); err != nil {
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
