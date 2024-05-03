package v1

import (
	"catssocial/handlers"

	"github.com/gofiber/fiber/v2"
)

func UserRoutes(router fiber.Router) {
	userHandler := handlers.User{}
	userGroupRoutes := router.Group("/user")
	userGroupRoutes.Post("/register", userHandler.Register)
	userGroupRoutes.Post("/login", userHandler.Login)
}
