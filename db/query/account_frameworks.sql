-- name: CreateAccountFramework :one
INSERT INTO account_frameworks (
    user_id,
    framework_id
)VALUES(
    $1,$2
)RETURNING *;


-- name: ListAccountFrameworksByUserID :many
SELECT 
    frameworks.framework_id,
    frameworks.tech_tag_id,
    frameworks.framework
FROM 
    account_frameworks
LEFT OUTER JOIN 
    frameworks 
ON 
    account_frameworks.framework_id = frameworks.framework_id 
WHERE 
    account_frameworks.user_Id = $1;
