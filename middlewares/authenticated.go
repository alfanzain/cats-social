package middlewares

import (
	"catssocial/configs"
	"log"

	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v2"
	"github.com/golang-jwt/jwt/v4"
)

func Authenticated() fiber.Handler {
	config, _ := configs.LoadConfig()

	return jwtware.New(jwtware.Config{
		SigningKey: []byte(config.JWTSecret),
		Filter: func(c *fiber.Ctx) bool {
			return false
		},
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			if err != nil {
				log.Println(err)
				return fiber.ErrUnauthorized
			}
			return c.Next()
		},
		SuccessHandler: func(c *fiber.Ctx) error {
			token := c.Locals("user").(*jwt.Token)
			claims := token.Claims.(jwt.MapClaims)

			// This line still bugged
			// if JWT payload user_id is:
			// - not string
			// - not defined
			userID := claims["user_id"].(string)
			
			c.Locals("user_id", userID)
			return c.Next()
		},
	})
}
