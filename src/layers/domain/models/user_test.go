package models

import (
	"testing"
)

func TestNewUser(t *testing.T) {
	// Testes de sucesso
	validCPF := "83103569009" // Este CPF precisa ser válido de acordo com as regras do `IsValidCPF`.

	user, err := NewUser(validCPF, "Lucas", "Albuquerque", "password123")
	if err != nil {
		t.Fatalf("Falha ao criar usuário com dados válidos: %v", err)
	}
	if user.CPF != validCPF {
		t.Errorf("Esperado CPF %s, mas recebeu %s", validCPF, user.CPF)
	}

	// Testes de falha
	_, err = NewUser("", "Lucas", "Albuquerque", "password123")
	if err == nil || err.Error() != "Cpf deve ser informado" {
		t.Error("Esperado erro de CPF não informado")
	}

	_, err = NewUser("invalidCPF", "Lucas", "Albuquerque", "password123")
	if err == nil || err.Error() != "Cpf com formato inválido" {
		t.Error("Esperado erro de formato de CPF inválido")
	}

	_, err = NewUser(validCPF, "", "Albuquerque", "password123")
	if err == nil || err.Error() != "Nome deve ser informado" {
		t.Error("Esperado erro de nome não informado")
	}

	_, err = NewUser(validCPF, "Lucas", "", "password123")
	if err == nil || err.Error() != "Sobrenome deve ser informado" {
		t.Error("Esperado erro de sobrenome não informado")
	}

	_, err = NewUser(validCPF, "Lucas", "Albuquerque", "")
	if err == nil || err.Error() != "Senha deve ser informada" {
		t.Error("Esperado erro de senha não informada")
	}
}

func TestValidateUserFields(t *testing.T) {
	// Estes testes são similares aos testes de `NewUser`, já que `NewUser` chama `validateUserFields`.
	// No entanto, é uma boa prática testar funções isoladamente para garantir que cada uma funcione corretamente.

	validCPF := "83103569009" // Este CPF precisa ser válido de acordo com as regras do `IsValidCPF`.

	err := validateUserFields(validCPF, "Lucas", "Albuquerque", "password123")
	if err != nil {
		t.Errorf("Falha ao validar campos de usuário com dados válidos: %v", err)
	}

	err = validateUserFields("", "Lucas", "Albuquerque", "password123")
	if err == nil || err.Error() != "Cpf deve ser informado" {
		t.Error("Esperado erro de CPF não informado")
	}

	err = validateUserFields("invalidCPF", "Lucas", "Albuquerque", "password123")
	if err == nil || err.Error() != "Cpf com formato inválido" {
		t.Error("Esperado erro de formato de CPF inválido")
	}

	err = validateUserFields(validCPF, "", "Albuquerque", "password123")
	if err == nil || err.Error() != "Nome deve ser informado" {
		t.Error("Esperado erro de nome não informado")
	}

	err = validateUserFields(validCPF, "Lucas", "", "password123")
	if err == nil || err.Error() != "Sobrenome deve ser informado" {
		t.Error("Esperado erro de sobrenome não informado")
	}

	err = validateUserFields(validCPF, "Lucas", "Albuquerque", "")
	if err == nil || err.Error() != "Senha deve ser informada" {
		t.Error("Esperado erro de senha não informada")
	}
}
