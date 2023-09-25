-- name: CreatePastWorks :one
INSERT INTO past_works (
    name,
    thumbnail_image,
    explanatory_text,
    award_data_id
  )
VALUES ($1, $2, $3, $4)
RETURNING *;
-- name: GetPastWorksByOpus :one
SELECT opus, name, thumbnail_image, explanatory_text, award_data_id, create_at, update_at, is_delete
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
-- name: DeletePastWorksByID :one
UPDATE past_works
SET is_delete = $1
WHERE opus = $2
RETURNING *;