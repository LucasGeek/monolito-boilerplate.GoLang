package middleware

import (
	"server/src/commons/shared"
	"strings"

	"github.com/gofiber/fiber/v2"
)

type JWTMiddleware struct {
	manager *shared.JWTManager
}

// NewJWTMiddleware cria um novo middleware para validação de JWT.
func NewJWTMiddleware(manager *shared.JWTManager) fiber.Handler {
	return (&JWTMiddleware{manager: manager}).Validate
}

// Validate é um middleware do Fiber que valida o JWT token em cada requisição.
func (j *JWTMiddleware) Validate(c *fiber.Ctx) error {
	// Extrai o token do header "Authorization", que frequentemente vem como "Bearer <token>"
	authHeader := c.Get("Authorization")
	if authHeader == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "autenticação requerida"})
	}

	splitToken := strings.Split(authHeader, "Bearer ")
	if len(splitToken) != 2 {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "formato de token inválido"})
	}

	token := splitToken[1]
	if _, err := j.manager.Verify(token); err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "token inválido"})
	}

	return c.Next() // Continue para o próximo middleware ou rota.
}
