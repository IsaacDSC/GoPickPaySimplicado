package domain

import "github.com/google/uuid"

const storekeeper = "STOREKEEPER"
const consumer = "consumer"

type UserEntity struct {
	ID           uuid.UUID
	CompleteName string
	Cpf_cnpj     string
	TypeUser     string
	Email        string
}

func NewUserEntity() *UserEntity {
	return new(UserEntity)
}

func (*UserEntity) ToDomain() {}
