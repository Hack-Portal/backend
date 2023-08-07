-- name: CreateAccountPastWorks :one
INSERT INTO account_past_works (opus, user_id)
VALUES ($1, $2)
RETURNING *;
-- name: GetAccountPastWorksByOpus :many
SELECT *
FROM account_past_works
WHERE opus = $1;
-- name: ListAccountPastWorks :many
SELECT *
FROM account_past_works;
-- name: DeleteAccountPastWorksByOpus :exec
DELETE FROM account_past_works
WHERE opus = $1;