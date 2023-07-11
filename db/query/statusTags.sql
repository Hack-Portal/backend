-- name: GetStatusTags :one
SELECT status_tags.status_id ,status_tags.status
FROM status_tags
LEFT OUTER JOIN hackathon_status_tags
ON status_tags.status_id = hackathon_status_tags.status_id
where hackathon_id = $1;

-- name: ListStatusTags :many
SELECT *
FROM status_tags;