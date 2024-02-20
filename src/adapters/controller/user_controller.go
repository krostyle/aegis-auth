package controller

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/krostyle/auth-systme-go/src/application/dto"
	"github.com/krostyle/auth-systme-go/src/application/interfaces"
)

type UserController struct {
	userUseCase interfaces.UserUseCaseInterface
}

func NewUserController(userUseCase interfaces.UserUseCaseInterface) *UserController {
	return &UserController{
		userUseCase: userUseCase,
	}
}

func (u *UserController) RegisterUser(c *fiber.Ctx) error {
	fmt.Println("Registering user...")
	var user dto.UserCreateDTO
	if err := c.BodyParser(&user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	err := u.userUseCase.RegisterUser(c.Context(), user)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	fmt.Println("User created successfully")

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "User created successfully",
	})
}

func (u *UserController) GetAllUsers(c *fiber.Ctx) error {
	fmt.Println("Getting all users...")
	users, err := u.userUseCase.GetAllUsers(c.Context())
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(users)
}

func (u *UserController) Login(c *fiber.Ctx) error {
	fmt.Println("Logging user...")
	var user dto.UserLoginDTO
	if err := c.BodyParser(&user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	token, err := u.userUseCase.LoginUser(c.Context(), user)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	fmt.Println("User logged in successfully")

	return c.Status(fiber.StatusOK).JSON(token)

}
