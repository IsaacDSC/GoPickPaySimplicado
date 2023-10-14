package contracts

import (
	"context"

	"github.com/IsaacDSC/GoPickPaySimplicado/internal/domain"
	"github.com/gofrs/uuid"
)

type TransactionRepositoriesInterface interface {
	GetTransactionsByUserID(ctx context.Context, UserID uuid.UUID) (output []domain.TransactionEntity, err error)
	InsertTransaction(ctx context.Context, input domain.TransactionEntity) (err error)
}
