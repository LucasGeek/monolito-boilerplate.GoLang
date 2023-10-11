package shared

import (
	"strconv"
	"strings"
)

// IsValidCPF checks if the given CPF string is valid.
func IsValidCPF(cpf string) bool {
	cpf = strings.Join(strings.Fields(cpf), "")
	if len(cpf) != 11 || strings.Count(cpf, cpf[0:1]) == 11 {
		return false
	}

	var sum, remainder int
	multipliers1 := []int{10, 9, 8, 7, 6, 5, 4, 3, 2}
	multipliers2 := []int{11, 10, 9, 8, 7, 6, 5, 4, 3, 2}

	for i := range multipliers1 {
		val, _ := strconv.Atoi(string(cpf[i]))
		sum += val * multipliers1[i]
	}

	remainder = sum % 11
	if remainder < 2 {
		remainder = 0
	} else {
		remainder = 11 - remainder
	}

	digit1 := strconv.Itoa(remainder)

	sum = 0
	for i := range multipliers2 {
		val, _ := strconv.Atoi(string(cpf[i]))
		sum += val * multipliers2[i]
	}

	remainder = sum % 11
	if remainder < 2 {
		remainder = 0
	} else {
		remainder = 11 - remainder
	}

	digit2 := strconv.Itoa(remainder)

	return cpf == cpf[:9]+digit1+digit2
}
