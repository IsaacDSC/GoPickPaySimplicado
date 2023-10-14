package domain

import (
	"github.com/IsaacDSC/GoPickPaySimplicado/external/sqlc"
	"github.com/google/uuid"
)

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

func (*UserEntity) ToDomain(input sqlc.User) UserEntity {
	return UserEntity{
		ID:           input.ID,
		CompleteName: input.CompleteName.String,
		Cpf_cnpj:     input.CpfCnpj,
		TypeUser:     input.TypeUser,
		Email:        input.Email,
	}
}
