package di

import (
	"gorm.io/gorm"
	"server/src/commons/config"
	"server/src/commons/shared"
	"server/src/layers/app/handlers"
	"server/src/layers/domain/repository"
	"server/src/layers/infrastructure/persistence"
	"server/src/layers/service/commands"
	"server/src/layers/service/queries"
	"time"

	"github.com/gofiber/fiber/v2/log"
)

type Container struct {
	AuthHandler  handlers.AuthHandler
	UserHandler  handlers.UserHandler
	JWT          *shared.JWTManager
	Argon2Config Argon2Config
}

// InitializeContainer configura todas as dependências para o aplicativo.
func InitializeContainer() *Container {
	cfg := config.LoadConfig()

	db := connectToDatabase()
	//defer persistence.Close(db)

	jwtManager := shared.NewJWTManager(cfg.JWTSecret, 24*time.Hour, (7*24)*time.Hour)
	argonManager := shared.NewArgon2Manager()

	userRepo := persistence.NewUserRepository(db)

	userHandler := initializeUserHandler(userRepo)
	authHandler := initializeAuthHandler(argonManager, jwtManager, userRepo)

	argonConfig := DefaultArgon2Config()

	return &Container{
		AuthHandler:  authHandler,
		UserHandler:  userHandler,
		JWT:          jwtManager,
		Argon2Config: argonConfig,
	}
}

// connectToDatabase estabelece uma conexão com o banco de dados.
func connectToDatabase() *gorm.DB {
	db, err := persistence.Connect()
	if err != nil {
		log.Fatalf("falha ao conectar ao banco de dados: %v", err)
	}
	return db
}

// initializeUserHandler cria um novo UserHandler com suas dependências necessárias.
func initializeUserHandler(repo repository.UserRepository) handlers.UserHandler {
	getUserQueryHandler := queries.GetUserQueryHandler{Repo: repo}

	return *handlers.NewUserHandler(getUserQueryHandler)
}

// initializeAuthHandler cria um novo AuthHandler com suas dependências necessárias.
func initializeAuthHandler(argonManager *shared.Argon2Manager, jwtManager *shared.JWTManager, repo repository.UserRepository) handlers.AuthHandler {
	createTokenHandler := commands.CreateTokenHandler{
		ArgonManager: argonManager,
		JWT:          jwtManager,
		Repo:         repo,
	}

	createUserHandler := commands.CreateUserHandler{
		ArgonManager: argonManager,
		Repo:         repo,
	}

	return *handlers.NewAuthHandler(createUserHandler, createTokenHandler)
}

type Argon2Config struct {
	Time    uint32
	Memory  uint32
	Threads uint8
	KeyLen  uint32
}

// DefaultArgon2Config fornece uma configuração padrão para Argon2.
func DefaultArgon2Config() Argon2Config {
	return Argon2Config{
		Time:    1,
		Memory:  64 * 1024,
		Threads: 4,
		KeyLen:  32,
	}
}
