package controller

import "github.com/gofiber/fiber/v2"

type HealthCheckController struct {
}

func NewHealthCheckController() *HealthCheckController {
	return &HealthCheckController{}
}

func (h *HealthCheckController) HealthCheck(c *fiber.Ctx) error {
	return c.Status(200).JSON("Health check OK")
}
