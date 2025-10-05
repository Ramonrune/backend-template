package controllers

import (
	"github.com/go-playground/validator"
	"github.com/gofiber/fiber/v2"
	"github.com/ramonrune/assina-backend/internal/app/domain/user"
)

type UserController struct {
	service *user.UserService
}

var validateUser = validator.New()

type CreateUserRequest struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (dto *CreateUserRequest) ValidateCreate() error {
	return validateUser.Struct(dto)
}

type UpdateUserRequest struct {
	Name string `json:"name"`
}

func (dto *UpdateUserRequest) ValidateUpdate() error {
	return validateUser.Struct(dto)
}

func (cc *UserController) Create(c *fiber.Ctx) error {

	var dto CreateUserRequest

	if err := c.BodyParser(&dto); err != nil {
		return c.
			Status(fiber.StatusBadRequest).
			JSON(fiber.Map{"error": "bad request"})
	}

	if err := dto.ValidateCreate(); err != nil {
		return c.
			Status(fiber.StatusBadRequest).
			JSON(fiber.Map{"error": err.Error()})
	}

	user, err := cc.service.Create(dto.Name, dto.Email, dto.Password)

	if err != nil {
		return c.
			Status(fiber.StatusInternalServerError).
			JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusCreated).JSON(user)
}

func NewUserController(service *user.UserService) *UserController {
	return &UserController{
		service: service,
	}
}
