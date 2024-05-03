package handlers

import (
	"catssocial/domains/responses"

	"github.com/gofiber/fiber/v2"
)

type User struct {
}

func (u *User) Login(c *fiber.Ctx) error {
	return responses.SuccessOK(c, responses.Payload{
		Message: "User logged successfully",
		Data: fiber.Map{
			"email":       "email",
			"name":        "name",
			"accessToken": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoiNCJ9.FF6Q5BH--DUOIMvmKwuwNfZnExE2MzCjmUYAPJTv5ec",
		},
	})
}

func (u *User) Register(c *fiber.Ctx) error {
	return responses.SuccessCreated(c, responses.Payload{
		Message: "User registered successfully",
		Data: fiber.Map{
			"email":       "email",
			"name":        "name",
			"accessToken": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoiNCJ9.FF6Q5BH--DUOIMvmKwuwNfZnExE2MzCjmUYAPJTv5ec",
		},
	})
}
