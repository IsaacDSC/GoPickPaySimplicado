package repositories

import (
	"context"

	"github.com/IsaacDSC/GoPickPaySimplicado/external/configs/database"
	"github.com/IsaacDSC/GoPickPaySimplicado/external/sqlc"
	"github.com/IsaacDSC/GoPickPaySimplicado/internal/domain"
	"github.com/google/uuid"
)

type UserRepository struct{}

func NewUserRepository() *UserRepository {
	return new(UserRepository)
}

func (*UserRepository) GetUserByID(
	ctx context.Context, userID uuid.UUID,
) (output domain.UserEntity, err error) {
	conn := database.DbConn()
	db := sqlc.New(conn)
	user, err := db.GetUserByID(ctx, userID)
	if err != nil {
		return
	}
	output = domain.NewUserEntity().ToDomain(user)
	return
}
