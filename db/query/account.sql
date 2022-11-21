-- name: CreateAccount :one
INSERT INTO account (
    person_id,
    amount
) VALUES (
    $1, $2
) RETURNING *;

-- name: GetAccount :one
SELECT * FROM account
WHERE id = $1 LIMIT 1;

-- name: ListAccount :many
SELECT * FROM account
WHERE person_id = $1
ORDER BY id
LIMIT $2
OFFSET $3;

-- name: UpdateAccount :one
UPDATE account 
SET  amount = $2
WHERE id = $1
RETURNING *;

-- name: DeleteAccount :exec
DELETE FROM account
WHERE id = $1;