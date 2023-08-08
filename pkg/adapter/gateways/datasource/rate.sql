-- name: CreateRateEntries :one
INSERT INTO rate_entries (user_id, rate)
VALUES($1, $2)
RETURNING *;
-- name: ListRateEntries :many
SELECT *
FROM rate_entries
WHERE user_id = $1
LIMIT $2 OFFSET $3;