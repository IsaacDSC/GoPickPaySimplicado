package service

import (
	"context"
	"strings"
	"testing"

	"github.com/IsaacDSC/GoPickPaySimplicado/external/mocks"
	"github.com/IsaacDSC/GoPickPaySimplicado/internal/domain"
	"github.com/IsaacDSC/GoPickPaySimplicado/internal/shared/dto"
	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestTransactionService(t *testing.T) {
	t.Run("Given init transaction payer", func(t *testing.T) {
		t.Run("When send input valid", func(t *testing.T) {
			t.Run("Then transaction not return error", func(t *testing.T) {
				ctl := gomock.NewController(t)
				defer ctl.Finish()
				userRepository := mocks.NewMockUserRepositoryInterface(ctl)
				transactionAuthGateway := mocks.NewMockOperationTransactionGatewayInterface(ctl)
				producer := mocks.NewMockIProducerQueue(ctl)
				userID := uuid.New()
				userRepository.EXPECT().GetUserByID(gomock.Any(), gomock.Any()).Return(
					domain.UserEntity{
						ID:           userID,
						CompleteName: "Advino Silva Name",
						Cpf_cnpj:     "1233213123123123",
						TypeUser:     "CONSUMER",
						Email:        "advino@gamil.com",
					}, nil,
				)
				transactionRepository := mocks.NewMockTransactionRepositoriesInterface(ctl)
				transactionRepository.EXPECT().GetTransactionsByUserID(gomock.Any(), gomock.Any()).Return(
					[]domain.Transactions{
						{
							ID:        uuid.New(),
							TypeUser:  "CONSUMER",
							Status:    "COMPLETE",
							Value:     "50000000",
							Operation: "CREDIT",
						},
						{
							ID:        uuid.New(),
							TypeUser:  "CONSUMER",
							Status:    "COMPLETE",
							Value:     "50000000",
							Operation: "CREDIT",
						},
					}, nil,
				)
				transactionRepository.EXPECT().InsertTransaction(gomock.Any(), gomock.Any()).Return(nil)
				service := NewTransactionService(
					userRepository, transactionRepository,
					transactionAuthGateway,
					producer,
					domain.NewTransactionEntity(domain.NewWalletEntity()),
					domain.NewTransactionEntity(domain.NewWalletEntity()),
				)
				service.input = dto.TransactionDtoInput{
					Value:   "10000000",
					PayerID: userID,
				}
				payerMailer, list_errors := service.executeTransactionPayer(context.Background())
				for index := range list_errors {
					assert.NoError(t, list_errors[index])
					assert.True(t, strings.Contains(payerMailer, "@"))
				}
			})
		})
	})

	t.Run("Given init transaction payer", func(t *testing.T) {
		t.Run("When send input valid", func(t *testing.T) {
			t.Run("Then transaction not return error", func(t *testing.T) {
				ctl := gomock.NewController(t)
				defer ctl.Finish()
				userRepository := mocks.NewMockUserRepositoryInterface(ctl)
				transactionAuthGateway := mocks.NewMockOperationTransactionGatewayInterface(ctl)
				producer := mocks.NewMockIProducerQueue(ctl)
				userID := uuid.New()
				userRepository.EXPECT().GetUserByID(gomock.Any(), gomock.Any()).Return(
					domain.UserEntity{
						ID:           userID,
						CompleteName: "Advino Silva Name",
						Cpf_cnpj:     "1233213123123123",
						TypeUser:     "STOREKEEPER",
						Email:        "advino@gamil.com",
					}, nil,
				)
				transactionRepository := mocks.NewMockTransactionRepositoriesInterface(ctl)
				transactionRepository.EXPECT().GetTransactionsByUserID(gomock.Any(), gomock.Any()).Return(
					[]domain.Transactions{
						{
							ID:        uuid.New(),
							TypeUser:  "CONSUMER",
							Status:    "COMPLETE",
							Value:     "50000000",
							Operation: "CREDIT",
						},
						{
							ID:        uuid.New(),
							TypeUser:  "CONSUMER",
							Status:    "COMPLETE",
							Value:     "50000000",
							Operation: "CREDIT",
						},
					}, nil,
				)
				transactionRepository.EXPECT().InsertTransaction(gomock.Any(), gomock.Any()).Return(nil)
				service := NewTransactionService(
					userRepository, transactionRepository,
					transactionAuthGateway,
					producer,
					domain.NewTransactionEntity(domain.NewWalletEntity()),
					domain.NewTransactionEntity(domain.NewWalletEntity()),
				)
				service.input = dto.TransactionDtoInput{
					Value:   "10000000",
					PayerID: userID,
				}
				payeeMailer, list_errors := service.executeTransactionPayee(context.Background())
				for index := range list_errors {
					assert.NoError(t, list_errors[index])
					assert.True(t, strings.Contains(payeeMailer, "@"))
				}
			})
		})
	})
}
