package main

import (
	"log"
	"tryhardplayoffs/internal/middleware"
	"tryhardplayoffs/internal/services"

	"github.com/labstack/echo/v4"
)

func main() {
	app := echo.New()

	ctn := services.NewContainer()

	defer ctn.Shutdown()

	app.Use(middleware.ContainerMiddleware(ctn))

	if err := app.Start(":8080"); err != nil {
		log.Fatal("Error starting server:", err)
	}
}
