package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/krostyle/auth-systme-go/src/application/dto"
	"github.com/krostyle/auth-systme-go/src/application/interfaces"
)

type PermissionController struct {
	permissionUseCase interfaces.PermissionCrudInterface
}

func NewPermissionController(permissionUseCase interfaces.PermissionCrudInterface) *PermissionController {

	return &PermissionController{
		permissionUseCase: permissionUseCase,
	}
}

func (p *PermissionController) CreatePermission(c *fiber.Ctx) error {
	var permission dto.PermissionCreateDTO
	if err := c.BodyParser(&permission); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	err := p.permissionUseCase.CreatePermission(c.Context(), &permission)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "Permission created successfully",
	})

}

func (p *PermissionController) GetPermissionByID(c *fiber.Ctx) error {
	id := c.Params("id")
	permission, err := p.permissionUseCase.GetPermissionByID(c.Context(), id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(permission)
}

func (p *PermissionController) UpdatePermission(c *fiber.Ctx) error {
	id := c.Params("id")
	var permission dto.PermissionUpdateDTO
	if err := c.BodyParser(&permission); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	err := p.permissionUseCase.UpdatePermission(c.Context(), id, &permission)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Permission updated successfully",
	})
}

func (p *PermissionController) DeletePermission(c *fiber.Ctx) error {
	id := c.Params("id")
	err := p.permissionUseCase.DeletePermission(c.Context(), id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Permission deleted successfully",
	})
}

func (p *PermissionController) GetAllPermissions(c *fiber.Ctx) error {
	permissions, err := p.permissionUseCase.GetAllPermissions(c.Context())
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(permissions)
}
