package persistence

import (
	"errors"
	"fmt"
	"gorm.io/gorm"
	"server/src/layers/domain/models"
)

// UserRepository representa o repositório de usuário.
type UserRepository struct {
	db *gorm.DB
}

// NewUserRepository cria uma nova instância de UserRepository.
func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

// Store insere um novo usuário e cria um evento relacionado.
func (ur *UserRepository) Store(user *models.User) (*models.User, error) {
	if err := ur.db.Create(user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

// FindByID busca um usuário pelo ID.
func (ur *UserRepository) FindByID(id string) (*models.User, error) {
	var user models.User
	if err := ur.db.First(&user, "id = ?", id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("usuário com o ID %s não encontrado", id)
		}
		return nil, err
	}
	return &user, nil
}

// FindByCPF busca um usuário pelo cpf.
func (ur *UserRepository) FindByCPF(cpf string) (*models.User, error) {
	var user models.User
	if err := ur.db.First(&user, "cpf = ?", cpf).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("usuário com o cpf %s não encontrado", cpf)
		}
		return nil, err
	}
	return &user, nil
}

// Update atualiza os detalhes do usuário e cria um evento relacionado.
func (ur *UserRepository) Update(user *models.User) error {
	return ur.db.Save(user).Error
}

// UpdatePassword atualiza a senha do usuário e cria um evento relacionado.
func (ur *UserRepository) UpdatePassword(userID string, hashedPassword string) error {
	return ur.db.Model(&models.User{}).Where("id = ?", userID).Update("password", hashedPassword).Error
}

// Delete remove um usuário e cria um evento relacionado.
func (ur *UserRepository) Delete(id string) error {
	return ur.db.Delete(&models.User{}, "id = ?", id).Error
}

// FindAllWithPagination busca todos os usuários com paginação.
func (ur *UserRepository) FindAllWithPagination(limit int, offset int) ([]*models.User, error) {
	var users []*models.User
	if err := ur.db.Limit(limit).Offset(offset).Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}
