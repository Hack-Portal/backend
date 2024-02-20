
-- name: CreatePastWorkTags :one
INSERT INTO past_work_tags (opus, tech_tag_id)
VALUES($1, $2)
RETURNING *;

-- name: ListPastWorkTagsByOpus :many
SELECT *
FROM past_work_tags
WHERE opus = $1;

-- name: DeletePastWorkTagsByOpus :exec
DELETE FROM past_work_tags
WHERE opus = $1;