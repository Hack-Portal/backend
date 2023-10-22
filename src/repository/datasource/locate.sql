-- name: CreateLocates :one
INSERT INTO locates (name) VALUES ($1) RETURNING *;

-- name: ListLocates :many
SELECT * FROM locates;

-- name: GetLocatesByID :one
SELECT * FROM locates WHERE locate_id = $1;

