package persistence

import (
	"fmt"
	"github.com/google/uuid"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"server/src/layers/domain/models"
	"testing"
)

func setupDatabase() (*gorm.DB, error) {
	return gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{})
}

func TestUserRepository_Store(t *testing.T) {
	db, _ := setupDatabase()
	repo := NewUserRepository(db)
	db.AutoMigrate(&models.User{})

	user := &models.User{
		CPF:       "83103569009",
		Password:  "password",
		FirstName: "Lucas",
		LastName:  "Albuquerque",
	}

	storedUser, err := repo.Store(user)
	if err != nil {
		t.Fatalf("Erro ao armazenar o usuário: %v", err)
	}
	if storedUser.CPF != user.CPF {
		t.Fatalf("CPF do usuário armazenado não corresponde. Esperado: %s, Obteve: %s", user.CPF, storedUser.CPF)
	}
}

func TestUserRepository_FindByID(t *testing.T) {
	db, _ := setupDatabase()
	repo := NewUserRepository(db)
	db.AutoMigrate(&models.User{})

	user := &models.User{
		CPF:       "83103569009",
		Password:  "password",
		FirstName: "Lucas",
		LastName:  "Albuquerque",
	}
	repo.Store(user)

	t.Run("Find by valid ID", func(t *testing.T) {
		foundUser, err := repo.FindByID(user.ID)
		if err != nil {
			t.Fatalf("Erro ao buscar o usuário pelo ID: %v", err)
		}
		if foundUser.ID != user.ID {
			t.Fatalf("Usuário encontrado não corresponde ao esperado.")
		}
	})

	t.Run("Find by invalid ID", func(t *testing.T) {
		_, err := repo.FindByID(uuid.New()) // Using a new random ID.
		if err == nil {
			t.Fatalf("Esperava um erro ao buscar um usuário inexistente, mas não obteve nenhum.")
		}
	})
}

func TestUserRepository_FindByCPF(t *testing.T) {
	db, _ := setupDatabase()
	repo := NewUserRepository(db)
	db.AutoMigrate(&models.User{})

	user := &models.User{
		CPF:       "83103569009",
		Password:  "password",
		FirstName: "Lucas",
		LastName:  "Albuquerque",
	}
	repo.Store(user)

	t.Run("Find by valid CPF", func(t *testing.T) {
		foundUser, err := repo.FindByCPF(user.CPF)
		if err != nil {
			t.Fatalf("Erro ao buscar o usuário pelo CPF: %v", err)
		}
		if foundUser.CPF != user.CPF {
			t.Fatalf("Usuário encontrado não corresponde ao esperado.")
		}
	})

	t.Run("Find by invalid CPF", func(t *testing.T) {
		_, err := repo.FindByCPF("999.999.999-99") // Using a random CPF.
		if err == nil {
			t.Fatalf("Esperava um erro ao buscar um usuário inexistente, mas não obteve nenhum.")
		}
	})
}

func TestUserRepository_Update(t *testing.T) {
	db, _ := setupDatabase()
	repo := NewUserRepository(db)
	db.AutoMigrate(&models.User{})

	user := &models.User{
		CPF:       "83103569009",
		Password:  "password",
		FirstName: "Lucas",
		LastName:  "Albuquerque",
	}
	repo.Store(user)

	t.Run("Update user details", func(t *testing.T) {
		user.FirstName = "Lucas"
		err := repo.Update(user)
		if err != nil {
			t.Fatalf("Erro ao atualizar o usuário: %v", err)
		}

		updatedUser, err := repo.FindByID(user.ID)
		if err != nil {
			t.Fatalf("Erro ao buscar o usuário atualizado: %v", err)
		}
		if updatedUser.FirstName != "Lucas" {
			t.Fatalf("Usuário não foi atualizado corretamente.")
		}
	})
}

func TestUserRepository_UpdatePassword(t *testing.T) {
	db, _ := setupDatabase()
	repo := NewUserRepository(db)
	db.AutoMigrate(&models.User{})

	user := &models.User{
		CPF:       "83103569009",
		Password:  "password",
		FirstName: "Lucas",
		LastName:  "Albuquerque",
	}
	repo.Store(user)

	t.Run("Update user password", func(t *testing.T) {
		newPassword := "new_password"
		err := repo.UpdatePassword(user.ID, newPassword)
		if err != nil {
			t.Fatalf("Erro ao atualizar a senha do usuário: %v", err)
		}

		updatedUser, _ := repo.FindByID(user.ID)
		if updatedUser.Password != newPassword {
			t.Fatalf("Senha do usuário não foi atualizada corretamente.")
		}
	})
}

func TestUserRepository_Delete(t *testing.T) {
	db, _ := setupDatabase()
	repo := NewUserRepository(db)
	db.AutoMigrate(&models.User{})

	user := &models.User{
		CPF:       "83103569009",
		Password:  "password",
		FirstName: "Lucas",
		LastName:  "Albuquerque",
	}
	repo.Store(user)

	t.Run("Delete user", func(t *testing.T) {
		err := repo.Delete(user.ID)
		if err != nil {
			t.Fatalf("Erro ao deletar o usuário: %v", err)
		}

		_, err = repo.FindByID(user.ID)
		if err == nil {
			t.Fatalf("Usuário não foi deletado corretamente.")
		}
	})
}

func TestUserRepository_FindAllWithPagination(t *testing.T) {
	db, _ := setupDatabase()
	repo := NewUserRepository(db)
	db.AutoMigrate(&models.User{})

	for i := 0; i < 10; i++ {
		user := &models.User{
			CPF:       fmt.Sprintf("123.456.789-%02d", i),
			Password:  "password",
			FirstName: fmt.Sprintf("User%d", i),
			LastName:  "Test",
		}
		_, err := repo.Store(user)
		if err != nil {
			t.Fatalf("Erro ao armazenar o usuário %d: %v", i, err)
		}
	}

	t.Run("Find users with pagination", func(t *testing.T) {
		users, err := repo.FindAllWithPagination(5, 0)
		if err != nil {
			t.Fatalf("Erro ao buscar usuários com paginação: %v", err)
		}
		if len(users) != 5 {
			t.Fatalf("Esperado 5 usuários, mas obteve: %d", len(users))
		}
	})
}
