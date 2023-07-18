-- name: ListFrameworks :many
SELECT
    *
FROM
    frameworks
LIMIT
    $1;

-- name: GetFrameworksByID :one
SELECT
    *
FROM
    frameworks
WHERE
    framework_id = $1;

-- name: DeleteFrameworksByID :exec
DELETE FROM
    frameworks
WHERE
    framework_id = $1;

-- name: UpdateFrameworksByID :one
UPDATE frameworks SET framework = $1 , tech_tag_id = $2 WHERE framework_id = $3 RETURNING *;