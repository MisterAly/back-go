-- name: CreatePerson :one
INSERT INTO person (
    name,
    document,
    phone
) VALUES (
    $1, $2, $3
) RETURNING *;

-- name: GetPerson :one
SELECT * FROM person
WHERE ID = $1 LIMIT 1;

-- name: ListPerson :many   
SELECT * FROM person
ORDER BY id
LIMIT $1
OFFSET $2;

-- name: UpdatePerson :one
UPDATE person 
SET 
    document = $2,
    phone = $3
WHERE 
    name = $1
RETURNING *;

-- name: DeletePerson :exec
DELETE FROM person
WHERE id = $1;