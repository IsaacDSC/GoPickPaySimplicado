package service

import (
	"context"
	"fmt"
	"time"

	"github.com/IsaacDSC/GoPickPaySimplicado/external/configs/queue"
	"github.com/IsaacDSC/GoPickPaySimplicado/internal/domain"
	"github.com/IsaacDSC/GoPickPaySimplicado/internal/infra/contracts"
	"github.com/IsaacDSC/GoPickPaySimplicado/internal/shared/dto"
)

type TransactionService struct {
	userRepo         contracts.UserRepositoryInterface
	transactionRepo  contracts.TransactionRepositoriesInterface
	checkOperation   contracts.OperationTransactionGatewayInterface
	notification     queue.IProducerQueue
	transactionPayer *domain.TransactionEntity //TODO: implementar contratos
	transactionPayee *domain.TransactionEntity //TODO: implementar contratos
	input            dto.TransactionDtoInput
}

func NewTransactionService(
	userRepo contracts.UserRepositoryInterface,
	transactionRepo contracts.TransactionRepositoriesInterface,
	checkOperation contracts.OperationTransactionGatewayInterface,
	notification queue.IProducerQueue,
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
	defer ts.transactionRepo.Rollback()
	ts.input = input
	payeeMailer, list_errors := ts.executeTransactionPayee(ctx)
	if len(list_errors) != 0 {
		return
	}
	payerMailer, list_errors := ts.executeTransactionPayer(ctx)
	if len(list_errors) != 0 {
		return
	}
	status := ts.checkOperation.TransactionAuth()
	fmt.Println("status", status)
	transactionAuthStatus := "COMPLETE"
	if status != "AUTHORIZED" {
		transactionAuthStatus = "TRANSACTION-AUTH-NOT-AUTHORIZED"
	}
	go ts.transactionRepo.UpdateStatusTransaction(ctx, ts.transactionPayee.ID, transactionAuthStatus)
	ts.transactionRepo.UpdateStatusTransaction(ctx, ts.transactionPayer.ID, transactionAuthStatus)
	ts.transactionRepo.Done()
	go ts.notification.TransactionNotificationMailer(ts.transactionPayee.ID, ts.transactionPayee.Operation, payeeMailer)
	go ts.notification.TransactionNotificationMailer(ts.transactionPayer.ID, ts.transactionPayer.Operation, payerMailer)

	//RECEBER DADOS DTO (PAYEE PAYER VALUE)
	//SAVE TRANSACTION PAYER
	//SAVE TRANSACTION PAYEE
	//CHECK-OPERATION
	//UPDATE TRANSACTION PAYER
	//UPDATE TRANSACTION PAYEE
	//NOTIFY EMAIL
	return
}

func (ts *TransactionService) executeTransactionPayer(ctx context.Context) (mailer string, list_errors []error) {
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
	ts.transactionPayer.TransactionFactory(
		payer.ID, payer.TypeUser, ts.input.Value, true, transactions,
	)
	list_errors = ts.transactionPayer.Transaction()
	if len(list_errors) != 0 {
		return
	}
	time.Sleep(time.Millisecond * time.Duration(ts.input.Sleep))
	err = ts.transactionRepo.InsertTransaction(ctx, ts.transactionPayer.Get())
	if err != nil {
		list_errors = append(list_errors, err)
		return
	}
	mailer = payer.Email
	return
}

func (ts *TransactionService) executeTransactionPayee(ctx context.Context) (mailer string, list_errors []error) {
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
		payee.ID, payee.TypeUser, ts.input.Value, false, transactions,
	)
	list_errors = ts.transactionPayee.Transaction()
	if len(list_errors) != 0 {
		return
	}
	time.Sleep(time.Millisecond * time.Duration(ts.input.Sleep))
	transaction := ts.transactionPayee.Get()
	err = ts.transactionRepo.InsertTransaction(ctx, transaction)
	if err != nil {
		list_errors = append(list_errors, err)
		return
	}
	status := ts.checkOperation.TransactionAuth()
	if status != "AUTHORIZED" {
		err := ts.transactionRepo.UpdateStatusTransaction(ctx, transaction.ID, status)
		list_errors = append(list_errors, err)
		return
	}
	err = ts.transactionRepo.UpdateStatusTransaction(ctx, transaction.ID, "COMPLETE")
	if err != nil {
		list_errors = append(list_errors, err)
	}
	mailer = payee.Email
	return
}
