-- name: CreateAccounts :one
INSERT INTO
    accounts (
        account_id,
        username,
        icon,
        explanatory_text,
        locate_id,
        rate,
        user_id,
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
        $9
    ) RETURNING *;

-- name: GetAccountsByID :one
SELECT
    account_id,
    username,
    icon,
    explanatory_text,
    locate_id,
    rate,
    user_id,
    show_locate,
    show_rate,
    create_at,
    update_at
FROM
    accounts
WHERE
    account_id = $1 AND is_delete = false;

-- name: GetAccountsByEmail :one
SELECT
    account_id,
    username,
    icon,
    explanatory_text,
    locate_id,
    rate,
    user_id,
    show_locate,
    show_rate,
    create_at,
    update_at
FROM
    accounts
WHERE
    user_id = (
        SELECT user_id FROM users WHERE email = $1
    ) AND is_delete = false;

-- name: ListAccounts :many
SELECT
    account_id,
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
    username LIKE $1 AND is_delete = false
LIMIT
    $2 OFFSET $3;

-- name: DeleteAccounts :one
UPDATE
    accounts
SET
    is_delete = true
WHERE
    account_id = $1 RETURNING *;

-- name: UpdateAccounts :one
UPDATE
    accounts
SET
    username = $2,
    icon = $3,
    explanatory_text = $4,
    locate_id = $5,
    rate = $6,
    show_locate = $7,
    show_rate = $8
WHERE
    account_id = $1 RETURNING *;


-- name: UpdateRateByID :one
UPDATE
    accounts
SET    
    rate = $2
WHERE
    account_id = $1 RETURNING *;