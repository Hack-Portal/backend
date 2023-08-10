-- name: CreateUsers :one
INSERT INTO users (
    email,
    hashed_password
)VALUES(
    $1,$2
)RETURNING *;

-- name: GetUsersByID :one
SELECT * FROM users WHERE user_id = $1;

-- name: GetUsersByEmail :one
SELECT * FROM users WHERE email = $1;

-- name: UpdateUsersByID :one
UPDATE users
SET email = $1,
    hashed_password = $2,
    update_at = $3
WHERE user_id = $4
RETURNING *;

-- name: DeleteUsersByID :exec
UPDATE users SET is_delete = $1 WHERE user_id = $2;