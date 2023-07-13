-- name: CreateFollow :one
INSERT INTO follows (
    to_user_id,
    from_user_id
  )
VALUES($1, $2)
RETURNING *;
-- name: ListFollowByToUserID :many
SELECT *
FROM follows
WHERE to_user_id = $1;

-- name: RemoveFollow :exec
DELETE FROM follows
WHERE to_user_id = $1
  AND from_user_id = $2;