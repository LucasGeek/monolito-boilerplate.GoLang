package queries

import (
	"errors"
	"server/src/layers/domain/models"
	"server/src/layers/domain/repository"
)

type GetUserQueryHandler struct {
	Repo repository.UserRepository
}

// GetUserByIDQuery representa a consulta para obter um usuário pelo ID
type GetUserByIDQuery struct {
	UserID string `json:"ID"`
}

// GetUserByCPFQuery representa a consulta para obter um usuário pelo cpf
type GetUserByCPFQuery struct {
	CPF string `json:"Cpf"`
}

// GetAllUsersQuery representa uma consulta para obter todos os usuários
type GetAllUsersQuery struct {
	Limit  int `json:"Limit"`  // limita o número de resultados retornados
	Offset int `json:"Offset"` // permite paginação dos resultados
}

func (g *GetUserQueryHandler) GetUserByIDHandle(query GetUserByIDQuery) (*models.User, error) {
	user, err := g.Repo.FindByID(query.UserID)
	if err != nil {
		if errors.Is(err, repository.ErrUserNotFound) {
			return nil, errors.New("usuário não encontrado")
		}
		return nil, errors.New("erro interno do servidor")
	}
	return user, nil
}

func (g *GetUserQueryHandler) GetUserByCPFHandle(query GetUserByCPFQuery) (*models.User, error) {
	user, err := g.Repo.FindByCPF(query.CPF)
	if err != nil {
		if errors.Is(err, repository.ErrUserNotFound) {
			return nil, errors.New("usuário não encontrado")
		}
		return nil, errors.New("erro interno do servidor")
	}
	return user, nil
}

func (g *GetUserQueryHandler) GetAllUsersHandle(query GetAllUsersQuery) ([]*models.User, error) {
	users, err := g.Repo.FindAllWithPagination(query.Limit, query.Offset)
	if err != nil {
		return nil, err
	}
	return users, nil
}
