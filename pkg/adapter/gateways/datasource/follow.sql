-- name: CreateFollows :one
INSERT INTO
  follows (to_user_id, from_user_id)
VALUES
($1, $2) RETURNING *;

-- name: ListFollowsByToUserID :many
SELECT
  *
FROM
  follows
WHERE
  to_user_id = $1;

-- name: DeleteFollows :exec
DELETE FROM
  follows
WHERE
  to_user_id = $1
  AND from_user_id = $2;