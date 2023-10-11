package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	JWTSecret string
	Port      int
}

func LoadConfig() *Config {
	// Carregar variáveis de ambiente do arquivo .env
	if err := godotenv.Load(); err != nil {
		log.Printf("erro ao carregar arquivo .env: %v", err) // Interrompe a execução aqui se o arquivo .env não for encontrado.
	}

	return &Config{
		JWTSecret: getEnv("JWT_SECRET", "api_secret"),
		Port:      getEnvAsInt("PORT", 3333),
	}
}

// getEnv obtém uma variável de ambiente ou retorna um valor padrão.
func getEnv(key, defaultValue string) string {
	value, exists := os.LookupEnv(key)
	if !exists {
		log.Printf("variável de ambiente %s não definida, usando valor padrão: %s", key, defaultValue)
		return defaultValue
	}
	return value
}

// getEnvAsInt tenta obter e converter uma variável de ambiente para int, ou retorna um valor padrão.
func getEnvAsInt(key string, defaultValue int) int {
	valueStr := getEnv(key, "")
	value, err := strconv.Atoi(valueStr)
	if err != nil {
		log.Printf("erro ao converter %s em int: %v, usando valor padrão: %d", key, err, defaultValue)
		return defaultValue
	}
	return value
}
