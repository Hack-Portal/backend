-- name: CreateRoomsFramework :one
INSERT INTO rooms_frameworks (
    room_id,
    framework_id
)VALUES(
    $1,$2
)RETURNING *;


-- name: ListRoomsFrameworks :many
SELECT 
    frameworks.framework_id,
    frameworks.tech_tag_id,
    frameworks.framework
FROM 
    rooms_frameworks
LEFT OUTER JOIN 
    frameworks 
ON 
    rooms_frameworks.framework_id = frameworks.framework_id 
WHERE 
    rooms_frameworks.room_id = $1;
