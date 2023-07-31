-- name: CreateAccountTags :one
INSERT INTO
    account_tags (user_id, tech_tag_id)
VALUES
($1, $2) RETURNING *;

-- name: ListAccountTagsByUserID :many
SELECT
    tech_tags.tech_tag_id,
    tech_tags.language
FROM
    account_tags
    LEFT OUTER JOIN tech_tags ON account_tags.tech_tag_id = tech_tags.tech_tag_id
WHERE
    account_tags.user_id = $1;

-- name: DeleteAccounttagsByUserID :exec
DELETE FROM
    account_tags
WHERE
    user_id = $1;