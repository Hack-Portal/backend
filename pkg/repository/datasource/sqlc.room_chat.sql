-- name: CreateChat :one
INSERT INTO room_chat (
    chat_id,
    room_id,
    account_id,
    message
)VALUES(
    $1,$2,$3,$4
)RETURNING *;
-- name: ListChat :many
SELECT * FROM room_chat WHERE room_id = $1 ORDER BY created_at DESC LIMIT $2 OFFSET $3;