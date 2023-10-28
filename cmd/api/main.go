package main

import (
	"path/filepath"

	"github.com/IsaacDSC/GoPickPaySimplicado/external/configs/env"
	"github.com/IsaacDSC/GoPickPaySimplicado/external/configs/queue"
	web "github.com/IsaacDSC/GoPickPaySimplicado/web/http"
	"github.com/IsaacDSC/GoPickPaySimplicado/web/tasks"
)

func init() {
	dotenv, _ := filepath.Abs("./.env")
	env.StartEnv(dotenv)
}

func main() {
	go queue.Consumer()
	go tasks.StartServerTasks()
	web.StartServerHttp()
}
