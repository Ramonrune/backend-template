package main

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/ramonrune/assina-backend/internal/infra/httpapi"
	"github.com/ramonrune/assina-backend/internal/infra/httpapi/controllers"
	"github.com/ramonrune/assina-backend/internal/infra/httpapi/routers"
	"github.com/valyala/fasthttp"
	"go.uber.org/fx"
)

func main() {
	err := godotenv.Load(".env")

	if err != nil {
		log.Printf("Couldn't load .env file")
	}

	app := fx.New(
		controllers.Module,
		routers.Module,
		httpapi.Module,
		fx.Invoke(func(*fasthttp.Server) {}),
	)

	app.Run()
}
