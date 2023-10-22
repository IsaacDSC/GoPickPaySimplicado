package http_controllers

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/IsaacDSC/GoPickPaySimplicado/internal/service"
	"github.com/IsaacDSC/GoPickPaySimplicado/internal/shared/container"
	"github.com/IsaacDSC/GoPickPaySimplicado/internal/shared/dto"
)

type TransactionController struct {
	Path    string
	service service.TransactionServiceInterface
}

func NewTransactionController(
	path string,
) *TransactionController {
	return &TransactionController{
		Path: path,
	}
}

func (tc *TransactionController) PostTransaction(w http.ResponseWriter, r *http.Request) {
	tc.service = container.NewTransactionContainer()
	defer r.Body.Close()
	payload, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		return
	}
	var body dto.TransactionDtoInput
	err = json.Unmarshal(payload, &body)
	// fmt.Printf("\n%+v\n", body)
	if err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		return
	}
	list_errors := tc.service.Transfer(r.Context(), body)
	if len(list_errors) > 0 {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(list_errors[0].Error()))
		return
	}
	w.WriteHeader(http.StatusCreated)

}
