package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"net/http"
	"server/src/commons/shared"
	"testing"
	"time"
)

const mockSecret = "test-secret"

var mockUserID = uuid.New()

func TestJWTMiddleware(t *testing.T) {
	app := fiber.New()

	manager := shared.NewJWTManager(mockSecret, time.Hour, 24*time.Hour)
	token, _, _ := manager.Generate(mockUserID)

	app.Use(NewJWTMiddleware(manager))

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	t.Run("No Authorization header", func(t *testing.T) {
		req, _ := http.NewRequest("GET", "/", nil)
		resp, err := app.Test(req)

		if err != nil {
			t.Fatalf("Expected no error, got %v", err)
		}

		if resp.StatusCode != fiber.StatusUnauthorized {
			t.Fatalf("Expected status %v, got %v", fiber.StatusUnauthorized, resp.StatusCode)
		}
	})

	t.Run("Invalid Token format", func(t *testing.T) {
		req, _ := http.NewRequest("GET", "/", nil)
		req.Header.Set("Authorization", "Token "+token)
		resp, err := app.Test(req)

		if err != nil {
			t.Fatalf("Expected no error, got %v", err)
		}

		if resp.StatusCode != fiber.StatusUnauthorized {
			t.Fatalf("Expected status %v, got %v", fiber.StatusUnauthorized, resp.StatusCode)
		}
	})

	t.Run("Valid Token", func(t *testing.T) {
		req, _ := http.NewRequest("GET", "/", nil)
		req.Header.Set("Authorization", "Bearer "+token)
		resp, err := app.Test(req)

		if err != nil {
			t.Fatalf("Expected no error, got %v", err)
		}

		if resp.StatusCode != fiber.StatusOK {
			t.Fatalf("Expected status %v, got %v", fiber.StatusOK, resp.StatusCode)
		}
	})

	t.Run("Invalid Token value", func(t *testing.T) {
		req, _ := http.NewRequest("GET", "/", nil)
		req.Header.Set("Authorization", "Bearer wrong.token.value")
		resp, err := app.Test(req)

		if err != nil {
			t.Fatalf("Expected no error, got %v", err)
		}

		if resp.StatusCode != fiber.StatusUnauthorized {
			t.Fatalf("Expected status %v, got %v", fiber.StatusUnauthorized, resp.StatusCode)
		}
	})
}
