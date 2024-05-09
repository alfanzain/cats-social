package userhandler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

type UserHandler struct {
	Log *logrus.Logger
}

type IUserHandler interface {
	Login(c *fiber.Ctx) error
	Register(c *fiber.Ctx) error
}

func New(logger *logrus.Logger) IUserHandler {
	return &UserHandler{
		Log: logger,
	}
}
