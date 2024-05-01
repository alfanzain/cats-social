package v1

import (
	"github.com/gofiber/fiber/v2"
)

func UserRoutes(router fiber.Router) {
	userGroupRoutes := router.Group("/user")
	userGroupRoutes.Post("/register", func(c *fiber.Ctx) error {
		return c.Status(201).JSON("POST /api/v1/user/register")
	})
	userGroupRoutes.Post("/login", func(c *fiber.Ctx) error {
		return c.Status(200).JSON("POST /api/v1/user/login")
	})
}
