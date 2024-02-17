package main

import (
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/krostyle/auth-systme-go/src/infrastructure/config"
)

func main() {
	app := fiber.New()
	config.Setup(app)
	port := os.Getenv("PORT")
	if err := app.Listen(":" + port); err != nil {
		panic(err)
	}
}
