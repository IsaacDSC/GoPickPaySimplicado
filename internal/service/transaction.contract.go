package service

import (
	"context"

	"github.com/IsaacDSC/GoPickPaySimplicado/internal/shared/dto"
)

type TransactionServiceInterface interface {
	Transfer(ctx context.Context, input dto.TransactionDtoInput) (list_errors []error)
}
