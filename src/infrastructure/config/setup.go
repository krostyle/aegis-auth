package setup

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/krostyle/auth-systme-go/src/application/usecase"
	"github.com/krostyle/auth-systme-go/src/infrastructure/api/router"
	"github.com/krostyle/auth-systme-go/src/infrastructure/database/postgres_client/configuration"
	"github.com/krostyle/auth-systme-go/src/infrastructure/database/postgres_client/repository"
	controller "github.com/krostyle/auth-systme-go/src/interfaces/controllers"
)

func Setup(app *fiber.App) {

	gormDB, err := configuration.NewPostgresDB()
	if err != nil {
		panic(fmt.Sprintf("Error connecting to database: %v", err))
	}
	permissionRepository := repository.NewPermissionRepository(gormDB)
	permissionUseCase := usecase.NewPermissionUseCase(permissionRepository)
	permissionController := controller.NewPermissionController(permissionUseCase)
	router.SetupRouter(app, permissionController)

}
