package userhandler

import (
	userrepo "catssocial/repositories/user"
	"catssocial/responses"
	userservice "catssocial/services/user"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type LoginPayload struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=5,max=15"`
}

func (u *UserHandler) Login(c *fiber.Ctx) error {
	user := new(LoginPayload)

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
	result, err := userService.Login(&userservice.LoginPayload{
		Email:    user.Email,
		Password: user.Password,
	})

	if err != nil {
		if err == userservice.ErrUserNotFound {
			return responses.ClientErrorNotFound(c, responses.ErrorPayload{
				Err: err.Error(),
			})
		} else if err == userservice.ErrInvalidPassword {
			return responses.ClientErrorBadRequest(c, responses.ErrorPayload{
				Err: err.Error(),
			})
		}
	}

	return responses.SuccessOK(c, responses.Payload{
		Message: "User logged successfully",
		Data:    result,
	})
}
