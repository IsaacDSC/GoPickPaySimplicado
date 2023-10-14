package domain

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/google/uuid"
)

type Transactions struct {
	ID            uuid.UUID
	TypeUser      string
	TotalBalancer string
	Status        string
	Value         string //TODO: implement BIG_FLOAT
	Operation     string
}

type WalletEntity struct {
	ID                   uuid.UUID
	UserID               uuid.UUID
	AggregateTransaction []Transactions
	totalBalancer        string
}

func NewWalletEntity() *WalletEntity {
	return new(WalletEntity)
}

func (we *WalletEntity) WalletFactoryEntity(transactions []Transactions) {
	we.ID = uuid.New()
	we.AggregateTransaction = transactions
}

func (we *WalletEntity) GetTotalBalancer() string {
	return we.totalBalancer
}

func (we *WalletEntity) CalculateTotalBalancer() (err error) {
	var (
		operation     []string
		balancer      []int64
		totalBalancer int64
	)
	for index := range we.AggregateTransaction {
		if we.AggregateTransaction[index].Status != "COMPLETE" {
			err = errors.New("Transaction-In-Process-Retry-Again-Latter")
			break
		}
		integerValue, err := strconv.ParseInt(we.AggregateTransaction[index].Value, 10, 64)
		if err != nil {
			break
		}
		balancer = append(balancer, integerValue)
		operation = append(operation, we.AggregateTransaction[index].Operation)
	}
	if err != nil {
		return
	}
	for index := range balancer {
		if operation[index] != "CREDIT" && operation[index] != "DEBIT" {
			err = errors.New("Invalid-Operation")
			break
		}
		if operation[index] == "CREDIT" {
			totalBalancer += balancer[index]
		}
		if operation[index] == "DEBIT" {
			totalBalancer -= balancer[index]
		}
	}
	we.totalBalancer = fmt.Sprintf("%d", totalBalancer)
	return
}
