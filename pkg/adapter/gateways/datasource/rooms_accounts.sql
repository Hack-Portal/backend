-- name: CreateRoomsAccounts :one
INSERT INTO rooms_accounts (
    account_id,
    room_id,
    is_owner,
    role
)VALUES(
    $1,$2,$3,$4
)RETURNING *;

-- name: GetRoomsAccountsByID :many
SELECT 
    accounts.account_id, 
    accounts.icon,
    rooms_accounts.is_owner,
    rooms_accounts.role    
FROM 
    rooms_accounts
LEFT OUTER JOIN 
    accounts 
ON 
    rooms_accounts.account_id = accounts.account_id 
WHERE 
    rooms_accounts.room_id = $1 ;

-- name: DeleteRoomsAccountsByID :exec
DELETE FROM rooms_accounts WHERE room_id = $1 AND account_id = $2;
