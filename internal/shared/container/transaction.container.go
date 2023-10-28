package container

import (
	"github.com/IsaacDSC/GoPickPaySimplicado/external/configs/queue"
	"github.com/IsaacDSC/GoPickPaySimplicado/internal/domain"
	"github.com/IsaacDSC/GoPickPaySimplicado/internal/infra/gateway"
	"github.com/IsaacDSC/GoPickPaySimplicado/internal/infra/repositories"
	"github.com/IsaacDSC/GoPickPaySimplicado/internal/service"
)

func NewTransactionContainer() service.TransactionServiceInterface {
	return service.NewTransactionService(
		repositories.NewUserRepository(),
		repositories.NewTransactionRepositories(),
		gateway.NewOperationTransactionGateway(),
		queue.NewProducerQueue(queue.AsyncClientConn()),
		domain.NewTransactionEntity(domain.NewWalletEntity()),
		domain.NewTransactionEntity(domain.NewWalletEntity()),
	)
}
