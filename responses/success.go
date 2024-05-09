package responses

import (
	"github.com/gofiber/fiber/v2"
)

func SuccessOK(c *fiber.Ctx, payload Payload) error {
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": payload.Message,
		"data":    payload.Data,
	})
}

func SuccessCreated(c *fiber.Ctx, payload Payload) error {
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": payload.Message,
		"data":    payload.Data,
	})
}
