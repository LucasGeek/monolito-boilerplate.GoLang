package api

import (
	"github.com/gofiber/contrib/swagger"
	"log"
	"server/src/layers/app/di"
	"server/src/layers/app/handlers"
	"server/src/layers/app/middleware"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/monitor"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

type FiberServer struct {
	App       *fiber.App
	Container *di.Container
}

func NewFiberServer(container *di.Container) *FiberServer {
	app := fiber.New(fiber.Config{
		// You can define some initial configurations here, if necessary.
	})

	// global middlewares
	setupGlobalMiddlewares(app)

	return &FiberServer{
		App:       app,
		Container: container,
	}
}

func setupGlobalMiddlewares(app *fiber.App) {
	cfg := swagger.Config{
		BasePath: "/",
		FilePath: "./docs/swagger.json",
		Path:     "swagger",
		Title:    "Swagger Boilerplate API Docs",
	}

	app.Use(
		swagger.New(cfg),
		recover.New(),
		logger.New(),
	)
}

func (server *FiberServer) SetupRoutes() {
	server.App.Get("/", monitor.New(monitor.Config{Title: "api metricas"}))

	server.setupAuthRoutes()
	server.setupUserRoutes()
}

func (server *FiberServer) setupAuthRoutes() {
	authHandler := handlers.NewAuthHandler(
		server.Container.AuthHandler.CreateUser,
		server.Container.AuthHandler.CreateToken,
	)

	server.App.Post("/sign-up", authHandler.SignUp)
	server.App.Post("/sign-in", authHandler.SignIn)
}

func (server *FiberServer) setupUserRoutes() {
	jwtMiddleware := middleware.NewJWTMiddleware(server.Container.JWT)
	secureGroup := server.App.Group("/users", jwtMiddleware)

	userHandler := handlers.NewUserHandler(
		server.Container.UserHandler.GetUser,
	)

	secureGroup.Get("/:id", userHandler.Get)
	secureGroup.Get("/", userHandler.GetAll)
}

func (server *FiberServer) Run(port int) {
	address := ":" + strconv.Itoa(port)

	log.Printf("iniciando o servidor na porta %s...\n", address)

	if err := server.App.Listen(address); err != nil {
		log.Fatalf("falha ao iniciar o servidor: %v", err)
	}
}
