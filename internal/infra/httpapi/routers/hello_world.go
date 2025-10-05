package routers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ramonrune/assina-backend/internal/infra/httpapi/controllers"
)

type HelloWorldRouter struct {
	controller *controllers.HelloWorldController
}

func (c *HelloWorldRouter) Load(r *fiber.App) {
	r.Get("/api/v1/hello-world", c.controller.Get)
}

func NewHelloWorldRouter(
	controller *controllers.HelloWorldController,
) *HelloWorldRouter {
	return &HelloWorldRouter{
		controller: controller,
	}
}
