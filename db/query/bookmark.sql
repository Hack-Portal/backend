-- name: CreateBookmark :one
INSERT INTO
    bookmarks(hackathon_id, user_id)
VALUES
    ($1, $2) RETURNING *;

-- name: ListBookmarkByUserID :many
SELECT
    *
FROM
    bookmarks
WHERE
    user_id = $1;

-- name: SoftRemoveBookmark :one
UPDATE
    bookmarks
SET
    is_delete = true
WHERE
    user_id = $1
    AND hackathon_id = $2 RETURNING *;