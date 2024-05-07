package handlers

import (
	"catssocial/configs"
	"catssocial/domains/responses"
	"catssocial/entities"
	"catssocial/helpers"
	"net/http"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type UserHandler struct{}

func NewUserHandler() *UserHandler {
	return &UserHandler{}
}

func (u *UserHandler) Register(c *fiber.Ctx) error {
	var request map[string]string

	if err := c.BodyParser(&request); err != nil {
		return c.SendStatus(http.StatusBadRequest)
	}

	password, _ := helpers.HashPassword(request["password"])

	user := entities.UserRegisterInput{
		Email:    request["email"],
		Name:     request["name"],
		Password: password,
	}

	// Create a new validator instance
	validate := validator.New()

	// Validate the User struct
	err := validate.Struct(user)
	if err != nil {

		return err
		// Validation failed, handle the error
		// return responses.SuccessCreated(c, responses.Payload{
		// 	Message: "User registered successfully",
		// 	Data: fiber.Map{
		// 		"email":       newEmail,
		// 		"name":        newUserName,
		// 		"accessToken": token,
		// 	},
		// })
	}

	// Insert user data into database
	query := `INSERT INTO users (email, name, password) VALUES ($1, $2, $3) RETURNING id, name, email`

	// Prepare the statement
	statement, err := configs.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer statement.Close()

	// Execute the statement with username and email as arguments
	var newUserID int
	var newUserName string
	var newEmail string
	err = statement.QueryRow(user.Email, user.Name, user.Password).Scan(&newUserID, &newUserName, &newEmail)
	if err != nil {
		return err
	}

	// Generate JWT
	token, err := helpers.GenerateAccessToken(strconv.Itoa(newUserID))
	if err != nil {
		return responses.ServerInternalServerError(c, err.Error())
	}

	return responses.SuccessCreated(c, responses.Payload{
		Message: "User registered successfully",
		Data: fiber.Map{
			"email":       newEmail,
			"name":        newUserName,
			"accessToken": token,
		},
	})
}

func (u *UserHandler) Login(c *fiber.Ctx) error {
	// var request map[string]string

	// if err := c.BodyParser(&request); err != nil {
	// 	return err
	// }

	// // Validate request body
	// if err := request.Validate(); err != nil {
	// 	return responses.ErrorBadRequest(ctx, err.Error())
	// }

	// usr := entity.User{
	// 	CredentialValue: req.CredentialValue,
	// 	CredentialType:  string(req.CredentialType),
	// 	Password:        req.Password,
	// }

	// // login user
	// result, err := u.Database.Login(ctx.UserContext(), usr)
	// if err != nil {
	// 	if err.Error() == "USER_NOT_FOUND" {
	// 		return responses.ErrorNotFound(ctx, err.Error())
	// 	}

	// 	if err.Error() == "INVALID_PASSWORD" {
	// 		return responses.ErrorBadRequest(ctx, err.Error())
	// 	}

	// 	return responses.ErrorInternalServerError(ctx, err.Error())
	// }

	// // generate access token
	// accessToken, err := utils.GenerateAccessToken(result.CredentialValue, result.Id)
	// if err != nil {
	// 	return responses.ErrorInternalServerError(ctx, err.Error())
	// }

	return responses.SuccessOK(c, responses.Payload{
		Message: "User logged successfully",
		Data: fiber.Map{
			"email":       "email",
			"name":        "name",
			"accessToken": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoiNCJ9.FF6Q5BH--DUOIMvmKwuwNfZnExE2MzCjmUYAPJTv5ec",
		},
	})
}
