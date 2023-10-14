
-- name: FindAllUsers :many
SELECT * FROM "user";

-- name: GetUserByID :one
SELECT * FROM "user" WHERE id = $1;
