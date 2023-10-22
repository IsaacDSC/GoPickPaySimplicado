-- name: GetTransactionByUserID :many
SELECT "transaction".*, "user".type_user 
FROM "transaction" join "user" on "transaction".user_id = "user".id
WHERE "user".id = $1;

-- name: CreateTransaction :exec
INSERT INTO
  "transaction" (id, user_id, value, operation, status)
VALUES
($1, $2, $3, $4, $5);

-- name: UpdateStatusTransaction :exec
UPDATE "transaction" SET status = $1 WHERE id = $2;