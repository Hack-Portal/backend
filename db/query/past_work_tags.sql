-- name: CreatePastWorkTag :one
INSERT INTO past_work_tags (opus, tech_tag_id)
VALUES($1, $2)
RETURNING *;
-- name: GetPastWorkTagsByOpus :many
SELECT *
FROM past_work_tags
WHERE opus = $1;
-- name: ListPastWorkTags :many
SELECT *
FROM past_work_tags;
-- name: DeletePastWorkTagsByOpus :exec
DELETE FROM past_work_tags
WHERE opus = $1;