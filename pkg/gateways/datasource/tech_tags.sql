-- name: GetTechTagByID :one
SELECT * FROM tech_tags WHERE tech_tag_id = $1;

-- name: ListTechTag :many
SELECT * FROM tech_tags ;

-- name: DeleteTechTagByID :exec
DELETE FROM tech_tags WHERE tech_tag_id = $1;

-- name: UpdateTechTagByID :one
UPDATE tech_tags SET language = $1 WHERE tech_tag_id = $1 RETURNING *;