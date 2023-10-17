package tests

import (
	"context"
	"fmt"
	"path/filepath"
	"testing"

	"github.com/IsaacDSC/GoPickPaySimplicado/external/configs/env"
	"github.com/IsaacDSC/GoPickPaySimplicado/internal/infra/repositories"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func init() {
	path, _ := filepath.Abs("../../../../.env")
	env.StartEnv(path)
}

func TestTransactionRepository(t *testing.T) {
	repo := repositories.NewTransactionRepositories()
	userID, err := uuid.Parse("fb3399b7-cc9c-4265-89fc-57a299995849")
	assert.NoError(t, err)
	transactions, err := repo.GetTransactionsByUserID(context.Background(), userID)
	assert.NoError(t, err)
	fmt.Printf("\n%+v\n", transactions)
}
