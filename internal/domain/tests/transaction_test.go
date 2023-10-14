package tests

import (
	"testing"

	"github.com/IsaacDSC/GoPickPaySimplicado/internal/domain"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestTransactionDomain(t *testing.T) {
	t.Run("Given wallet bigger than value transfer", func(t *testing.T) {
		t.Run("When execute at transaction with status AWAIT_PERMISSION", func(t *testing.T) {
			t.Run("Then not Return Error", func(t *testing.T) {
				entity := domain.NewTransactionEntity()
				wallet := domain.NewWalletEntity()
				transactions := []domain.Transactions{
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
						Operation: "CREDIT",
					},
					{
						ID:        uuid.New(),
						Status:    "COMPLETE",
						Value:     "10000",
						Operation: "DEBIT",
					},
				}
				wallet.WalletFactoryEntity(transactions)
				entity.TransactionFactory(
					"CONSUMER", "10000", "DEBIT", *wallet,
				)
				listErrors := entity.Transaction()
				assert.Len(t, listErrors, 1)
				assert.Equal(t, "Transaction-In-Process-Retry-Again-Latter", listErrors[0])
				assert.Equal(t, "NOT_AUTHORIZED", entity.Status)
			})
		})
		t.Run("When execute at transaction with status COMPLETE", func(t *testing.T) {
			t.Run("Then not Return Error", func(t *testing.T) {
				entity := domain.NewTransactionEntity()
				wallet := domain.NewWalletEntity()
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
				wallet.WalletFactoryEntity(transactions)
				entity.TransactionFactory(
					"CONSUMER", "10000", "DEBIT", *wallet,
				)
				listErrors := entity.Transaction()
				assert.Len(t, listErrors, 0)
				assert.Equal(t, "AWAIT_PERMISSION", entity.Status)
			})
		})
	})

	t.Run("Given wallet smaller than value transfer", func(t *testing.T) {
		t.Run("When execute at transaction", func(t *testing.T) {
			t.Run("Then not Return Error", func(t *testing.T) {
				entity := domain.NewTransactionEntity()
				wallet := domain.NewWalletEntity()
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
				wallet.WalletFactoryEntity(transactions)
				entity.TransactionFactory(
					"CONSUMER", "20000", "DEBIT", *wallet,
				)
				listErrors := entity.Transaction()
				assert.Len(t, listErrors, 1)
				assert.Equal(t, listErrors[0], "UNAUTHORIZED-TRANSFER-WALLET")
				assert.Equal(t, "NOT-AUTHORIZED", entity.Status)
			})
		})
	})

}
