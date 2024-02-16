package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/krostyle/auth-systme-go/src/application/dto"
	"github.com/krostyle/auth-systme-go/src/application/interfaces"
)

type UserController struct {
	userCrud interfaces.UserCrudInterface
}

func NewUserController(userUseCase interfaces.UserCrudInterface) *UserController {
	return &UserController{
		userCrud: userUseCase,
	}
}

func (u *UserController) CreateUser(c *fiber.Ctx) error {
	var user dto.UserCreateDTO
	if err := c.BodyParser(&user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	err := u.userCrud.CreateUser(c.Context(), &user)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "User created successfully",
	})

}

func (u *UserController) GetAllUsers(c *fiber.Ctx) error {
	users, err := u.userCrud.GetAllUsers(c.Context())
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(users)
}

func (u *UserController) GetUserByID(c *fiber.Ctx) error {
	id := c.Params("id")
	user, err := u.userCrud.GetUserByID(c.Context(), id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(user)
}

func (u *UserController) UpdateUser(c *fiber.Ctx) error {
	id := c.Params("id")
	var user dto.UserUpdateDTO
	if err := c.BodyParser(&user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	err := u.userCrud.UpdateUser(c.Context(), id, &user)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "User updated successfully",
	})
}

func (u *UserController) DeleteUser(c *fiber.Ctx) error {
	id := c.Params("id")
	err := u.userCrud.DeleteUser(c.Context(), id)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "User deleted successfully",
	})
}
