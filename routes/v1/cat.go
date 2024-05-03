package v1

import (
	"catssocial/middlewares"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func CatRoutes(router fiber.Router) {
	catGroupRoutes := router.Group("/cat")
	catGroupRoutes.Post("", middlewares.Authenticated(), func(c *fiber.Ctx) error {
		return c.Status(200).JSON("POST /api/v1/cat")
	})
	catGroupRoutes.Get("", middlewares.Authenticated(), func(c *fiber.Ctx) error {
		return c.Status(200).JSON("GET /api/v1/cat")
	})
	catGroupRoutes.Put("/:id", middlewares.Authenticated(), func(c *fiber.Ctx) error {
		response := fmt.Sprintf("PUT /api/v1/cat/%s", c.Params("id"))
		return c.Status(200).JSON(response)
	})
	catGroupRoutes.Delete("/:id", middlewares.Authenticated(), func(c *fiber.Ctx) error {
		response := fmt.Sprintf("DELETE /api/v1/cat/%s", c.Params("id"))
		return c.Status(200).JSON(response)
	})
}
