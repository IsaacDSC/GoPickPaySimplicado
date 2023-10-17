package domain

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/IsaacDSC/GoPickPaySimplicado/external/sqlc"
	"github.com/google/uuid"
)

// var status = map[string][]string{
// 	"status": []string{"CREATED", "NOT_AUTHORIZED", "AWAIT_PERMISSION", "DELIVERY", "NOTIFICATION", "COMPLETE"},
// }

type TransactionEntity struct {
	ID              uuid.UUID
	UserID          uuid.UUID
	TypeUser        string
	totalBalancer   string
	Status          string
	Value           string //TODO: implement BIG_FLOAT
	Operation       string
	aggregateWallet *WalletEntity
	transactions    []Transactions
	list_errors     []error
}

func NewTransactionEntity(aggregateWallet *WalletEntity) *TransactionEntity {
	return &TransactionEntity{
		aggregateWallet: aggregateWallet,
	}
}

func (te *TransactionEntity) ToDomain(input sqlc.GetTransactionByUserIDRow) TransactionEntity {
	return TransactionEntity{
		ID:        input.ID,
		TypeUser:  input.TypeUser,
		Status:    input.Status,
		Value:     input.Value,
		UserID:    input.UserID,
		Operation: input.Operation.String,
	}
}

func (te *TransactionEntity) TransactionFactory(
	typeUser string,
	value string,
	isPayer bool,
	transactions []Transactions,
) {
	if isPayer {
		te.Operation = "DEBIT"
	} else {
		te.Operation = "CREDIT"
	}
	te.TypeUser = typeUser
	te.Status = "CREATED"
	te.Value = value
	te.transactions = transactions
}

func (te *TransactionEntity) Transaction() (list_errors []error) {
	te.aggregateWallet.WalletFactoryEntity(te.transactions)
	te.validateValue()
	if len(te.list_errors) != 0 {
		return
	}
	te.validateTypeUserTransaction()
	if len(te.list_errors) != 0 {
		return
	}
	if te.Operation == "DEBIT" && te.TypeUser == "CONSUMER" {
		err := te.aggregateWallet.CalculateTotalBalancer()
		if err != nil {
			list_errors = append(list_errors, err)
			te.Status = "NOT_AUTHORIZED"
			return
		}
		te.totalBalancer = te.aggregateWallet.totalBalancer
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
		te.list_errors = append(te.list_errors, errors.New(msg))
		te.Status = "NOT-AUTHORIZED"
		return
	}
	if integerValuer <= 0 {
		te.list_errors = append(te.list_errors, errors.New("VALUE NOT AUTHORIZED"))
		te.Status = "NOT-AUTHORIZED"
		return
	}
}

func (te *TransactionEntity) validateOperationDebit() {
	value, err := strconv.ParseFloat(te.Value, 64)
	if err != nil {
		te.list_errors = append(te.list_errors, errors.New("Error-Parse-StringToFloat"))
		te.Status = "NOT-AUTHORIZED"
		return
	}
	totalBalancer, err := strconv.ParseFloat(te.totalBalancer, 64)
	if err != nil {
		te.list_errors = append(te.list_errors, errors.New("Error-Parse-StringToFloat"))
		te.Status = "NOT-AUTHORIZED"
		return
	}
	if value > totalBalancer {
		te.list_errors = append(te.list_errors, errors.New("UNAUTHORIZED-TRANSFER-WALLET"))
		te.Status = "NOT-AUTHORIZED"
		return
	}
}

func (te *TransactionEntity) validateTypeUserTransaction() {
	if te.Operation == "DEBIT" && te.TypeUser != "CONSUMER" {
		msg := fmt.Sprintf("NOT-AUTHORIZED-TYPE-USER-%s-OPERATION-%s", te.TypeUser, te.Operation)
		te.list_errors = append(te.list_errors, errors.New(msg))
		return
	}
}
