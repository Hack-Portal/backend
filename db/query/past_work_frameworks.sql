-- name: CreatePastWorkFrameworks :one
INSERT INTO past_work_frameworks (opus, framework_id)
VALUES ($1, $2)
RETURNING *;
-- name: GetPastWorkFrameworksByOpus :many
SELECT *
FROM past_work_frameworks
WHERE opus = $1;
-- name: ListPastWorkFrameworks :many
SELECT *
FROM past_work_frameworks;
-- name: DeletePastWorkFrameworksByOpus :exec
DELETE FROM past_work_frameworks
WHERE opus = $1;