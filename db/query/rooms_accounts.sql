-- name: CreateRoomsAccounts :one
INSERT INTO rooms_accounts (
    user_id,
    room_id,
    is_owner
)VALUES(
    $1,$2,$3
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
    rooms_accounts.user_id = accounts.user_id 
WHERE 
    rooms_accounts.room_id = $1 ;
-- name: RemoveAccountInRoom :one
DELETE FROM rooms_accounts WHERE room_id = $1 AND user_id = $2 RETURNING *;
