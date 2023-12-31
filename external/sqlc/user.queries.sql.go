// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.16.0
// source: user.queries.sql

package sqlc

import (
	"context"
	"database/sql"

	"github.com/google/uuid"
)

const findAllUsers = `-- name: FindAllUsers :many
SELECT id, complete_name, cpf_cnpj, type_user, email, password, updated_at, created_at FROM "user"
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
			&i.UpdatedAt,
			&i.CreatedAt,
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
SELECT id, complete_name, cpf_cnpj, type_user, email, password, updated_at, created_at FROM "user" WHERE id = $1
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
		&i.UpdatedAt,
		&i.CreatedAt,
	)
	return i, err
}

const transactionsOnUser = `-- name: TransactionsOnUser :many
select "user".id, complete_name, cpf_cnpj, type_user, email, password, "user".updated_at, "user".created_at, transaction.id, user_id, value, operation, status, transaction.updated_at, transaction.created_at from "user" join "transaction" on "transaction".user_id = "user".id
`

type TransactionsOnUserRow struct {
	ID           uuid.UUID
	CompleteName sql.NullString
	CpfCnpj      string
	TypeUser     string
	Email        string
	Password     sql.NullString
	UpdatedAt    sql.NullTime
	CreatedAt    sql.NullTime
	ID_2         uuid.UUID
	UserID       uuid.UUID
	Value        string
	Operation    sql.NullString
	Status       string
	UpdatedAt_2  sql.NullTime
	CreatedAt_2  sql.NullTime
}

func (q *Queries) TransactionsOnUser(ctx context.Context) ([]TransactionsOnUserRow, error) {
	rows, err := q.db.QueryContext(ctx, transactionsOnUser)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []TransactionsOnUserRow
	for rows.Next() {
		var i TransactionsOnUserRow
		if err := rows.Scan(
			&i.ID,
			&i.CompleteName,
			&i.CpfCnpj,
			&i.TypeUser,
			&i.Email,
			&i.Password,
			&i.UpdatedAt,
			&i.CreatedAt,
			&i.ID_2,
			&i.UserID,
			&i.Value,
			&i.Operation,
			&i.Status,
			&i.UpdatedAt_2,
			&i.CreatedAt_2,
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
