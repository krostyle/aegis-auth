package router

import (
	"github.com/gofiber/fiber/v2"
	controller "github.com/krostyle/auth-systme-go/src/interfaces/controllers"
)

func SetupRouter(app *fiber.App, permissionController *controller.PermissionController) {
	permissionRoutes := app.Group(("/permissions"))
	permissionRoutes.Post("/", permissionController.CreatePermission)
	permissionRoutes.Get("/:id", permissionController.GetPermissionByID)
	permissionRoutes.Put("/:id", permissionController.UpdatePermission)
	permissionRoutes.Delete("/:id", permissionController.DeletePermission)
	permissionRoutes.Get("/", permissionController.GetAllPermissions)
}
