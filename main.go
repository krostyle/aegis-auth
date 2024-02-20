package main

import (
	"fmt"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/krostyle/auth-systme-go/src/infrastructure/config"
	"github.com/krostyle/auth-systme-go/src/infrastructure/middleware"
)

func main() {
	app := fiber.New()
	app.Use(middleware.CORS())
	config.Setup(app)
	port := os.Getenv("PORT")
	env := os.Getenv("ENV")
	if env != "production" {
		port = "3000"
	}

	fmt.Println("Server running on port: " + port)
	if err := app.Listen(":" + port); err != nil {
		panic(err)
	}
}
