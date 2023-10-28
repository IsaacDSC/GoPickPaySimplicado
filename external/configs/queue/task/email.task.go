package task

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/hibiken/asynq"
)

const (
	// NotificationTransactionEmailTask is a name of the task type
	// for sending a notification email the transaction.
	NotificationTransactionMailerEvent = "email:transaction"
)

func NewNotificationTransactionEmailTask(
	transactionID uuid.UUID, mailer, operation string, sentIn time.Time,
) *asynq.Task {
	payload := map[string]interface{}{
		"transaction_id": transactionID,
		"mailer":         mailer,
		"operation":      operation,
		"sent_in":        sentIn.String(),
	}
	input, err := json.Marshal(payload)
	if err != nil {
		fmt.Printf("Error converter payload: %w\n", err.Error())
	}

	return asynq.NewTask(
		NotificationTransactionMailerEvent,
		input, asynq.MaxRetry(5),
		asynq.Timeout(1*time.Minute),
	)
}
