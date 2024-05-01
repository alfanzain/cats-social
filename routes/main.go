package routes

import (
	v1 "cats-social/routes/v1"

	"github.com/gofiber/fiber/v2"
)

func RouteRegister(app *fiber.App) {
	apiGroupRoutes := app.Group("/api")

	v1GroupRoutes := apiGroupRoutes.Group("/v1")
	v1.UserRoutes(v1GroupRoutes)
	v1.CatRoutes(v1GroupRoutes)
}
