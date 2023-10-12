package handlers

import (
	"github.com/google/uuid"
	"server/src/layers/service/queries"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type UserHandler struct {
	GetUser queries.GetUserQueryHandler
}

// NewUserHandler retorna uma nova instância de UserHandler
func NewUserHandler(getUser queries.GetUserQueryHandler) *UserHandler {
	return &UserHandler{
		GetUser: getUser,
	}
}

// Get recupera informações do usuário com base no ID fornecido na URL
func (h *UserHandler) Get(c *fiber.Ctx) error {
	idStr := c.Params("id")

	id, err := uuid.Parse(idStr)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "ID inválido"})
	}

	query := queries.GetUserByIDQuery{
		UserID: id,
	}

	user, err := h.GetUser.GetUserByIDHandle(query)

	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"message": err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(user)
}

// Get recupera informações de varios usuários
func (h *UserHandler) GetAll(c *fiber.Ctx) error {
	const defaultLimit = 10
	const defaultOffset = 0

	limitStr := c.Params("limit")
	offsetStr := c.Params("offset")

	limit, err := strconv.Atoi(limitStr)
	if err != nil || limit <= 0 {
		limit = defaultLimit
	}

	offset, err := strconv.Atoi(offsetStr)
	if err != nil || offset < 0 {
		offset = defaultOffset
	}

	query := queries.GetAllUsersQuery{
		Limit:  limit,
		Offset: offset,
	}

	users, err := h.GetUser.GetAllUsersHandle(query)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"message": err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(users)
}
