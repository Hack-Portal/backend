-- name: GetRoomsAccountsRolesIDByIDs :one
SELECT rooms_account_id
FROM rooms_accounts
WHERE room_id = $1
  AND account_id = $2;

-- name: CreateRoomsAccountsRoles :one
INSERT INTO rooms_accounts_roles (rooms_account_id, role_id)
VALUES ($1, $2)
RETURNING *;

-- name: ListRoomsAccountsRolesByIDs :many
SELECT roles.role_id,
  roles.role
FROM roles
  LEFT OUTER JOIN rooms_accounts_roles ON rooms_accounts_roles.role_id = roles.role_id
WHERE rooms_accounts_roles.rooms_account_id = (
    SELECT rooms_account_id
    FROM rooms_accounts
    WHERE room_id = $1
      AND account_id = $2
  );

-- name: ListRoomsAccountsRolesByID :many
SELECT roles.role_id,
  roles.role
FROM roles
  LEFT OUTER JOIN rooms_accounts_roles ON rooms_accounts_roles.role_id = roles.role_id
WHERE rooms_accounts_roles.rooms_account_id = $1;

-- name: DeleteRoomsAccountsRolesByID :exec
DELETE FROM rooms_accounts_roles
WHERE rooms_account_id = $1;