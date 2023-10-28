package gateway

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/IsaacDSC/GoPickPaySimplicado/internal/shared/dto"
)

type OperationTransactionGateway struct{}

func NewOperationTransactionGateway() *OperationTransactionGateway {
	return new(OperationTransactionGateway)
}

func (*OperationTransactionGateway) TransactionAuth() string {
	url := "https://run.mocky.io/v3/052ea2d9-e51c-4fcf-bcd7-db3ceb4395f6"
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Add("User-Agent", "api/picpay-simplificado")
	res, _ := http.DefaultClient.Do(req)
	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)
	if res.StatusCode != http.StatusOK {
		return "BAD-RESPONSE-TRANSACTION_AUTH"
	}
	var data dto.OperationAuthResponse
	if err := json.Unmarshal(body, &data); err != nil {
		return "UN-PROCESSABLE-ENTITY"
	}
	if data.Status != "AUTHORIZED" {
		return "UNAUTHORIZED"
	}
	return data.Status
}
