--name ListStatusTags :many
SELECT *
FROM status_tags
where hackathon_id = $1;
