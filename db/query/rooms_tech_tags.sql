-- name: CreateRoomsTechTag :one
INSERT INTO rooms_tech_tags (
    room_id,
    tech_tag_id
)VALUES(
    $1,$2
)RETURNING *;

-- name: GetRoomsTechTags :many
SELECT 
    tech_tags.tech_tag_id,
    tech_tags.tech_tag
FROM 
    rooms_tech_tags
LEFT OUTER JOIN 
    tech_tags 
ON 
    rooms_tech_tags.tech_tag_id = tech_tags.tech_tag_id
WHERE 
    rooms_tech_tags.room_id = $1 ;