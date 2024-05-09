package responses

import (
	"github.com/gofiber/fiber/v2"
)

func ServerErrorInternalServerError(c *fiber.Ctx) error {
	return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
		"message": "Internal server error",
		"data":    nil,
	})
}
