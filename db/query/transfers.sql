-- name: CreateTranfer :one
INSERT INTO transfers (
  from_account_id,
  to_account_id, 
  amount
) VALUES (
  $1, $2 , $3
)
RETURNING *;


-- name: GetTranfer :one
SELECT * FROM transfers
WHERE from_account_id = $1 LIMIT 1;

-- name: ListTranfers :many
SELECT * FROM transfers
ORDER BY id
LIMIT $1
OFFSET $2;
