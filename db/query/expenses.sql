-- name: CreateExpenses :one
INSERT INTO expenses (
    account_id,
    value
) VALUES (
    $1, $2
) RETURNING *;

--name: GetExpense :one
SELECT * FROM expenses
WHERE id = $1 LIMIT 1;

-- name: ListExpenses :many
SELECT * FROM expenses
WHERE account_id = $1
ORDER BY id
LIMIT $2
OFFSET $3;

-- name: UpdateExpenses :one
UPDATE expenses 
SET  value = $2
WHERE id = $1
RETURNING *;

-- name: DeleteExpense :exec
DELETE FROM expenses
WHERE id = $1;