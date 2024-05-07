package v1

import (
	// "catssocial/domains/enums"
	// "fmt"

	"github.com/gofiber/fiber/v2"
)

func MatchRoutes(router fiber.Router) {
	// matchGroupRoutes := router.Group("/cat")
	// matchGroupRoutes.Post("", func(c *fiber.Ctx) error {
	// 	return c.Status(200).JSON("POST /api/v1/cat")
	// })
	// matchGroupRoutes.Get("", func(c *fiber.Ctx) error {
	// 	return c.Status(200).JSON("GET /api/v1/cat")
	// })
	// matchGroupRoutes.Put("/:id", func(c *fiber.Ctx) error {
	// 	response := fmt.Sprintf("PUT /api/v1/cat/%s %s", c.Params("id"), enums.CatRaceEnum.Persian)
	// 	return c.Status(200).JSON(response)
	// })
	// matchGroupRoutes.Delete("/:id", func(c *fiber.Ctx) error {
	// 	response := fmt.Sprintf("DELETE /api/v1/cat/%s", c.Params("id"))
	// 	return c.Status(200).JSON(response)
	// })
}
