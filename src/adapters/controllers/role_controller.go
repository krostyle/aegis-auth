package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/krostyle/auth-systme-go/src/application/dto"
	"github.com/krostyle/auth-systme-go/src/application/interfaces"
)

type RoleController struct {
	roleUseCase interfaces.RoleUseCaseInterface
}

func NewRoleController(roleUseCase interfaces.RoleUseCaseInterface) *RoleController {
	return &RoleController{
		roleUseCase: roleUseCase,
	}
}

func (r *RoleController) CreateRole(c *fiber.Ctx) error {
	var role dto.RoleCreateDTO
	if err := c.BodyParser(&role); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	err := r.roleUseCase.CreateRole(c.Context(), &role)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "Role created successfully",
	})
}

func (r *RoleController) GetAllRoles(c *fiber.Ctx) error {
	roles, err := r.roleUseCase.GetAllRoles(c.Context())
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(roles)
}
