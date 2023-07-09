-- name: CreatePastWorkTag :one
INSERT INTO past_work_tags (
    opus,
    tech_tag_id
)VALUES(
    $1,$2
)RETURNING *;

-- name: GetPastWorkTags :many
SELECT 
    tech_tags.tech_tag_id,
    tech_tags.tech_tag
FROM 
    past_work_tags
LEFT OUTER JOIN 
    tech_tags 
ON 
    past_work_tags.tech_tag_id = tech_tags.tech_tag_id 
WHERE 
    past_work_tags.opus = $1 ;

