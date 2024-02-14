package router

import (
	"github.com/gofiber/fiber/v2"
	controller "github.com/krostyle/auth-systme-go/src/interfaces/controllers"
)

func SetupRouter(app *fiber.App, permissionController *controller.PermissionController) {
	permissionRoutes := app.Group(("/permission"))
	permissionRoutes.Post("/", permissionController.CreatePermission)
}
