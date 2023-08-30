
-- name: CreatePastWorkFrameworks :one
INSERT INTO past_work_frameworks (opus, framework_id)
VALUES ($1, $2)
RETURNING *;

-- name: ListPastWorkFrameworksByOpus :many
SELECT *
FROM past_work_frameworks
WHERE opus = $1;

-- name: DeletePastWorkFrameworksByOpus :exec
DELETE FROM past_work_frameworks
WHERE opus = $1;