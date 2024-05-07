package responses

import (
	"github.com/gofiber/fiber/v2"
)

func ServerInternalServerError(c *fiber.Ctx, message string) error {
	return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
		"message": message,
		"data":    nil,
	})
}
