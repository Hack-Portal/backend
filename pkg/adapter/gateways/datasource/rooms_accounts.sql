-- name: CreateRoomsAccounts :one
INSERT INTO rooms_accounts (
    user_id,
    room_id,
    is_owner
)VALUES(
    $1,$2,$3
)RETURNING *;

-- name: GetRoomsAccountsByID :many
SELECT 
    accounts.user_id, 
    accounts.icon,
    rooms_accounts.is_owner
FROM 
    rooms_accounts
LEFT OUTER JOIN 
    accounts 
ON 
    rooms_accounts.user_id = accounts.user_id 
WHERE 
    rooms_accounts.room_id = $1 ;

-- name: DeleteRoomsAccountsByID :exec
DELETE FROM rooms_accounts WHERE room_id = $1 AND user_id = $2;
