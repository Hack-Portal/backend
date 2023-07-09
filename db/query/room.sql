-- name: CreateRoom :one
INSERT INTO rooms (
    room_id,
    hackathon_id,
    title,
    description,
    member_limit,
    is_status
)VALUES(
    $1,$2,$3,$4,$5,$6
)RETURNING *;

-- name: GetRoom :one
SELECT * FROM rooms WHERE room_id = $1 ;

-- name: ListRoom :many

SELECT 
    * 
FROM 
    rooms 
WHERE 
    member_limit > (
        SELECT count(*) 
        FROM rooms_accounts 
        WHERE rooms_accounts.room_id = rooms.room_id
        ) 
    AND
    is_status = TRUE 
LIMIT $1;

