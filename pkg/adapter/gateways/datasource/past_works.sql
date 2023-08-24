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
SELECT *,
  (
    SELECT count(*)
    FROM likes
    WHERE likes.opus = past_works.opus
  ) as like_count,
  CASE
    WHEN $2 IS NOT NULL AND EXISTS (
      SELECT 1
      FROM likes
      WHERE likes.opus = $1 AND likes.account_id = $2
    ) THEN TRUE
    ELSE FALSE
  END as is_liked
FROM past_works
WHERE opus = $1;

-- name: ListPastWorks :many
SELECT opus,
  name,
  explanatory_text,
  (
    SELECT count(*)
    FROM likes
    WHERE likes.opus = past_works.opus
  ) as like_count,
  CASE
    WHEN $2 IS NOT NULL AND EXISTS (
      SELECT 1
      FROM likes
      WHERE likes.opus = past_works.opus AND likes.account_id = $3
    ) THEN TRUE
    ELSE FALSE
  END as is_liked
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