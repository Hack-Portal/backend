-- name: ListFrameworks :many
SELECT *
FROM frameworks;

-- name: GetFrameworksByID :one
SELECT *
FROM frameworks
WHERE framework_id = $1;

-- name: DeleteFrameworksByID :exec
DELETE FROM frameworks
WHERE framework_id = $1;

-- name: UpdateFrameworksByID :one
UPDATE frameworks
SET framework = $1,
    tech_tag_id = $2
WHERE framework_id = $3
RETURNING *;

-- name: CreateFrameworks :one
INSERT INTO
  frameworks (framework, tech_tag_id)
VALUES
($1, $2) RETURNING *;