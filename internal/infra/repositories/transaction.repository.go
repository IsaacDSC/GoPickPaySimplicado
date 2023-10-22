package repositories

import (
	"context"
	"database/sql"

	"github.com/IsaacDSC/GoPickPaySimplicado/external/configs/database"
	"github.com/IsaacDSC/GoPickPaySimplicado/external/sqlc"
	"github.com/IsaacDSC/GoPickPaySimplicado/internal/domain"
	"github.com/google/uuid"
)

type TransactionRepositories struct {
	db *sqlc.Queries
	tx *sql.Tx
}

func NewTransactionRepositories() *TransactionRepositories {
	t := new(TransactionRepositories)
	tx, _ := database.DbConn().Begin()
	t.tx = tx
	t.db = sqlc.New(tx)
	return t
}

func (tr *TransactionRepositories) GetTransactionsByUserID(
	ctx context.Context, UserID uuid.UUID,
) (output []domain.Transactions, err error) {
	transactions, err := tr.db.GetTransactionByUserID(ctx, UserID)
	if err != nil {
		return
	}
	for index := range transactions {
		out := domain.Transactions{
			ID:        transactions[index].ID,
			TypeUser:  transactions[index].TypeUser,
			Status:    transactions[index].Status,
			Value:     transactions[index].Value,
			Operation: transactions[index].Operation.String,
		}
		output = append(output, out)
	}
	return
}

func (tr *TransactionRepositories) InsertTransaction(
	ctx context.Context, input domain.TransactionEntity,
) (err error) {
	err = tr.db.CreateTransaction(ctx, sqlc.CreateTransactionParams{
		ID:        input.ID,
		UserID:    input.UserID,
		Value:     input.Value,
		Operation: sql.NullString{String: input.Operation, Valid: true},
		Status:    input.Status,
	})
	return
}

func (tr *TransactionRepositories) UpdateStatusTransaction(
	ctx context.Context, transactionID uuid.UUID, status string,
) (err error) {
	err = tr.db.UpdateStatusTransaction(ctx, sqlc.UpdateStatusTransactionParams{
		Status: status,
		ID:     transactionID,
	})
	return
}

func (tr *TransactionRepositories) Done() {
	tr.tx.Commit()
}

func (tr *TransactionRepositories) Rollback() {
	tr.tx.Rollback()
}
