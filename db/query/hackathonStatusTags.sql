--name CreateHackathonStatusTag :one
INSERT INTO hackathon_status_tags (
    hackathon_id,
    status_id
  )VALUES(
    $1,
    $2
  )
RETURNING *;

--name GetStatusTag :many
SELECT *
FROM hackathon_status_tags
WHERE status_id = $1;