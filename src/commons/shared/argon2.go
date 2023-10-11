package shared

import (
	"crypto/rand"
	"encoding/base64"
	"errors"
	"strings"

	"golang.org/x/crypto/argon2"
)

const (
	// Constantes de configuração para o Argon2. Ajuste conforme necessário.
	SaltSize  = 16
	KeySize   = 32
	Time      = 1
	Memory    = 64 * 1024
	Threads   = 4
	separator = "$"
)

type Argon2Manager struct{}

func NewArgon2Manager() *Argon2Manager {
	return &Argon2Manager{}
}

// generateSalt cria um novo salt aleatório
func (a *Argon2Manager) generateSalt() ([]byte, error) {
	salt := make([]byte, SaltSize)
	if _, err := rand.Read(salt); err != nil {
		return nil, err
	}
	return salt, nil
}

// HashPassword cria e retorna um hash Argon2 da senha fornecida
func (a *Argon2Manager) HashPassword(password string) (string, error) {
	salt, err := a.generateSalt()
	if err != nil {
		return "", err
	}

	hash := argon2.IDKey([]byte(password), salt, Time, Memory, Threads, KeySize)

	// O hash retornado é uma combinação do salt e hash, separados por "$"
	encodedSalt := base64.RawStdEncoding.EncodeToString(salt)
	encodedHash := base64.RawStdEncoding.EncodeToString(hash)
	return encodedSalt + separator + encodedHash, nil
}

// VerifyPassword verifica se a senha fornecida corresponde ao hash Argon2 fornecido
func (a *Argon2Manager) VerifyPassword(password, encodedHash string) (bool, error) {
	parts := strings.Split(encodedHash, separator)
	if len(parts) != 2 {
		return false, errors.New("formato de hash inválido")
	}

	salt, err := base64.RawStdEncoding.DecodeString(parts[0])
	if err != nil {
		return false, err
	}

	expectedHash, err := base64.RawStdEncoding.DecodeString(parts[1])
	if err != nil {
		return false, err
	}

	calculatedHash := argon2.IDKey([]byte(password), salt, Time, Memory, Threads, KeySize)

	return string(calculatedHash) == string(expectedHash), nil
}
