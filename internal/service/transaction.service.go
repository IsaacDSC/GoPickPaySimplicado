package service

import (
	"context"
	"fmt"

	"github.com/IsaacDSC/GoPickPaySimplicado/internal/domain"
	"github.com/IsaacDSC/GoPickPaySimplicado/internal/infra/contracts"
	"github.com/IsaacDSC/GoPickPaySimplicado/internal/shared/dto"
)

type TransactionServiceInterface interface {
	Transfer() []error
}

type TransactionService struct {
	userRepo         contracts.UserRepositoryInterface
	transactionRepo  contracts.TransactionRepositoriesInterface
	checkOperation   contracts.OperationTransactionGatewayInterface
	notification     contracts.NotificationMailerInterface
	transactionPayer *domain.TransactionEntity //TODO: implementar contratos
	transactionPayee *domain.TransactionEntity //TODO: implementar contratos
	input            dto.TransactionDtoInput
}

func NewTransactionService(
	userRepo contracts.UserRepositoryInterface,
	transactionRepo contracts.TransactionRepositoriesInterface,
	checkOperation contracts.OperationTransactionGatewayInterface,
	notification contracts.NotificationMailerInterface,
	transactionPayer *domain.TransactionEntity,
	transactionPayee *domain.TransactionEntity,
) *TransactionService {
	return &TransactionService{
		userRepo,
		transactionRepo,
		checkOperation,
		notification,
		transactionPayer,
		transactionPayee,
		dto.TransactionDtoInput{},
	}
}

func (ts *TransactionService) Transfer(ctx context.Context, input dto.TransactionDtoInput) (list_errors []error) {
	ts.input = input
	list_errors = ts.executeTransactionPayee(ctx)
	if len(list_errors) != 0 {
		fmt.Println("primeiro error")
		return
	}
	list_errors = ts.executeTransactionPayer(ctx)
	if len(list_errors) != 0 {
		fmt.Println("segundo error")
		return
	}
	//RECEBER DADOS DTO (PAYEE PAYER VALUE)
	//CHECK-OPERATION
	//SAVE TRANSACTION PAYER
	// SAVE TRANSACTION PAYEE
	//NOTIFY EMAIL
	return
}

func (ts *TransactionService) executeTransactionPayer(ctx context.Context) (list_errors []error) {
	payer, err := ts.userRepo.GetUserByID(ctx, ts.input.PayerID)
	if err != nil {
		list_errors = append(list_errors, err)
		return
	}
	transactionsPayer, err := ts.transactionRepo.GetTransactionsByUserID(ctx, payer.ID)
	if err != nil {
		list_errors = append(list_errors, err)
		return
	}
	var transactions []domain.Transactions
	for index := range transactionsPayer {
		transactions = append(transactions, domain.Transactions{
			ID:        transactionsPayer[index].ID,
			TypeUser:  transactionsPayer[index].TypeUser,
			Status:    transactionsPayer[index].Status,
			Value:     transactionsPayer[index].Value,
			Operation: transactionsPayer[index].Operation,
		})
	}
	ts.transactionPayee.TransactionFactory(
		payer.TypeUser, ts.input.Value, true, transactions,
	)
	list_errors = ts.transactionPayee.Transaction()
	if len(list_errors) != 0 {
		return
	}
	err = ts.transactionRepo.InsertTransaction(ctx, *ts.transactionPayee)
	if err != nil {
		list_errors = append(list_errors, err)
		return
	}
	return
}

func (ts *TransactionService) executeTransactionPayee(ctx context.Context) (list_errors []error) {
	payee, err := ts.userRepo.GetUserByID(ctx, ts.input.PayeeID)
	if err != nil {
		list_errors = append(list_errors, err)
		return
	}
	transactionsPayee, err := ts.transactionRepo.GetTransactionsByUserID(ctx, payee.ID)
	if err != nil {
		list_errors = append(list_errors, err)
		return
	}
	var transactions []domain.Transactions
	for index := range transactionsPayee {
		transactions = append(transactions, domain.Transactions{
			ID:        transactionsPayee[index].ID,
			TypeUser:  transactionsPayee[index].TypeUser,
			Status:    transactionsPayee[index].Status,
			Value:     transactionsPayee[index].Value,
			Operation: transactionsPayee[index].Operation,
		})
	}
	ts.transactionPayee.TransactionFactory(
		payee.TypeUser, ts.input.Value, false, transactions,
	)
	list_errors = ts.transactionPayee.Transaction()
	if len(list_errors) != 0 {
		return
	}
	err = ts.transactionRepo.InsertTransaction(ctx, *ts.transactionPayer)
	if err != nil {
		list_errors = append(list_errors, err)
		return
	}
	return
}
