package repositories

import (
	"context"
	"database/sql"

	"github.com/IsaacDSC/GoPickPaySimplicado/external/configs/database"
	"github.com/IsaacDSC/GoPickPaySimplicado/external/sqlc"
	"github.com/IsaacDSC/GoPickPaySimplicado/internal/domain"
	"github.com/google/uuid"
)

type TransactionRepositories struct{}

func (*TransactionRepositories) GetTransactionsByUserID(
	ctx context.Context, UserID uuid.UUID,
) (output []domain.TransactionEntity, err error) {
	conn := database.DbConn()
	db := sqlc.New(conn)
	transactions, err := db.GetTransactionByUserID(ctx, UserID)
	if err != nil {
		return
	}
	for index := range transactions {
		out := domain.NewTransactionEntity().ToDomain(transactions[index])
		output = append(output, out)
	}
	return
}

func (*TransactionRepositories) InsertTransaction(
	ctx context.Context, input domain.TransactionEntity,
) (err error) {
	conn := database.DbConn()
	db := sqlc.New(conn)
	err = db.CreateTransaction(ctx, sqlc.CreateTransactionParams{
		ID:        input.ID,
		UserID:    input.UserID,
		Value:     input.Value,
		Operation: sql.NullString{String: input.Operation, Valid: true},
		Status:    input.Status,
	})
	return
}
