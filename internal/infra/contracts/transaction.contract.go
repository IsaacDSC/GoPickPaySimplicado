package contracts

import (
	"context"

	"github.com/IsaacDSC/GoPickPaySimplicado/internal/domain"
	"github.com/google/uuid"
)

type TransactionRepositoriesInterface interface {
	GetTransactionsByUserID(ctx context.Context, UserID uuid.UUID) (output []domain.Transactions, err error)
	InsertTransaction(ctx context.Context, input domain.TransactionEntity) (err error)
}
