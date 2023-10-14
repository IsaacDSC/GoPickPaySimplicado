-- name: GetTransactionByUserID :many
SELECT
  *
FROM
  "transaction"
WHERE
  user_id = $1;

-- name: CreateTransaction :exec
INSERT INTO
  "transaction" (id, user_id, value, operation, status)
VALUES
($1, $2, $3, $4, $5);