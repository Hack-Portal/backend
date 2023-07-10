-- name: CreatePastWorkFramework :one
INSERT INTO past_work_frameworks (
    opus,
    framework_id
)VALUES(
    $1,$2
)RETURNING *;


-- name: ListPastWorkFrameworks :many
SELECT 
    frameworks.framework_id,
    frameworks.tech_tag_id,
    frameworks.framework
FROM 
    past_work_frameworks
LEFT OUTER JOIN 
    frameworks 
ON 
    past_work_frameworks.framework_id = frameworks.framework_id 
WHERE 
    past_work_frameworks.opus = $1;
