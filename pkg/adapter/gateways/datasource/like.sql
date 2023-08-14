-- name: CreateLikes :one
INSERT INTO
    likes(opus, account_id)
VALUES
    ($1, $2) RETURNING *;

-- name: ListLikesByID :many
SELECT
    *
FROM
    likes
WHERE
    account_id = $1 AND is_delete = false;

-- name: DeleteLikesByID :one
UPDATE
    likes
SET
    is_delete = true
WHERE
    account_id = $1
    AND opus = $2 RETURNING *;