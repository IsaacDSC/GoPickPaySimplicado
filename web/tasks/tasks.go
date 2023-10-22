package tasks

import (
	"fmt"
	"log"
	"net/http"

	"github.com/hibiken/asynq"
	"github.com/hibiken/asynqmon"
)

func StartServerTasks() {
	h := asynqmon.New(asynqmon.Options{
		RootPath:     "/monitoring", // RootPath specifies the root for asynqmon app
		RedisConnOpt: asynq.RedisClientOpt{Addr: ":6379"},
	})

	http.Handle(h.RootPath()+"/", h)
	fmt.Println("[ * ] started task 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
