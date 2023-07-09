-- name: GetAccountAuth :one
SELECT 
    user_id,
    hashed_password,
    email
FROM 
    accounts
WHERE
    user_id = $1;

-- name: GetAccount :one
SELECT 
    user_id,
    username,
    icon,
    explanatory_text,
    locate_id,
    rate,
    show_locate,
    show_rate
FROM
    accounts
WHERE
    user_id = $1;

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
WHERE username LIKE $1
LIMIT $2
OFFSET $3;

-- name: CreateAccount :one
INSERT INTO accounts (
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
)VALUES(
    $1,$2,$3,$4,$5,$6,$7,$8,$9,$10
)RETURNING *;