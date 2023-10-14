package domain

import (
	"fmt"
	"strconv"

	"github.com/google/uuid"
)

var status = map[string][]string{
	"status": []string{"CREATED", "NOT_AUTHORIZED", "AWAIT_PERMISSION", "DELIVERY", "NOTIFICATION", "COMPLETE"},
}

type TransactionEntity struct {
	ID              uuid.UUID
	TypeUser        string
	TotalBalancer   string
	Status          string
	Value           string //TODO: implement BIG_FLOAT
	Operation       string
	aggregateWallet WalletEntity
	list_errors     []string
}

func NewTransactionEntity() *TransactionEntity {
	return new(TransactionEntity)
}

func (te *TransactionEntity) TransactionFactory(
	typeUser string,
	value string,
	operation string,
	aggregateWallet WalletEntity,
) {
	te.TypeUser = typeUser
	te.Status = "CREATED"
	te.Value = value
	te.Operation = operation
	te.aggregateWallet = aggregateWallet
}

func (te *TransactionEntity) Transaction() (list_errors []string) {
	te.validateValue()
	if te.Operation == "DEBIT" && te.TypeUser == "CONSUMER" {
		err := te.aggregateWallet.CalculateTotalBalancer()
		if err != nil {
			list_errors = append(list_errors, err.Error())
			te.Status = "NOT_AUTHORIZED"
			return
		}
		te.TotalBalancer = te.aggregateWallet.totalBalancer
		te.validateOperationDebit()
	}
	list_errors = te.list_errors
	if (len(te.list_errors)) == 0 {
		te.Status = "AWAIT_PERMISSION"
	}
	return
}

func (te *TransactionEntity) validateValue() {
	integerValuer, err := strconv.Atoi(te.Value)
	if err != nil {
		msg := fmt.Sprintf("VALUE: %s - NOT AUTHORIZED", te.Value)
		te.list_errors = append(te.list_errors, msg)
		te.Status = "NOT-AUTHORIZED"
		return
	}
	if integerValuer <= 0 {
		te.list_errors = append(te.list_errors, "VALUE NOT AUTHORIZED")
		te.Status = "NOT-AUTHORIZED"
		return
	}
}

func (te *TransactionEntity) validateOperationDebit() {
	value, err := strconv.ParseFloat(te.Value, 64)
	if err != nil {
		te.list_errors = append(te.list_errors, "Error-Parse-StringToFloat")
		te.Status = "NOT-AUTHORIZED"
		return
	}
	totalBalancer, err := strconv.ParseFloat(te.TotalBalancer, 64)
	if err != nil {
		te.list_errors = append(te.list_errors, "Error-Parse-StringToFloat")
		te.Status = "NOT-AUTHORIZED"
		return
	}
	if value > totalBalancer {
		te.list_errors = append(te.list_errors, "UNAUTHORIZED-TRANSFER-WALLET")
		te.Status = "NOT-AUTHORIZED"
		return
	}
}
