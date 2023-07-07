-- name: ListLocates :many
SELECT * FROM locates;

-- name: GetLocate :one
SELECT * FROM locates WHERE locate_id = $1;