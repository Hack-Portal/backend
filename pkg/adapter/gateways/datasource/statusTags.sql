-- name: GetStatusTagsByHackathonID :many
SELECT status_tags.status_id ,status_tags.status
FROM status_tags
LEFT OUTER JOIN hackathon_status_tags
ON status_tags.status_id = hackathon_status_tags.status_id
where hackathon_id = $1;

-- name: ListStatusTags :many
SELECT *
FROM status_tags;

-- name: GetStatusTagByStatusID :one
SELECT * FROM status_tags WHERE status_id = $1;

-- name: DeleteStatusTagByStatusID :exec
DELETE FROM
    status_tags
WHERE status_id = $1;

-- name: UpdateStatusTagByStatusID :one
UPDATE
    status_tags
SET
    status = $1
WHERE status_id = $1 RETURNING *;