-- name: ListLocates :many
SELECT * FROM locates;

-- name: GetLocateByID :one
SELECT * FROM locates WHERE locate_id = $1;