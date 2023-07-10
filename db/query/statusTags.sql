--name ListStatusTags :many
SELECT *
FROM status_tags
ORDER BY status_tag_id;

--name CreateHackathonStatusTag :one


--name GetHackathonStatusTag :one
