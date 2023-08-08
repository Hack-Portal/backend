-- name: ListLocates :many
SELECT * FROM locates;

-- name: GetLocatesByID :one
SELECT * FROM locates WHERE locate_id = $1;