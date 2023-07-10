--name ListStatusTags :many
SELECT *
FROM status_tags
where hackathon_id = $1;

--name CreateHackathonStatusTag :one


--name GetHackathonStatusTag :one
