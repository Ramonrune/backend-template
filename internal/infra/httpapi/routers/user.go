package routers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ramonrune/assina-backend/internal/infra/httpapi/controllers"
)

type UserRouter struct {
	controller *controllers.UserController
}

func (c *UserRouter) Load(r *fiber.App) {
	r.Get("/api/v1/user", c.controller.Get)
	r.Post("/api/v1/user", c.controller.Create)
	r.Put("/api/v1/user", c.controller.Update)
}

func NewUserRouter(
	controller *controllers.UserController,
) *UserRouter {
	return &UserRouter{
		controller: controller,
	}
}
