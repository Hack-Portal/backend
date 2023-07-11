-- name: ListFrameworks :many
SELECT * FROM frameworks LIMIT $1 ;
-- name: GetFrameworksByID :one
SELECT * FROM frameworks WHERE framework_id = $1 ;