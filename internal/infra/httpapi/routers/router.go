package routers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

type Router interface {
	Load()
}

func MakeRouter(
	helloWorldRouter *HelloWorldRouter,
	userRouter *UserRouter,
) *fiber.App {

	cfg := fiber.Config{
		AppName:       "Assina.Backend",
		CaseSensitive: true,
		Prefork:       false,
	}

	r := fiber.New(cfg)

	r.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "Origin, Content-Type, Accept, Authorization",
		AllowMethods: "GET, POST, PUT, PATCH, DELETE, OPTIONS",
	}))

	r.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	helloWorldRouter.Load(r)
	userRouter.Load(r)

	return r
}
