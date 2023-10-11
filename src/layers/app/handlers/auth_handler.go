package handlers

import (
	"server/src/layers/service/commands"

	"github.com/gofiber/fiber/v2"
)

type AuthHandler struct {
	CreateUser  commands.CreateUserHandler
	CreateToken commands.CreateTokenHandler
}

func NewAuthHandler(createUser commands.CreateUserHandler, createToken commands.CreateTokenHandler) *AuthHandler {
	return &AuthHandler{
		CreateUser:  createUser,
		CreateToken: createToken,
	}
}

func (h *AuthHandler) SignUp(c *fiber.Ctx) error {
	var input createUserInput

	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": err.Error()})
	}

	newUserCommand := commands.CreateUserCommand{
		CPF:       input.CPF,
		Password:  input.Password,
		FirstName: input.FirstName,
		LastName:  input.LastName,
	}

	err := newUserCommand.Validate()
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": err.Error()})
	}

	user, err := h.CreateUser.Handle(newUserCommand)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": err.Error()})
	}

	return c.Status(fiber.StatusCreated).JSON(user)
}

func (h *AuthHandler) SignIn(c *fiber.Ctx) error {
	var input createTokenInput

	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": err.Error()})
	}

	newTokenCommand := commands.CreateTokenCommand{
		CPF:      input.CPF,
		Password: input.Password,
	}

	err := newTokenCommand.Validate()
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": err.Error()})
	}

	token, err := h.CreateToken.Handle(newTokenCommand)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(token)
}

type createUserInput struct {
	CPF       string `json:"Cpf"`
	Password  string `json:"Password"`
	FirstName string `json:"FirstName"`
	LastName  string `json:"LastName"`
}

type createTokenInput struct {
	CPF      string `json:"Cpf"`
	Password string `json:"Password"`
}
