package dto

import "github.com/google/uuid"

type TransactionDtoInput struct {
	Value   string    `json:"value"`
	PayerID uuid.UUID `json:"payerId"`
	PayeeID uuid.UUID `json:"payeeId"`
	Sleep   int       `json:"sleep"`
}
