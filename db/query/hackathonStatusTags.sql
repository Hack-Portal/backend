-- name: CreateHackathonStatusTag :one
INSERT INTO hackathon_status_tags (
    hackathon_id,
    status_id
  )VALUES(
    $1,
    $2
  )
RETURNING *;

-- name: GetHackathonStatusTags :many
SELECT *
FROM hackathon_status_tags
WHERE hackathon_id = $1;