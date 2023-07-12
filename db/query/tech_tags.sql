-- name: GetTechTagByID :one
SELECT * FROM tech_tags WHERE tech_tag_id = $1;

-- name: ListTechTag :many
SELECT * FROM tech_tags ;