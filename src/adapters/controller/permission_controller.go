package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/krostyle/auth-systme-go/src/application/dto"
	"github.com/krostyle/auth-systme-go/src/application/interfaces"
)

type PermissionController struct {
	permissionUseCase interfaces.PermissionUseCaseInterface
}

func NewPermissionController(permissionUseCase interfaces.PermissionUseCaseInterface) *PermissionController {
	return &PermissionController{
		permissionUseCase: permissionUseCase,
	}
}

func (p *PermissionController) CreatePermission(c *fiber.Ctx) error {
	var permissionDTO dto.PermissionCreateDTO
	if err := c.BodyParser(&permissionDTO); err != nil {
		c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
		return err
	}

	err := p.permissionUseCase.CreatePermission(c.Context(), permissionDTO)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
		return err
	}

	c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "Permission created successfully",
	})
	return nil

}

func (p *PermissionController) GetAllPermissions(c *fiber.Ctx) error {
	permissions, err := p.permissionUseCase.GetAllPermissions(c.Context())
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(permissions)
}
