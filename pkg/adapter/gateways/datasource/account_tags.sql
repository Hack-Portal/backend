-- name: CreateAccountTags :one
INSERT INTO
    account_tags (account_id, tech_tag_id)
VALUES
($1, $2) RETURNING *;

-- name: ListAccountTagsByUserID :many
SELECT
    tech_tags.tech_tag_id,
    tech_tags.language
    tech_tags.icon
FROM
    account_tags
    LEFT OUTER JOIN tech_tags ON account_tags.tech_tag_id = tech_tags.tech_tag_id
WHERE
    account_tags.account_id = $1;

-- name: DeleteAccountTagsByUserID :exec
DELETE FROM
    account_tags
WHERE
    account_id = $1;