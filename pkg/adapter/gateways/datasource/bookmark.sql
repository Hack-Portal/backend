-- name: CreateBookmarks :one
INSERT INTO
    bookmarks(hackathon_id, user_id)
VALUES
    ($1, $2) RETURNING *;

-- name: ListBookmarksByID :many
SELECT
    *
FROM
    bookmarks
WHERE
    user_id = $1 AND is_delete = false;

-- name: DeleteBookmarksByID :one
UPDATE
    bookmarks
SET
    is_delete = true
WHERE
    user_id = $1
    AND hackathon_id = $2 RETURNING *;