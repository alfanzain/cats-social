package userhandler

import (
	userservice "catssocial/services/user"
	userrepo "catssocial/repositories/user"
	"catssocial/responses"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type RegisterPayload struct {
	Email    string `json:"email" validate:"required,email"`
	Name     string `json:"name" validate:"required,min=5,max=50"`
	Password string `json:"password" validate:"required,min=5,max=15"`
}

func (u *UserHandler) Register(c *fiber.Ctx) error {
	user := new(RegisterPayload)

	if err := c.BodyParser(&user); err != nil {
		return responses.ClientErrorBadRequest(c, responses.ErrorPayload{
			Err: err.Error(),
		})
	}

	validate := validator.New()
	err := validate.Struct(user)
	if err != nil {
		return responses.ClientErrorBadRequest(c, responses.ErrorPayload{
			Err: err.Error(),
		})
	}

	userService := userservice.New(userrepo.New())
	result, err := userService.Register(&userservice.RegisterPayload{
		Name:     user.Name,
		Email:    user.Email,
		Password: user.Password,
	})

	if err == userservice.ErrEmailAlreadyRegistered {
		return responses.ClientErrorConflict(c, responses.ErrorPayload{
			Message: err.Error(),
			Err:     "",
		})
	}

	if err != nil {
		return responses.ServerErrorInternalServerError(c)
	}

	return responses.SuccessCreated(c, responses.Payload{
		Message: "User registered successfully",
		Data:    result,
	})
}
