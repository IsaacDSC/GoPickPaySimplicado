package tests

import (
	"context"
	"fmt"
	"path/filepath"
	"testing"

	"github.com/IsaacDSC/GoPickPaySimplicado/external/configs/env"
	"github.com/IsaacDSC/GoPickPaySimplicado/internal/domain"
	"github.com/IsaacDSC/GoPickPaySimplicado/internal/infra/repositories"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func init() {
	path, _ := filepath.Abs("../../../../.env")
	env.StartEnv(path)
}

func TestTransactionRepository(t *testing.T) {
	t.Skip()
	repo := repositories.NewTransactionRepositories()
	userID, err := uuid.Parse("fb3399b7-cc9c-4265-89fc-57a299995849")
	assert.NoError(t, err)
	transactions, err := repo.GetTransactionsByUserID(context.Background(), userID)
	assert.NoError(t, err)
	fmt.Printf("\n%+v\n", transactions)
}

func TestTransactionUOWRepository(t *testing.T) {
	ID := uuid.New()
	repo := repositories.NewTransactionRepositories()
	userID, err := uuid.Parse("fb3399b7-cc9c-4265-89fc-57a299995849")
	assert.NoError(t, err)
	transactions, err := repo.GetTransactionsByUserID(context.Background(), userID)
	assert.NoError(t, err)
	fmt.Printf("\n%+v\n", transactions)
	err = repo.InsertTransaction(context.Background(), domain.TransactionEntity{
		ID:        ID,
		UserID:    userID,
		Value:     "9000000",
		Operation: "CREDIT",
		Status:    "COMPLETE",
	})
	assert.NoError(t, err)
	repo.Done()
}
