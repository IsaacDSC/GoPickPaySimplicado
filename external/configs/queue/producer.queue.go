package queue

import (
	"log"
	"time"

	"github.com/IsaacDSC/GoPickPaySimplicado/external/configs/queue/task"
	"github.com/google/uuid"
	"github.com/hibiken/asynq"
)

type IProducerQueue interface {
	TransactionNotificationMailer(transactionID uuid.UUID, operation string, mailer string)
}

type ProducerQueue struct {
	client *asynq.Client
}

func NewProducerQueue(client *asynq.Client) *ProducerQueue {
	return &ProducerQueue{
		client,
	}
}

func (pq *ProducerQueue) TransactionNotificationMailer(
	transactionID uuid.UUID, operation string, mailer string,
) {
	delay := 1 * time.Minute
	taskNotificationMailer := task.NewNotificationTransactionEmailTask(
		transactionID, mailer, operation, time.Now().Add(delay),
	)

	if _, err := pq.client.Enqueue(
		taskNotificationMailer,
		asynq.Queue("low"),
		asynq.ProcessIn(delay),
	); err != nil {
		log.Fatal(err)
	}
}
