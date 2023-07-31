-- name: CreatePastWorks :one
INSERT INTO past_works (
    opus,
    name,
    thumbnail_image,
    explanatory_text
  )
VALUES ($1, $2, $3, $4)
RETURNING *;
-- name: GetPastWorksByOpus :one
SELECT *
FROM past_works
WHERE opus = $1;
-- name: ListPastWorks :many
SELECT opus,
  name,
  explanatory_text
FROM past_works
LIMIT $1 OFFSET $2;