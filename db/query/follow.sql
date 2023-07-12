-- name CreateFollow :one
INSERT INTO follows (
    to_user_id,
    form_user_id,
    create_at
  )
VALUES($1, $2, $3)
RETURNING *;
-- name ListFollow :many
SELECT *
FROM follows
WHERE to_user_id = $1;
-- name RemoveFollow :exec
DELETE FROM follows
WHERE to_user_id = $1
  AND form_user_id = $2;