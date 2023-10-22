package handler

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/hibiken/asynq"
)

func HandlerNotificationTransactionEmail(c context.Context, t *asynq.Task) error {
	// Get int with the user ID from the given task.
	input := t.Payload()
	var body map[string]interface{}
	err := json.Unmarshal(input, &body)
	if err != nil {
		return err
	}

	// Dummy message to the worker's output.
	fmt.Printf("Reason: time is up (%v)\n", body)

	return nil
}
