-- name: CreateBookmark :one
INSERT INTO bookmarks(
    hackathon_id,
    user_id
)VALUES(
    $1,$2
)RETURNING *;

-- name: ListBookmark :many
SELECT * FROM bookmarks WHERE user_id = $1 ;