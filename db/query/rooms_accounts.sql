-- name: CreateRoomsAccounts :one
INSERT INTO rooms_accounts (
    user_id,
    room_id
)VALUES(
    $1,$2
)RETURNING *;

-- name: GetRoomsAccounts :many
SELECT 
    accounts.user_id,
    accounts.username,  
    accounts.icon
FROM 
    rooms_accounts
LEFT OUTER JOIN 
    accounts 
ON 
    rooms_accounts.tech_tag_id = tech_tag_id 
WHERE 
    rooms_accounts.room_id = $1 ;