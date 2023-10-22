package main

import (
	"fmt"
	"path/filepath"

	"github.com/IsaacDSC/GoPickPaySimplicado/external/configs/env"
	web "github.com/IsaacDSC/GoPickPaySimplicado/web/http"
)

func init() {
	dotenv, _ := filepath.Abs("./.env")
	env.StartEnv(dotenv)
}

func main() {
	fmt.Println("[ * ] start application 3000")
	web.StartServerHttp()
}
