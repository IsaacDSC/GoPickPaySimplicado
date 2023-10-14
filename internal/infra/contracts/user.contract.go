package contracts

import (
	"context"

	"github.com/IsaacDSC/GoPickPaySimplicado/internal/domain"
	"github.com/gofrs/uuid"
)

type UserRepositoryInterface interface {
	GetUserByID(ctx context.Context, userID uuid.UUID) (output domain.UserEntity, err error)
}
