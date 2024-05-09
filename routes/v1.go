package routes

import (
	userhandler "catssocial/handlers/user"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

type RouteConfig struct {
	App            *fiber.App
	UserHandler    userhandler.IUserHandler
	AuthMiddleware fiber.Handler
}

func (c *RouteConfig) Setup() {
	c.SetupGuestRoute()
	c.SetupAuthRoute()
}

func (c *RouteConfig) SetupGuestRoute() {
	v1 := c.App.Group("/v1")
	userGroupRoutes := v1.Group("/user")
	userGroupRoutes.Post("/register", c.UserHandler.Register)
	userGroupRoutes.Post("/login", c.UserHandler.Login)
}

func (c *RouteConfig) SetupAuthRoute() {
	c.App.Use(c.AuthMiddleware)
	v1 := c.App.Group("/v1")

	catGroupRoutes := v1.Group("/cat")
	catGroupRoutes.Post("", func(c *fiber.Ctx) error {
		return c.Status(200).JSON("POST /api/v1/cat")
	})
	catGroupRoutes.Get("", func(c *fiber.Ctx) error {
		return c.Status(200).JSON("GET /api/v1/cat")
	})
	catGroupRoutes.Put("/:id", func(c *fiber.Ctx) error {
		response := fmt.Sprintf("PUT /api/v1/cat/%s", c.Params("id"))
		return c.Status(200).JSON(response)
	})
	catGroupRoutes.Delete("/:id", func(c *fiber.Ctx) error {
		response := fmt.Sprintf("DELETE /api/v1/cat/%s", c.Params("id"))
		return c.Status(200).JSON(response)
	})

	matchGroupRoutes := v1.Group("/match")
	matchGroupRoutes.Post("", func(c *fiber.Ctx) error {
		return c.Status(200).JSON("POST /api/v1/match")
	})
}
