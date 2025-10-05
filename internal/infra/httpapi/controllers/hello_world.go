package controllers

import (
	"github.com/gofiber/fiber/v2"
)

type HelloWorldController struct {
}

func NewHelloWorldController() *HelloWorldController {
	return &HelloWorldController{}
}

func (cc *HelloWorldController) Get(c *fiber.Ctx) error {

	return c.Status(fiber.StatusOK).JSON(
		fiber.Map{
			"name": "ola mundo",
		},
	)
}
