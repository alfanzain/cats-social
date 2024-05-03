package responses

import (
	"github.com/gofiber/fiber/v2"
)

type Payload struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func SuccessOK(c *fiber.Ctx, payload Payload) error {
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status": payload.Message,
		"data":   payload.Data,
	})
}

func SuccessCreated(c *fiber.Ctx, payload Payload) error {
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"status": payload.Message,
		"data":   payload.Data,
	})
}
