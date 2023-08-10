-- name: CreateFollows :one
INSERT INTO
  follows (to_account_id, from_account_id)
VALUES
($1, $2) RETURNING *;

-- name: ListFollowsByToUserID :many
SELECT
  *
FROM
  follows
WHERE
  to_account_id = $1;

-- name: DeleteFollows :exec
DELETE FROM
  follows
WHERE
  to_account_id = $1
  AND from_account_id = $2;