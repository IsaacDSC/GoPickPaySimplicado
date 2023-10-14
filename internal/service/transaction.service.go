package service

import (
	"github.com/IsaacDSC/GoPickPaySimplicado/internal/domain"
	"github.com/IsaacDSC/GoPickPaySimplicado/internal/infra/contracts"
)

type TransactionServiceInterface interface {
	Transfer() []error
}

type TransactionService struct {
	userRepo        contracts.UserRepositoryInterface
	transactionRepo contracts.TransactionRepositoriesInterface
	checkOperation  contracts.OperationTransactionGatewayInterface
	notification    contracts.NotificationMailerInterface
	transaction     domain.TransactionEntity //TODO: implementar contratos
	wallet          domain.WalletEntity      //TODO: implementar contratos
}

func (*TransactionService) Transfer() (list_errors []error) {
	//RECEBER DADOS DTO (PAYEE PAYER VALUE)
	//BUSCAR PAYER
	//BUSCAR PAYEE
	// VALIDAR TRANSACTION PAYER
	// VALIDAR TRANSACTION PAYEE
	//CHECK-OPERATION
	//SAVE TRANSACTION PAYER
	// SAVE TRANSACTION PAYEE
	//NOTIFY EMAIL
	return
}
