
-- name: FindAllUsers :many
SELECT * FROM "user";

-- name: GetUserByID :one
SELECT * FROM "user" WHERE id = $1;


-- name: TransactionsOnUser :many
select * from "user" join "transaction" on "transaction".user_id = "user".id;
