-- name: CreateAccounts :one
INSERT INTO
    accounts (
        account_id,
        email,
        username,
        icon,
        explanatory_text,
        locate_id,
        rate,
        character,
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

-- name: GetAccountsByID :one
SELECT
    *
FROM
    accounts
WHERE
    account_id = $1 AND is_delete = false;

-- name: GetAccountsByEmail :one
SELECT
    *
FROM
    accounts
WHERE
    email = $1 AND is_delete = false;

-- name: ListAccounts :many
SELECT
    *
FROM
    accounts
WHERE
    is_delete = false
ORDER BY
    rate DESC
LIMIT
    $1 OFFSET $2;

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
    character = $7,
    show_locate = $8,
    show_rate = $9,
    update_at = $10,
    twitter_link = $11,
    github_link = $12,
    discord_link = $13
WHERE
    account_id = $1 RETURNING *;


-- name: UpdateRateByID :one
UPDATE
    accounts
SET    
    rate = $2,
    update_at = $3
WHERE
    account_id = $1 RETURNING *;

-- name: CheckAccount :one
SELECT
    count(*)
FROM
    accounts
WHERE 
    account_id = $1 AND email = $2;