package setup

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	controller "github.com/krostyle/auth-systme-go/src/adapters/controllers"
	"github.com/krostyle/auth-systme-go/src/application/usecase"
	"github.com/krostyle/auth-systme-go/src/infrastructure/api/router"
	"github.com/krostyle/auth-systme-go/src/infrastructure/database/postgres_client/configuration"
	"github.com/krostyle/auth-systme-go/src/infrastructure/database/postgres_client/repository"
	"github.com/krostyle/auth-systme-go/src/infrastructure/service"
)

func Setup(app *fiber.App) {

	err := godotenv.Load()
	if err != nil {
		panic(fmt.Sprintf("Error loading .env file: %v", err))
	}

	gormDB, err := configuration.NewPostgresDB()
	if err != nil {
		panic(fmt.Sprintf("Error connecting to database: %v", err))
	}

	identifierGenerator := service.NewUUIDGenerator()

	permissionRepository := repository.NewPermissionRepository(gormDB)
	permissionUseCase := usecase.NewPermissionUseCase(permissionRepository, identifierGenerator)
	permissionController := controller.NewPermissionController(permissionUseCase)

	roleRepository := repository.NewRoleRepository(gormDB)
	roleUseCase := usecase.NewRoleUseCase(roleRepository, identifierGenerator)
	roleController := controller.NewRoleController(roleUseCase)

	router.SetupRouter(app, permissionController, roleController)

}
