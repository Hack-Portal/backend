-- name: CreateHackathonStatusTags :one
INSERT INTO hackathon_status_tags (
    hackathon_id,
    status_id
  )VALUES(
    $1,
    $2
  )
RETURNING *;

-- name: ListHackathonStatusTagsByID :many
SELECT *
FROM hackathon_status_tags
WHERE hackathon_id = $1;

-- name: DeleteHackathonStatusTagsByID :exec
DELETE FROM hackathon_status_tags WHERE hackathon_id = $1;