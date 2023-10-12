package shared

import (
	"testing"
)

func TestIsValidCPF(t *testing.T) {
	tests := []struct {
		cpf      string
		expected bool
	}{
		{"52998224725", true},    // Válido
		{"52998224724", false},   // Inválido
		{"00000000000", false},   // Inválido (todos os dígitos são iguais)
		{"529 98 224 725", true}, // Válido (com espaços)
		{"5299822", false},       // Inválido (tamanho menor que 11)
		{"5299822472500", false}, // Inválido (tamanho maior que 11)
		{"52998224725 ", true},   // Válido (com espaço no final)
		{" 52998224725", true},   // Válido (com espaço no começo)
		{"52a98224725", false},   // Inválido (contém letras)
		{"", false},              // Inválido (string vazia)
	}

	for _, test := range tests {
		result := IsValidCPF(test.cpf)
		if result != test.expected {
			t.Errorf("Expected IsValidCPF(%s) to be %v, but got %v", test.cpf, test.expected, result)
		}
	}
}
