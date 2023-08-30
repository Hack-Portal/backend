-- name: GetRolesByID :one
SELECT * FROM roles WHERE role_id = $1;

-- name: CreateRoles :one
INSERT INTO roles (
    role
)VALUES($1)
RETURNING *;

-- name: ListRoles :many
SELECT * FROM roles;