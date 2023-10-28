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
	ctl := gomock.NewController(t)
	defer ctl.Finish()
	t.Run("Given init transaction payer", func(t *testing.T) {
		userRepository := mocks.NewMockUserRepositoryInterface(ctl)
		transactionAuthGateway := mocks.NewMockOperationTransactionGatewayInterface(ctl)
		producer := mocks.NewMockIProducerQueue(ctl)
		t.Run("When positive balancer and type user payer is consumer ", func(t *testing.T) {
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
			t.Run("Then should transaction not return error", func(t *testing.T) {
				service.input = dto.TransactionDtoInput{
					Value:   "10000000",
					PayerID: userID,
				}
				payerMailer, list_errors := service.executeTransactionPayer(context.Background())
				for index := range list_errors {
					assert.NoError(t, list_errors[index])
					assert.True(t, strings.Contains(payerMailer, "@"))
				}
				transaction := service.transactionPayer.Get()
				assert.Equal(t, "DEBIT", transaction.Operation)
				assert.Equal(t, "10000000", transaction.Value)
			})
		})

		t.Run("When storekeeper payer to consumer", func(t *testing.T) {
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
			t.Run("Then should transaction return error not authorized", func(t *testing.T) {
				_, list_errors := service.executeTransactionPayer(context.Background())
				for index := range list_errors {
					assert.Error(t, list_errors[index])
					assert.EqualError(t, list_errors[index], "NOT-AUTHORIZED-STOREKEEPER-PAYER-TO-CONSUMER")
				}
			})
		})

		t.Run("When insufficient balancer to payer", func(t *testing.T) {
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
						Operation: "DEBIT",
					},
				}, nil,
			)
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
			t.Run("Then should return error unauthorized transfer-wallet", func(t *testing.T) {
				_, list_errors := service.executeTransactionPayer(context.Background())
				for index := range list_errors {
					assert.Error(t, list_errors[index])
					assert.EqualError(t, list_errors[index], "UNAUTHORIZED-TRANSFER-WALLET")
				}
			})
		})
	})

	t.Run("Given init transaction payee", func(t *testing.T) {
		userRepository := mocks.NewMockUserRepositoryInterface(ctl)
		transactionAuthGateway := mocks.NewMockOperationTransactionGatewayInterface(ctl)
		producer := mocks.NewMockIProducerQueue(ctl)
		t.Run("When type user storekeeper", func(t *testing.T) {
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
						TypeUser:  "STOREKEEPER",
						Status:    "COMPLETE",
						Value:     "50000000",
						Operation: "CREDIT",
					},
					{
						ID:        uuid.New(),
						TypeUser:  "STOREKEEPER",
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
			t.Run("Then should return correctly mailer and not error", func(t *testing.T) {
				mailer, list_errors := service.executeTransactionPayee(context.Background())
				for index := range list_errors {
					assert.NoError(t, list_errors[index])
				}
				assert.Equal(t, "advino@gamil.com", mailer)
				transaction := service.transactionPayee.Get()
				assert.Equal(t, "CREDIT", transaction.Operation)
				assert.Equal(t, "10000000", transaction.Value)
			})
		})

		t.Run("When transaction receiver payee is consumer", func(t *testing.T) {
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
			t.Run("Then should not return error and set transaction credit payee", func(t *testing.T) {
				mailer, list_errors := service.executeTransactionPayee(context.Background())
				for index := range list_errors {
					assert.NoError(t, list_errors[index])
				}
				assert.Equal(t, "advino@gamil.com", mailer)
				transaction := service.transactionPayee.Get()
				assert.Equal(t, "CREDIT", transaction.Operation)
				assert.Equal(t, "10000000", transaction.Value)
			})
		})
	})

	//TODO: implementation complete transaction service
	// t.Run("Given init transfer transactions", func(t *testing.T) {
	// 	userRepository := mocks.NewMockUserRepositoryInterface(ctl)
	// 	transactionAuthGateway := mocks.NewMockOperationTransactionGatewayInterface(ctl)
	// 	producer := mocks.NewMockIProducerQueue(ctl)
	// 	t.Run("When", func(t *testing.T) {
	// 		userID := uuid.New()
	// 		userRepository.EXPECT().GetUserByID(gomock.Any(), gomock.Any()).Return(
	// 			domain.UserEntity{
	// 				ID:           userID,
	// 				CompleteName: "Advino Silva Name",
	// 				Cpf_cnpj:     "1233213123123123",
	// 				TypeUser:     "CONSUMER",
	// 				Email:        "advino@gamil.com",
	// 			}, nil,
	// 		)
	// 		transactionRepository := mocks.NewMockTransactionRepositoriesInterface(ctl)
	// 		transactionRepository.EXPECT().GetTransactionsByUserID(gomock.Any(), gomock.Any()).Return(
	// 			[]domain.Transactions{
	// 				{
	// 					ID:        uuid.New(),
	// 					TypeUser:  "CONSUMER",
	// 					Status:    "COMPLETE",
	// 					Value:     "50000000",
	// 					Operation: "CREDIT",
	// 				},
	// 				{
	// 					ID:        uuid.New(),
	// 					TypeUser:  "CONSUMER",
	// 					Status:    "COMPLETE",
	// 					Value:     "50000000",
	// 					Operation: "CREDIT",
	// 				},
	// 			}, nil,
	// 		)
	// 		transactionRepository.EXPECT().InsertTransaction(gomock.Any(), gomock.Any()).Return(nil)
	// 		service := NewTransactionService(
	// 			userRepository, transactionRepository,
	// 			transactionAuthGateway,
	// 			producer,
	// 			domain.NewTransactionEntity(domain.NewWalletEntity()),
	// 			domain.NewTransactionEntity(domain.NewWalletEntity()),
	// 		)
	// 		t.Run("Then ", func(t *testing.T) {
	// 			list_errors := service.Transfer(context.Background(), dto.TransactionDtoInput{
	// 				Value:   "10000000",
	// 				PayerID: userID,
	// 				PayeeID: userID,
	// 			})
	// 			for index := range list_errors {
	// 				assert.NoError(t, list_errors[index])
	// 			}
	// 		})
	// 	})
	// })

}
