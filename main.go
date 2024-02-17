package main

import (
	"os"

	"github.com/gofiber/fiber/v2"
	setup "github.com/krostyle/auth-systme-go/src/infrastructure/config"
)

func main() {
	app := fiber.New()
	setup.Setup(app)
	port := os.Getenv("APP_PORT")
	if err := app.Listen(port); err != nil {
		panic(err)
	}
}
