package routes

import (
	v1 "catssocial/routes/v1"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	apiGroupRoutes := app.Group("/api")

	v1GroupRoutes := apiGroupRoutes.Group("/v1")
	v1.UserRoutes(v1GroupRoutes,)
	v1.CatRoutes(v1GroupRoutes)
	// v1.MatchRoutes(v1GroupRoutes)
}
