-- name: CreateRateEntities :one
INSERT INTO rate_entities (account_id, rate)
VALUES($1, $2)
RETURNING *;

-- name: ListRateEntities :many
SELECT *
FROM rate_entities
WHERE account_id = $1
LIMIT $2 OFFSET $3;