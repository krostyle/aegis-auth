package config

import (
	"fmt"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"github.com/krostyle/auth-systme-go/src/adapters/controller"
	"github.com/krostyle/auth-systme-go/src/application/crud"

	"github.com/krostyle/auth-systme-go/src/infrastructure/api/router"
	"github.com/krostyle/auth-systme-go/src/infrastructure/database/postgres_client/configuration"
	"github.com/krostyle/auth-systme-go/src/infrastructure/database/postgres_client/persistence"
	"github.com/krostyle/auth-systme-go/src/infrastructure/service"
)

func Setup(app *fiber.App) {

	if _, err := os.Stat(".env"); err == nil {
		err := godotenv.Load()
		if err != nil {
			panic(fmt.Sprintf("Error loading .env file: %v", err))
		}
	}

	gormDB, err := configuration.NewPostgresDB()
	if err != nil {
		panic(fmt.Sprintf("Error connecting to database: %v", err))
	}

	identifierGenerator := service.NewUUIDGenerator()
	passwordHasher := service.NewPasswordHasher()

	permissionRepository := persistence.NewPermissionRepository(gormDB)
	permissionCrud := crud.NewPermissionCrud(permissionRepository, identifierGenerator)
	permissionController := controller.NewPermissionController(permissionCrud)

	roleRepository := persistence.NewRoleRepository(gormDB)
	roleCrud := crud.NewRoleCrud(roleRepository, identifierGenerator)
	roleController := controller.NewRoleController(roleCrud)

	userRepository := persistence.NewUserRepository(gormDB)
	userCrud := crud.NewUserCrud(userRepository, identifierGenerator, passwordHasher)
	userController := controller.NewUserController(userCrud)

	router.SetupRouter(app, permissionController, roleController, userController)

}
