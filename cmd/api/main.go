package main

import (
	"path/filepath"

	"github.com/IsaacDSC/GoPickPaySimplicado/external/configs/env"
	"github.com/IsaacDSC/GoPickPaySimplicado/external/configs/queue"
	web "github.com/IsaacDSC/GoPickPaySimplicado/web/http"
	"github.com/IsaacDSC/GoPickPaySimplicado/web/tasks"
	"github.com/google/uuid"
)

func init() {
	dotenv, _ := filepath.Abs("./.env")
	env.StartEnv(dotenv)
}

func main() {
	go queue.Consumer()
	go tasks.StartServerTasks()
	producer := queue.NewProducerQueue(queue.AsyncClientConn())
	producer.TransactionNotificationMailer(
		uuid.New(), "isaacdsc@gmail.com",
	)
	web.StartServerHttp()
}
