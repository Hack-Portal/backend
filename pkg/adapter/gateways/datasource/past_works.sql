-- name: CreatePastWorks :one
INSERT INTO past_works (
    name,
    thumbnail_image,
    explanatory_text,
    award_data_id
  )
VALUES ($1, $2, $3,$4)
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

-- name: UpdatePastWorksByID :one
UPDATE past_works
SET name = $1,
    thumbnail_image = $2,
    explanatory_text = $3,
    award_data_id = $4,
    update_at = $5
WHERE opus = $6
RETURNING *;
