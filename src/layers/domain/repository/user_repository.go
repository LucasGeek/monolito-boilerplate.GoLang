package repository

import (
	"errors"
	"fmt"
	"server/src/layers/domain/models"
)

var (
	ErrUserExists   = errors.New("usuário já existe")
	ErrUserNotFound = errors.New("usuário não encontrado")
)

// UserRepository define a interface que qualquer armazenamento de usuário deve implementar
type UserRepository interface {
	Store(user *models.User) (*models.User, error)
	FindByID(id string) (*models.User, error)
	FindByCPF(cpf string) (*models.User, error)
	Update(user *models.User) error
	UpdatePassword(userID, hashedPassword string) error
	Delete(id string) error
	FindAllWithPagination(limit int, offset int) ([]*models.User, error)
}

// MockUserRepository é uma implementação fictícia do UserRepository para testes
type MockUserRepository struct {
	users map[string]*models.User
}

// NewMockUserRepository cria uma nova instância do MockUserRepository
func NewMockUserRepository() *MockUserRepository {
	return &MockUserRepository{users: make(map[string]*models.User)}
}

// Store adiciona um novo usuário ao armazenamento fictício
func (m *MockUserRepository) Store(user *models.User) error {
	if _, exists := m.users[fmt.Sprintf("%d", user.ID)]; exists {
		return ErrUserExists
	}
	m.users[fmt.Sprintf("%d", user.ID)] = user
	return nil
}

// FindByID retorna um usuário pelo ID do armazenamento fictício
func (m *MockUserRepository) FindByID(id string) (*models.User, error) {
	if user, exists := m.users[id]; exists {
		return user, nil
	}
	return nil, ErrUserNotFound
}

// FindByCPF retorna um usuário pelo e-mail do armazenamento fictício
func (m *MockUserRepository) FindByCPF(cpf string) (*models.User, error) {
	for _, user := range m.users {
		if user.CPF == cpf {
			return user, nil
		}
	}
	return nil, ErrUserNotFound
}

// Update atualiza um usuário existente no armazenamento fictício
func (m *MockUserRepository) Update(user *models.User) error {
	if _, exists := m.users[fmt.Sprintf("%d", user.ID)]; !exists {
		return ErrUserNotFound
	}
	m.users[fmt.Sprintf("%d", user.ID)] = user
	return nil
}

func (m *MockUserRepository) UpdatePassword(userID string, hashedPassword string) error {
	user, exists := m.users[userID]
	if !exists {
		return ErrUserNotFound
	}
	user.Password = hashedPassword
	return nil
}

// Delete remove um usuário pelo ID do armazenamento fictício
func (m *MockUserRepository) Delete(id string) error {
	if _, exists := m.users[id]; !exists {
		return ErrUserNotFound
	}
	delete(m.users, id)
	return nil
}

func (m *MockUserRepository) FindAllWithPagination(limit int, offset int) ([]*models.User, error) {
	if offset >= len(m.users) {
		return []*models.User{}, nil
	}

	usersSlice := make([]*models.User, 0, len(m.users))
	for _, user := range m.users {
		usersSlice = append(usersSlice, user)
	}

	end := offset + limit
	if end > len(usersSlice) {
		end = len(usersSlice)
	}

	return usersSlice[offset:end], nil
}
