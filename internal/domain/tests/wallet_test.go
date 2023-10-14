package tests

import (
	"testing"

	"github.com/IsaacDSC/GoPickPaySimplicado/internal/domain"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestWalletDomain(t *testing.T) {
	t.Run("Given multiples transactions the user with operations the credit and debit and status complete", func(t *testing.T) {
		t.Run("When calculate transaction mount wallet", func(t *testing.T) {
			t.Run("Then total amount in wallet is 20000", func(t *testing.T) {
				transactions := []domain.Transactions{
					{
						ID:        uuid.New(),
						Status:    "COMPLETE",
						Value:     "10000",
						Operation: "CREDIT",
					},
					{
						ID:        uuid.New(),
						Status:    "COMPLETE",
						Value:     "10000",
						Operation: "CREDIT",
					},
					{
						ID:        uuid.New(),
						Status:    "COMPLETE",
						Value:     "10000",
						Operation: "CREDIT",
					},
					{
						ID:        uuid.New(),
						Status:    "COMPLETE",
						Value:     "10000",
						Operation: "DEBIT",
					},
				}
				entity := domain.NewWalletEntity()
				entity.WalletFactoryEntity(transactions)
				err := entity.CalculateTotalBalancer()
				assert.NoError(t, err)
				assert.Equal(t, "20000", entity.GetTotalBalancer())
			})

			t.Run("Then total amount in wallet is 10000", func(t *testing.T) {
				transactions := []domain.Transactions{
					{
						ID:        uuid.New(),
						Status:    "COMPLETE",
						Value:     "10000",
						Operation: "CREDIT",
					},
					{
						ID:        uuid.New(),
						Status:    "COMPLETE",
						Value:     "10000",
						Operation: "CREDIT",
					},
					{
						ID:        uuid.New(),
						Status:    "COMPLETE",
						Value:     "10000",
						Operation: "DEBIT",
					},
				}
				entity := domain.NewWalletEntity()
				entity.WalletFactoryEntity(transactions)
				err := entity.CalculateTotalBalancer()
				assert.NoError(t, err)
				assert.Equal(t, "10000", entity.GetTotalBalancer())
			})

			t.Run("Then total amount in wallet is -30000", func(t *testing.T) {
				transactions := []domain.Transactions{
					{
						ID:        uuid.New(),
						Status:    "COMPLETE",
						Value:     "10000",
						Operation: "DEBIT",
					},
					{
						ID:        uuid.New(),
						Status:    "COMPLETE",
						Value:     "10000",
						Operation: "DEBIT",
					},
					{
						ID:        uuid.New(),
						Status:    "COMPLETE",
						Value:     "10000",
						Operation: "DEBIT",
					},
				}
				entity := domain.NewWalletEntity()
				entity.WalletFactoryEntity(transactions)
				err := entity.CalculateTotalBalancer()
				assert.NoError(t, err)
				assert.Equal(t, "-30000", entity.GetTotalBalancer())
			})

		})
	})

	t.Run("Given multiples transactions the user with operations the credit and debit and status different complete", func(t *testing.T) {
		t.Run("When calculate transaction mount wallet", func(t *testing.T) {
			t.Run("Then return error Transaction-In-Process-Retry-Again-Latter", func(t *testing.T) {
				transactions := []domain.Transactions{
					{
						ID:        uuid.New(),
						Status:    "COMPLETE",
						Value:     "10000",
						Operation: "CREDIT",
					},
					{
						ID:        uuid.New(),
						Status:    "COMPLETE",
						Value:     "10000",
						Operation: "CREDIT",
					},
					{
						ID:        uuid.New(),
						Status:    "AWAIT_PERMISSION",
						Value:     "10000",
						Operation: "CREDIT",
					},
					{
						ID:        uuid.New(),
						Status:    "COMPLETE",
						Value:     "10000",
						Operation: "DEBIT",
					},
				}
				entity := domain.NewWalletEntity()
				entity.WalletFactoryEntity(transactions)
				err := entity.CalculateTotalBalancer()
				assert.Error(t, err)
				assert.EqualError(t, err, "Transaction-In-Process-Retry-Again-Latter")
			})
		})
	})

	t.Run("Given not-found transactions the user", func(t *testing.T) {
		t.Run("When calculate total amount in wallet", func(t *testing.T) {
			t.Run("Then zero amount in wallet", func(t *testing.T) {
				transactions := []domain.Transactions{}
				entity := domain.NewWalletEntity()
				entity.WalletFactoryEntity(transactions)
				err := entity.CalculateTotalBalancer()
				assert.NoError(t, err)
				assert.Equal(t, "0", entity.GetTotalBalancer())
			})
		})
	})

}
