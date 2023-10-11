package commands

import (
	"errors"
	"fmt"
	"github.com/google/uuid"
	"server/src/commons/shared"
	"server/src/layers/domain/repository"
)

type CreateTokenHandler struct {
	Repo         repository.UserRepository
	ArgonManager *shared.Argon2Manager
	JWT          *shared.JWTManager
}

// CreateTokenCommand representa a intenção de criar um token para um usuário existente
type CreateTokenCommand struct {
	CPF      string `json:"Cpf"`
	Password string `json:"Password"`
}

type TokenResponse struct {
	User SimplifiedUser `json:"User"`
	Key  SimplifiedKey  `json:"Key"`
}

type SimplifiedUser struct {
	ID        uuid.UUID `json:"ID"`
	FirstName string    `json:"FirstName"`
	LastName  string    `json:"LastName"`
	CPF       string    `json:"Cpf"`
}

type SimplifiedKey struct {
	Token        string `json:"Token"`
	RefreshToken string `json:"RefreshToken"`
}

// Validate realiza validações básicas no comando CreateTokenCommand
func (c *CreateTokenCommand) Validate() error {
	fields := map[string]interface{}{
		"Cpf":      c.CPF,
		"Password": c.Password,
	}

	for fieldName, value := range fields {
		if value == "" {
			return fmt.Errorf("%s é necessário", fieldName)
		}
	}
	return nil
}

// Handle processa o comando CreateTokenCommand e gera um JWT para o usuário
func (c *CreateTokenHandler) Handle(command CreateTokenCommand) (*TokenResponse, error) {
	// Busca o usuário com base no cpf fornecido
	user, err := c.Repo.FindByCPF(command.CPF)
	if err != nil || user == nil {
		return nil, errors.New("cpf ou senha inválidos")
	}

	// Compara a senha fornecida com a hash armazenada usando argon2
	match, err := c.ArgonManager.VerifyPassword(command.Password, user.Password)
	if err != nil || !match {
		return nil, errors.New("cpf ou senha inválidos")
	}

	// Gera o JWT para o usuário
	token, refreshToken, err := c.JWT.Generate(user.ID)
	if err != nil {
		return nil, errors.New("falha ao gerar o token")
	}

	response := &TokenResponse{
		User: SimplifiedUser{
			ID:        user.ID,
			FirstName: user.FirstName,
			LastName:  user.LastName,
			CPF:       user.CPF,
		},
		Key: SimplifiedKey{
			Token:        token,
			RefreshToken: refreshToken,
		},
	}

	return response, nil
}
