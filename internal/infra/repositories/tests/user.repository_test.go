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

func TestUserRepository(t *testing.T) {
	repo := repositories.NewUserRepository()
	userID, err := uuid.Parse("e9d94c21-4471-4189-9e83-cc6697089740")
	assert.NoError(t, err)
	user, err := repo.GetUserByID(context.Background(), userID)
	assert.NoError(t, err)
	fmt.Printf("\n%+v\n", user)
}
