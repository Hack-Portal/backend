-- name: CreateBookmarks :one
INSERT INTO
    bookmarks(opus, account_id)
VALUES
    ($1, $2) RETURNING *;

-- name: ListBookmarksByID :many
SELECT
    *
FROM
    bookmarks
WHERE
    account_id = $1 AND is_delete = false;

-- name: DeleteBookmarksByID :one
UPDATE
    bookmarks
SET
    is_delete = true
WHERE
    account_id = $1
    AND opus = $2 RETURNING *;