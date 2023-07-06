-- name: CreateUsers :one
INSERT INTO users (
    hashed_password,
    email
)VALUES(
    $1,$2
)RETURNING *;