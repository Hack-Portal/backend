-- name: CreateAccountPastWorks :one
INSERT INTO account_past_works (opus, account_id)
VALUES ($1, $2)
RETURNING *;
-- name: ListAccountPastWorksByOpus :many
SELECT *
FROM account_past_works
WHERE opus = $1;
-- name: DeleteAccountPastWorksByOpus :exec
DELETE FROM account_past_works
WHERE opus = $1;