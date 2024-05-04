-- name: GetEntry :one
SELECT * FROM entries
WHERE id = $1 LIMIT 1;

-- name: ListEntries :many
SELECT * FROM entries
ORDER BY id
OFFSET $1
LIMIT $2
;

-- name: CreateEntry :one
INSERT INTO entries (
  amount, account_id
) VALUES (
  $1, $2
)
RETURNING *;
