package httpapi

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/valyala/fasthttp"
)

func NewServer(
	router *fiber.App,
) *fasthttp.Server {
	addr := fmt.Sprintf(":%s", "8080")

	go func() {
		if err := router.Listen(addr); err != nil {
			log.Fatalf("Failed to start server: %v", err)
		}
		log.Printf("Servidor rodando na porta %s", addr)
	}()
	return router.Server()
}
