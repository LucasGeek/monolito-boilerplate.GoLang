package models

import (
	"errors"
	"server/src/commons/shared"
)

// User representa o modelo de domínio para um usuário.
type User struct {
	Base
	CPF       string `json:"Cpf"`
	Password  string `json:"-"`
	FirstName string `json:"FirstName"`
	LastName  string `json:"LastName"`
}

// NewUser é um construtor para o modelo User.
func NewUser(cpf, firstName, lastName, password string) (*User, error) {
	if err := validateUserFields(cpf, firstName, lastName, password); err != nil {
		return nil, err
	}

	return &User{
		CPF:       cpf,
		Password:  password,
		FirstName: firstName,
		LastName:  lastName,
	}, nil
}

// validateUserFields verifica se os campos obrigatórios estão preenchidos e se o CPF é válido.
func validateUserFields(cpf, firstName, lastName, password string) error {
	if !shared.IsValidCPF(cpf) {
		return errors.New("Cpf com formato inválido")
	}
	if cpf == "" {
		return errors.New("Cpf deve ser informado")
	}
	if firstName == "" {
		return errors.New("Nome deve ser informado")
	}
	if lastName == "" {
		return errors.New("Sobrenome deve ser informado")
	}
	if password == "" {
		return errors.New("Senha deve ser informada")
	}
	return nil
}
