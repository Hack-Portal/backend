-- name: CreateRoom :one
INSERT INTO rooms (
    room_id,
    hackathon_id,
    title,
    description,
    member_limit
)VALUES(
    $1,$2,$3,$4,$5
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
    hackathon_id IN (
        SELECT hackathon_id
        FROM hackathons
        WHERE expired > $1
    ) ;

