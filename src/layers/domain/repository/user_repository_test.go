package repository

import (
	"fmt"
	"github.com/google/uuid"
	"server/src/layers/domain/models"
	"testing"
)

func TestMockUserRepository_Store(t *testing.T) {
	repo := NewMockUserRepository()
	user := &models.User{CPF: "83103569009", FirstName: "Lucas", LastName: "Albuquerque"}

	// Teste para adicionar um novo usuário
	err := repo.Store(user)
	if err != nil {
		t.Fatalf("Erro ao armazenar o usuário: %v", err)
	}

	// Teste para verificar se o usuário já existe
	err = repo.Store(user)
	if err != ErrUserExists {
		t.Fatalf("Esperado erro de usuário já existe, mas obteve: %v", err)
	}
}

func TestMockUserRepository_FindByID(t *testing.T) {
	repo := NewMockUserRepository()
	user := &models.User{CPF: "83103569009", FirstName: "Lucas", LastName: "Albuquerque"}
	repo.Store(user)

	// Imprima todos os usuários no repositório após a inserção
	fmt.Printf("Usuários no repositório: %+v\n", repo.users)
	// Imprima o ID do usuário após a inserção
	fmt.Printf("ID do usuário inserido: %v\n", user.ID)

	t.Run("Buscar usuário existente por ID", func(t *testing.T) {
		foundUser, err := repo.FindByID(user.ID)
		if err != nil {
			t.Errorf("Erro ao buscar usuário por ID: %v", err)
		}
		if foundUser.ID != user.ID {
			t.Errorf("Usuário encontrado não corresponde ao esperado")
		}
	})

	// Teste para buscar usuário com ID inexistente
	t.Run("Buscar usuário inexistente por ID", func(t *testing.T) {
		_, err := repo.FindByID(uuid.New())
		if err != ErrUserNotFound {
			t.Errorf("Esperado erro de usuário não encontrado, mas obteve: %v", err)
		}
	})
}

func TestMockUserRepository_FindByCPF(t *testing.T) {
	repo := NewMockUserRepository()
	user := &models.User{CPF: "83103569009", FirstName: "Lucas", LastName: "Albuquerque"}
	repo.Store(user)

	// Teste para buscar usuário por CPF
	foundUser, err := repo.FindByCPF("83103569009")
	if err != nil {
		t.Fatalf("Erro ao buscar usuário por CPF: %v", err)
	}
	if foundUser.CPF != user.CPF {
		t.Fatalf("Usuário encontrado não corresponde ao esperado")
	}

	// Teste para buscar usuário com CPF inexistente
	_, err = repo.FindByCPF("111.222.333-44")
	if err != ErrUserNotFound {
		t.Fatalf("Esperado erro de usuário não encontrado, mas obteve: %v", err)
	}
}

func TestMockUserRepository_Update(t *testing.T) {
	repo := NewMockUserRepository()
	user := &models.User{CPF: "83103569009", FirstName: "Lucas", LastName: "Albuquerque"}
	repo.Store(user)

	// Atualizar usuário
	user.FirstName = "Jane"
	err := repo.Update(user)
	if err != nil {
		t.Fatalf("Erro ao atualizar usuário: %v", err)
	}
	updatedUser, _ := repo.FindByID(user.ID)
	if updatedUser.FirstName != "Jane" {
		t.Fatalf("Usuário não foi atualizado corretamente")
	}

	// Teste para atualizar usuário inexistente
	user.ID = uuid.New()
	err = repo.Update(user)
	if err != ErrUserNotFound {
		t.Fatalf("Esperado erro de usuário não encontrado, mas obteve: %v", err)
	}
}

func TestMockUserRepository_UpdatePassword(t *testing.T) {
	repo := NewMockUserRepository()
	user := &models.User{CPF: "83103569009", FirstName: "Lucas", LastName: "Albuquerque", Password: "oldPassword"}
	repo.Store(user)

	// Atualizar senha do usuário
	err := repo.UpdatePassword(user.ID, "newPassword")
	if err != nil {
		t.Fatalf("Erro ao atualizar senha: %v", err)
	}
	updatedUser, _ := repo.FindByID(user.ID)
	if updatedUser.Password != "newPassword" {
		t.Fatalf("Senha do usuário não foi atualizada corretamente")
	}
}

func TestMockUserRepository_Delete(t *testing.T) {
	repo := NewMockUserRepository()
	user := &models.User{CPF: "83103569009", FirstName: "Lucas", LastName: "Albuquerque"}
	repo.Store(user)

	// Deletar usuário
	err := repo.Delete(user.ID)
	if err != nil {
		t.Fatalf("Erro ao deletar usuário: %v", err)
	}
	_, err = repo.FindByID(user.ID)
	if err != ErrUserNotFound {
		t.Fatalf("Usuário não foi deletado corretamente")
	}
}

func TestMockUserRepository_FindAllWithPagination(t *testing.T) {
	repo := NewMockUserRepository()
	for i := 0; i < 10; i++ {
		user := &models.User{CPF: fmt.Sprintf("123.456.789-%02d", i), FirstName: "Lucas", LastName: "Albuquerque"}
		repo.Store(user)
	}

	// Buscar todos os usuários com paginação
	users, err := repo.FindAllWithPagination(5, 0)
	if err != nil {
		t.Fatalf("Erro ao buscar usuários: %v", err)
	}
	if len(users) != 5 {
		t.Fatalf("Esperado 5 usuários, mas obteve: %d", len(users))
	}
}
