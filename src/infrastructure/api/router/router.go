package router

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/krostyle/auth-systme-go/src/adapters/controller"
)

func SetupRouter(
	app *fiber.App,
	userController *controller.UserController,
	permissionController *controller.PermissionController,
	healthCheckController *controller.HealthCheckController) {

	fmt.Println("Setting up health check routes")
	healthCheckRoutes := app.Group("/health")
	healthCheckRoutes.Get("/", healthCheckController.HealthCheck)

	fmt.Println("Setting up user routes")
	userRoutes := app.Group("/users")
	userRoutes.Post("/", userController.RegisterUser)
	userRoutes.Get("/", userController.GetAllUsers)
	userRoutes.Post("/login", userController.Login)

	fmt.Println("Setting up permission routes")
	permissionRoutes := app.Group("/permissions")
	permissionRoutes.Post("/", permissionController.CreatePermission)
	permissionRoutes.Get("/", permissionController.GetAllPermissions)

}
