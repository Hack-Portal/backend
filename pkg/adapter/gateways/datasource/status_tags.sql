-- name: CreateStatusTags :one
INSERT INTO status_tags (status) VALUES ($1) RETURNING *;

-- name: GetStatusTagsByTag :one
SELECT status_id ,status
FROM status_tags
where status_id = $1;

-- name: GetStatusTagsByHackathonID :one
SELECT status_tags.status_id ,status_tags.status
FROM status_tags
LEFT OUTER JOIN hackathon_status_tags
ON status_tags.status_id = hackathon_status_tags.status_id
where hackathon_id = $1;


-- name: ListStatusTags :many
SELECT *
FROM status_tags;

-- name: DeleteStatusTagsByStatusID :exec
DELETE FROM
    status_tags
WHERE status_id = $1;
