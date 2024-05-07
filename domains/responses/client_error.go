package responses

import (
	"github.com/gofiber/fiber/v2"
)

func ClientErrorBadRequest(c *fiber.Ctx, errorPayload ErrorPayload) error {
	return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
		"message": "Invalid request body",
		"error":   errorPayload.Err,
		"data":    nil,
	})
}

func ClientErrorConflict(c *fiber.Ctx, errorPayload ErrorPayload) error {
	return c.Status(fiber.StatusConflict).JSON(fiber.Map{
		"message": errorPayload.Message,
		"error":   errorPayload.Err,
		"data":    nil,
	})
}
