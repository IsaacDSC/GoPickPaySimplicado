package queue

import (
	"fmt"
	"log"

	"github.com/IsaacDSC/GoPickPaySimplicado/external/configs/queue/handler"
	"github.com/IsaacDSC/GoPickPaySimplicado/external/configs/queue/task"
	"github.com/hibiken/asynq"
)

func Consumer() {
	worker := AsyncServerConn()

	mux := asynq.NewServeMux()

	mux.HandleFunc(
		task.NotificationTransactionMailerEvent,
		handler.HandlerNotificationTransactionEmail,
	)

	fmt.Println("[ * ] started consumer task")
	if err := worker.Run(mux); err != nil {
		log.Fatal(err)
	}

}
