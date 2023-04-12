-- name: CreateEntery :one
INSERT INTO enteries (
  account_id,
  amount
) VALUES (
  $1, $2
) RETURNING *;

-- name: GetEntery :one
SELECT * FROM enteries
WHERE id = $1 LIMIT 1;

-- name: ListEnteries :many
SELECT * FROM enteries
ORDER BY id
LIMIT $1
OFFSET $2;

-- name: UpdateEntery :exec
UPDATE enteries
SET amount = $2
WHERE id = $1
RETURNING *;

-- name: DeleteEntery :exec
DELETE FROM enteries
WHERE id = $1;