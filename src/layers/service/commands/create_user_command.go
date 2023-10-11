package commands

import (
	"errors"
	"fmt"
	"server/src/commons/shared"
	"server/src/layers/domain/models"
	"server/src/layers/domain/repository"
)

type CreateUserHandler struct {
	Repo         repository.UserRepository
	ArgonManager *shared.Argon2Manager
}

// CreateUserCommand representa a intenção de criar um novo usuário
type CreateUserCommand struct {
	FirstName string `json:"FirstName"`
	LastName  string `json:"LastName"`
	CPF       string `json:"Cpf"`
	Password  string `json:"Password"`
}

// Validate realiza validações básicas no comando CreateUserCommand
func (c *CreateUserCommand) Validate() error {
	fields := map[string]interface{}{
		"FirstName": c.FirstName,
		"LastName":  c.LastName,
		"Password":  c.Password,
		"Cpf":       c.CPF,
	}

	for fieldName, value := range fields {
		if value == "" {
			return fmt.Errorf("%s é necessário", fieldName)
		}
	}
	return nil
}

func (h *CreateUserHandler) Handle(command CreateUserCommand) (*models.User, error) {
	// Busca o usuário com base no cpf fornecido
	user, _ := h.Repo.FindByCPF(command.CPF)
	if user != nil {
		return nil, errors.New("cpf já cadastrado")
	}

	// Hash da senha com argon2
	hashedPassword, err := h.ArgonManager.HashPassword(command.Password)
	if err != nil {
		return nil, errors.New("erro ao criptografar a senha")
	}

	// Convertendo 'command' para um 'User'
	newUser, err := models.NewUser(command.CPF, command.FirstName, command.LastName, hashedPassword)
	if err != nil {
		return nil, err
	}

	return h.Repo.Store(newUser)
}
