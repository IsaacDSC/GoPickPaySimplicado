package dto

import "github.com/google/uuid"

type TransactionDtoInput struct {
	Value   string    `json:"value"`
	PayerID uuid.UUID `json:"payer"`
	PayeeID uuid.UUID `json:"payee"`
}
