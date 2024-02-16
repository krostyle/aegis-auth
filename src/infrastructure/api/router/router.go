package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/krostyle/auth-systme-go/src/adapters/controller"
)

func SetupRouter(
	app *fiber.App,
	permissionController *controller.PermissionController,
	roleController *controller.RoleController,
	userController *controller.UserController) {

	permissionRoutes := app.Group(("/permissions"))
	permissionRoutes.Post("/", permissionController.CreatePermission)
	permissionRoutes.Get("/:id", permissionController.GetPermissionByID)
	permissionRoutes.Put("/:id", permissionController.UpdatePermission)
	permissionRoutes.Delete("/:id", permissionController.DeletePermission)
	permissionRoutes.Get("/", permissionController.GetAllPermissions)

	roleRoutes := app.Group(("/roles"))
	roleRoutes.Post("/", roleController.CreateRole)
	roleRoutes.Get("/", roleController.GetAllRoles)
	roleRoutes.Get("/:id", roleController.GetRoleByID)
	roleRoutes.Put("/:id", roleController.UpdateRole)
	roleRoutes.Delete("/:id", roleController.DeleteRole)

	userRoutes := app.Group(("/users"))
	userRoutes.Post("/", userController.CreateUser)
	userRoutes.Get("/", userController.GetAllUsers)
	userRoutes.Get("/:id", userController.GetUserByID)
	userRoutes.Put("/:id", userController.UpdateUser)
	userRoutes.Delete("/:id", userController.DeleteUser)
}
