-- name: CreateTechTags :one
INSERT INTO tech_tags (language) VALUES ($1) RETURNING * ;

-- name: GetTechTagsByID :one
SELECT *
FROM tech_tags
WHERE tech_tag_id = $1;

-- name: ListTechTags :many
SELECT *
FROM tech_tags;

-- name: DeleteTechTagsByID :exec
DELETE FROM tech_tags
WHERE tech_tag_id = $1;

-- name: UpdateTechTagsByID :one
UPDATE tech_tags
SET language = $2
WHERE tech_tag_id = $1
RETURNING *;