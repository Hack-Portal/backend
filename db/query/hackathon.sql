-- name: CreateHackathon :one
INSERT INTO hackathons (
    name,
    icon,
    description,
    link,
    expired,
    start_date,
    term
  )
VALUES(
    $1,$2,$3,$4,$5,$6,$7
  )
RETURNING *;
-- name: ListHackathons :many
SELECT *
FROM hackathons
WHERE expired > $1
ORDER BY hackathon_id
LIMIT $2 OFFSET $3;

-- name: GetHackathon :one
SELECT *
FROM hackathons
WHERE hackathon_id = $1;