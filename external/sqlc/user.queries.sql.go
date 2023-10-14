// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.16.0
// source: user.queries.sql

package sqlc

import (
	"context"

	"github.com/google/uuid"
)

const findAllUsers = `-- name: FindAllUsers :many
SELECT id, complete_name, cpf_cnpj, type_user, email, password, created_at, updated_at FROM "user"
`

func (q *Queries) FindAllUsers(ctx context.Context) ([]User, error) {
	rows, err := q.db.QueryContext(ctx, findAllUsers)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []User
	for rows.Next() {
		var i User
		if err := rows.Scan(
			&i.ID,
			&i.CompleteName,
			&i.CpfCnpj,
			&i.TypeUser,
			&i.Email,
			&i.Password,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getUserByID = `-- name: GetUserByID :one
SELECT id, complete_name, cpf_cnpj, type_user, email, password, created_at, updated_at FROM "user" WHERE id = $1
`

func (q *Queries) GetUserByID(ctx context.Context, id uuid.UUID) (User, error) {
	row := q.db.QueryRowContext(ctx, getUserByID, id)
	var i User
	err := row.Scan(
		&i.ID,
		&i.CompleteName,
		&i.CpfCnpj,
		&i.TypeUser,
		&i.Email,
		&i.Password,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}