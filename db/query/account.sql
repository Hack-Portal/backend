-- name: CreateAccount :one
INSERT INTO
    accounts (
        user_id,
        username,
        icon,
        explanatory_text,
        locate_id,
        rate,
        hashed_password,
        email,
        show_locate,
        show_rate
    )
VALUES
    (
        $1,
        $2,
        $3,
        $4,
        $5,
        $6,
        $7,
        $8,
        $9,
        $10
    ) RETURNING *;

-- name: GetAccountByID :one
SELECT
    user_id,
    username,
    icon,
    explanatory_text,
    (
        SELECT
            name
        FROM
            locates
        WHERE
            locate_id = accounts.locate_id
    ) as locate,
    rate,
    hashed_password,
    email,
    show_locate,
    show_rate,
    create_at,
    update_at
FROM
    accounts
WHERE
    user_id = $1;

-- name: GetAccountByEmail :one
SELECT
    user_id,
    username,
    icon,
    explanatory_text,
    (
        SELECT
            name
        FROM
            locates
        WHERE
            locate_id = accounts.locate_id
    ) as locate,
    rate,
    hashed_password,
    email,
    show_locate,
    show_rate,
    create_at,
    update_at
FROM
    accounts
WHERE
    email = $1;

-- name: ListAccounts :many
SELECT
    user_id,
    username,
    icon,
    (
        SELECT
            name
        FROM
            locates
        WHERE
            locate_id = accounts.locate_id
    ) as locate,
    rate,
    show_locate,
    show_rate
FROM
    accounts
WHERE
    username LIKE $1
LIMIT
    $2 OFFSET $3;

-- name: SoftDeleteAccount :one
UPDATE
    accounts
SET
    is_delete = true
WHERE
    user_id = $1 RETURNING *;