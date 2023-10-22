package uow

import (
	"context"
	"database/sql"

	"github.com/IsaacDSC/GoPickPaySimplicado/external/configs/database"
)

type UnitOfWork struct {
	db *sql.Tx
}

func NewUnitOfWork() *UnitOfWork {
	uow := new(UnitOfWork)
	return uow
}

func (uow *UnitOfWork) Get(ctx context.Context) *sql.Tx {
	conn := database.DbConn()
	transaction, err := conn.Begin()
	if err != nil {
		panic(err)
	}
	uow.db = transaction
	return uow.db
}
