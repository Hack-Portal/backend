-- name: GetListStatusTags :one
SELECT *
FROM status_tags
where status_id = $1;

-- name: ListStatusTags :many
SELECT *
FROM status_tags;
