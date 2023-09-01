-- name: CreateRooms :one
INSERT INTO rooms (
        room_id,
        hackathon_id,
        title,
        description,
        member_limit,
        include_rate,
        is_closing
    )
VALUES($1, $2, $3, $4, $5, $6, $7)
RETURNING *;
-- name: GetRoomsByID :one
SELECT *
FROM rooms
WHERE room_id = $1 AND is_delete = false;


-- name: ListRooms :many
SELECT *
FROM rooms
WHERE is_delete = false
LIMIT $1 OFFSET $2;

-- name: DeleteRoomsByID :one
UPDATE rooms
SET is_delete = true
WHERE room_id = $1
RETURNING *;

-- name: UpdateRoomsByID :one
UPDATE rooms
SET hackathon_id = $1,
    title = $2,
    description = $3,
    member_limit = $4,
    update_at = $5,
    is_closing = $6
WHERE room_id = $7
RETURNING *;