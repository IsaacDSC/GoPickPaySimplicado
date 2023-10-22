package http_controllers

import "net/http"

type HealthCheckController struct {
	Path string
}

func NewHealthCheckController(path string) *HealthCheckController {
	controller := new(HealthCheckController)
	controller.Path = path
	return controller
}

func (*HealthCheckController) Index(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Alive"))
}
